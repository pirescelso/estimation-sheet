package usecase

import (
	"context"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/mapper"
)

type CreateCompetenceUseCase struct {
	repository domain.EstimationRepository
}

type CreateCompetenceInputDTO struct {
	Code string `json:"code" validate:"required,max=20"`
	Name string `json:"name" validate:"required,max=50"`
}

type CreateCompetenceOutputDTO struct {
	mapper.CompetenceOutput
}

func NewCreateCompetenceUseCase(repo domain.EstimationRepository) *CreateCompetenceUseCase {
	return &CreateCompetenceUseCase{repo}
}

func (uc *CreateCompetenceUseCase) Execute(ctx context.Context, input CreateCompetenceInputDTO) (*CreateCompetenceOutputDTO, error) {

	competence := domain.NewCompetence(input.Code, input.Name)
	if err := uc.repository.CreateCompetence(ctx, competence); err != nil {
		return nil, err
	}

	createdCompetence, err := uc.repository.GetCompetence(ctx, competence.CompetenceID)
	if err != nil {
		return nil, err
	}

	output := mapper.CompetenceOutputFromDomain(*createdCompetence)

	return &CreateCompetenceOutputDTO{output}, nil
}

type UpdateCompetenceUseCase struct {
	repository domain.EstimationRepository
}

type UpdateCompetenceInputDTO struct {
	CompetenceID string  `json:"competence_id" validate:"required,uuid4"`
	Code         *string `json:"code" validate:"omitempty,max=20"`
	Name         *string `json:"name" validate:"omitempty,max=50"`
}

type UpdateCompetenceOutputDTO struct {
	mapper.CompetenceOutput
}

func NewUpdateCompetenceUseCase(repo domain.EstimationRepository) *UpdateCompetenceUseCase {
	return &UpdateCompetenceUseCase{repo}
}

func (uc *UpdateCompetenceUseCase) Execute(ctx context.Context, input UpdateCompetenceInputDTO) (*UpdateCompetenceOutputDTO, error) {
	competence, err := uc.repository.GetCompetence(ctx, input.CompetenceID)
	if err != nil {
		return nil, err
	}

	competence.ChangeCode(input.Code)
	competence.ChangeName(input.Name)

	err = competence.Validate()
	if err != nil {
		return nil, err
	}

	err = uc.repository.UpdateCompetence(ctx, competence)
	if err != nil {
		return nil, err
	}

	updated, err := uc.repository.GetCompetence(ctx, competence.CompetenceID)
	if err != nil {
		return nil, err
	}

	output := mapper.CompetenceOutputFromDomain(*updated)

	return &UpdateCompetenceOutputDTO{output}, nil
}

type DeleteCompetenceUseCase struct {
	repository domain.EstimationRepository
}

type DeleteCompetenceInputDTO struct {
	CompetenceID string `json:"competence_id" validate:"required"`
}

type DeleteCompetenceOutputDTO struct{}

func NewDeleteCompetenceUseCase(repo domain.EstimationRepository) *DeleteCompetenceUseCase {
	return &DeleteCompetenceUseCase{repo}
}

func (uc *DeleteCompetenceUseCase) Execute(ctx context.Context, input DeleteCompetenceInputDTO) (*DeleteCompetenceOutputDTO, error) {
	err := uc.repository.DeleteCompetence(ctx, input.CompetenceID)
	if err != nil {
		return nil, err
	}
	return &DeleteCompetenceOutputDTO{}, nil
}

type GetCompetenceUseCase struct {
	repository domain.EstimationRepository
}

type GetCompetenceInputDTO struct {
	CompetenceID string `json:"competence_id" validate:"required"`
}

type GetCompetenceOutputDTO struct {
	mapper.CompetenceOutput
}

func NewGetCompetenceUseCase(repo domain.EstimationRepository) *GetCompetenceUseCase {
	return &GetCompetenceUseCase{repo}
}

func (uc *GetCompetenceUseCase) Execute(ctx context.Context, input GetCompetenceInputDTO) (*GetCompetenceOutputDTO, error) {
	competence, err := uc.repository.GetCompetence(ctx, input.CompetenceID)
	if err != nil {
		return nil, err
	}
	output := mapper.CompetenceOutputFromDomain(*competence)
	return &GetCompetenceOutputDTO{output}, nil
}
