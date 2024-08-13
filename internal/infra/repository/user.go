package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *estimationRepositoryPostgres) CreateUser(ctx context.Context, user *domain.User) error {
	err := r.queries.InsertUser(ctx, db.InsertUserParams{
		UserID:    user.UserID,
		Email:     user.Email,
		UserName:  user.UserName,
		Name:      user.Name,
		UserType:  user.UserType.String(),
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return common.NewConflictError(fmt.Errorf("user with email %s already exists", user.Email))
			}
			return common.NewConflictError(err)
		}
	}

	return err
}

func (r *estimationRepositoryPostgres) GetUser(ctx context.Context, userID string) (*domain.User, error) {
	userModel, err := r.queries.FindUserById(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, common.NewNotFoundError(fmt.Errorf("user with id %s not found", userID))
		}
		return nil, err
	}

	props := domain.RestoreUserProps{
		UserID:    userModel.UserID,
		Email:     userModel.Email,
		UserName:  userModel.UserName,
		Name:      userModel.Name,
		UserType:  domain.UserType(userModel.UserType),
		CreatedAt: userModel.CreatedAt.Time,
		UpdatedAt: userModel.UpdatedAt.Time,
	}

	user := domain.RestoreUser(props)
	err = user.Validate()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *estimationRepositoryPostgres) UpdateUser(ctx context.Context, user *domain.User) error {
	err := r.queries.UpdateUser(ctx, db.UpdateUserParams{
		UserID:    user.UserID,
		Email:     user.Email,
		UserName:  user.UserName,
		Name:      user.Name,
		UserType:  user.UserType.String(),
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return common.NewNotFoundError(fmt.Errorf("user with id %s not found", user.UserID))
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return common.NewConflictError(fmt.Errorf("user with email %s already exists", user.Email))
			}
			return common.NewConflictError(err)
		}
	}

	return err
}

func (r *estimationRepositoryPostgres) DeleteUser(ctx context.Context, userID string) error {
	_, err := r.queries.DeleteUser(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return common.NewNotFoundError(fmt.Errorf("user with id %s not found", userID))
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23503" {
				if pgErr.ConstraintName == "baselines_manager_id_fkey" {
					return common.NewConflictError(fmt.Errorf("cannot delete user id %s with baseline as manager", userID))
				}
				if pgErr.ConstraintName == "baselines_estimator_id_fkey" {
					return common.NewConflictError(fmt.Errorf("cannot delete user id %s with baseline as estimator", userID))
				}
				return common.NewConflictError(fmt.Errorf("cannot delete user id %s with relations: %w", userID, err))
			}
		}
		return common.NewConflictError(fmt.Errorf("cannot delete user id %s: %w", userID, err))
	}
	return nil
}
