package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/celsopires1999/estimation/internal/mapper"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *EstimationService) CreateUser(ctx context.Context, input CreateUserInputDTO) (*CreateUserOutputDTO, error) {
	_, err := s.queries.FindUserByEmail(ctx, input.Email)
	if err == nil {
		return nil, common.NewConflictError(fmt.Errorf("user with email %s already exists", input.Email))
	}
	params := db.InsertUserParams{
		UserID:    uuid.New().String(),
		Email:     input.Email,
		UserName:  input.UserName,
		Name:      input.Name,
		UserType:  input.UserType,
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	err = s.queries.InsertUser(ctx, params)
	if err != nil {
		return nil, err
	}

	created, err := s.queries.FindUserById(ctx, params.UserID)
	if err != nil {
		return nil, err
	}

	userOutput := mapper.UserOutputFromDb(created)

	output := &CreateUserOutputDTO{userOutput}

	return output, nil
}

func (s *EstimationService) UpdateUser(ctx context.Context, input UpdateUserInputDTO) (*UpdateUserOutputDTO, error) {
	current, err := s.queries.FindUserById(ctx, input.UserID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, common.NewNotFoundError(fmt.Errorf("user with id: %s not found", input.UserID))
		}
		return nil, err
	}
	params := db.UpdateUserParams{}

	params.UserID = current.UserID
	params.UpdatedAt = pgtype.Timestamp{Time: time.Now(), Valid: true}

	params.Email = copyPatch(input.Email, current.Email)
	params.UserName = copyPatch(input.UserName, current.UserName)
	params.Name = copyPatch(input.Name, current.Name)
	params.UserType = copyPatch(input.UserType, current.UserType)

	err = s.queries.UpdateUser(ctx, params)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil, common.NewConflictError(fmt.Errorf("user with email %s already exists", *input.Email))
			}
			return nil, common.NewConflictError(err)
		}
	}

	updated, err := s.queries.FindUserById(ctx, params.UserID)
	if err != nil {
		return nil, err
	}

	userOutput := mapper.UserOutputFromDb(updated)

	output := &UpdateUserOutputDTO{userOutput}

	return output, nil
}

func (s *EstimationService) GetUser(ctx context.Context, userID string) (*GetUserOutputDTO, error) {
	user, err := s.queries.FindUserById(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, common.NewNotFoundError(fmt.Errorf("user with id %s not found", userID))
		}
		return nil, err
	}

	userOutput := mapper.UserOutputFromDb(user)

	output := &GetUserOutputDTO{userOutput}

	return output, nil
}

func (s *EstimationService) DeleteUser(ctx context.Context, userID string) error {
	_, err := s.queries.DeleteUser(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return common.NewNotFoundError(fmt.Errorf("user with id %s not found", userID))
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23503" {
				if pgErr.ConstraintName == "baselines_estimator_id_fkey" {
					return common.NewConflictError(fmt.Errorf("cannot delete user id %s with baseline", userID))
				}
				return common.NewConflictError(fmt.Errorf("cannot delete user id %s with relations: %w", userID, err))
			}
		}
		return common.NewConflictError(fmt.Errorf("cannot delete user id %s: %w", userID, err))
	}
	return err
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

type GetUserOutputDTO struct {
	mapper.UserOutput
}

func (s *EstimationService) ListUsers(ctx context.Context, input ListUsersInputDTO) (*ListUsersOutputDTO, error) {
	users, err := s.queries.FindAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	usersOutput := make([]mapper.UserOutput, len(users))
	for i, user := range users {
		usersOutput[i] = mapper.UserOutputFromDb(user)
	}

	return &ListUsersOutputDTO{usersOutput}, nil
}

type ListUsersInputDTO struct{}

type ListUsersOutputDTO struct {
	Users []mapper.UserOutput `json:"users"`
}
