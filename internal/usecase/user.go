package usecase

import (
	"context"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/mapper"
)

type CreateUserUseCase struct {
	repository domain.EstimationRepository
}

type CreateUserInputDTO struct {
	Email    string `json:"email" validate:"required,email"`
	UserName string `json:"user_name" validate:"required"`
	Name     string `json:"name" validate:"required"`
	UserType string `json:"user_type" validate:"required,oneof=manager estimator"`
}

type CreateUserOutputDTO struct {
	mapper.UserOutput
}

func NewCreateUserUseCase(repo domain.EstimationRepository) *CreateUserUseCase {
	return &CreateUserUseCase{repo}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, input CreateUserInputDTO) (*CreateUserOutputDTO, error) {
	user := domain.NewUser(input.Email, input.UserName, input.Name, domain.UserType(input.UserType))
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	if err := uc.repository.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	createdUser, err := uc.repository.GetUser(ctx, user.UserID)
	if err != nil {
		return nil, err
	}

	output := mapper.UserOutputFromDomain(*createdUser)

	return &CreateUserOutputDTO{output}, nil
}

type UpdateUserUseCase struct {
	repository domain.EstimationRepository
}

type UpdateUserInputDTO struct {
	UserID   string  `json:"user_id" validate:"required,uuid4"`
	Email    *string `json:"email" validate:"omitempty,email"`
	UserName *string `json:"user_name" validate:"omitempty"`
	Name     *string `json:"name" validate:"omitempty"`
	UserType *string `json:"user_type" validate:"omitempty,oneof=manager estimator"`
}

type UpdateUserOutputDTO struct {
	mapper.UserOutput
}

func NewUpdateUserUseCase(repo domain.EstimationRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{repo}
}

func (uc *UpdateUserUseCase) Execute(ctx context.Context, input UpdateUserInputDTO) (*UpdateUserOutputDTO, error) {
	user, err := uc.repository.GetUser(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	user.ChangeEmail(input.Email)
	user.ChangeUserName(input.UserName)
	user.ChangeName(input.Name)
	user.ChangeUserType(input.UserType)

	err = user.Validate()
	if err != nil {
		return nil, err
	}

	err = uc.repository.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	updated, err := uc.repository.GetUser(ctx, user.UserID)
	if err != nil {
		return nil, err
	}

	output := mapper.UserOutputFromDomain(*updated)

	return &UpdateUserOutputDTO{output}, nil
}

type DeleteUserUseCase struct {
	repository domain.EstimationRepository
}

type DeleteUserInputDTO struct {
	UserID string `json:"user_id" validate:"required,uuid4"`
}

type DeleteUserOutputDTO struct{}

func NewDeleteUserUseCase(repo domain.EstimationRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{repo}
}

func (uc *DeleteUserUseCase) Execute(ctx context.Context, input DeleteUserInputDTO) (*DeleteUserOutputDTO, error) {
	err := uc.repository.DeleteCompetence(ctx, input.UserID)
	if err != nil {
		return nil, err
	}
	return &DeleteUserOutputDTO{}, nil
}

type GetUserUseCase struct {
	repository domain.EstimationRepository
}

type GetUserInputDTO struct {
	UserID string `json:"competence_id" validate:"required,uuid4"`
}

type GetUserOutputDTO struct {
	mapper.UserOutput
}

func NewGetUserUseCase(repo domain.EstimationRepository) *GetUserUseCase {
	return &GetUserUseCase{repo}
}

func (uc *GetUserUseCase) Execute(ctx context.Context, input GetUserInputDTO) (*GetUserOutputDTO, error) {
	user, err := uc.repository.GetUser(ctx, input.UserID)
	if err != nil {
		return nil, err
	}
	output := mapper.UserOutputFromDomain(*user)
	return &GetUserOutputDTO{output}, nil
}
