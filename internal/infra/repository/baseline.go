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

func (r *estimationRepositoryPostgres) CreateBaseline(ctx context.Context, baseline *domain.Baseline) error {
	err := r.queries.InsertBaseline(ctx, db.InsertBaselineParams{
		BaselineID:  baseline.BaselineID,
		Code:        baseline.Code,
		Review:      baseline.Review,
		Title:       baseline.Title,
		Description: pgtype.Text{String: baseline.Description, Valid: true},
		StartDate:   pgtype.Date{Time: baseline.StartDate, Valid: true},
		Duration:    baseline.Duration,
		ManagerID:   baseline.ManagerID,
		EstimatorID: baseline.EstimatorID,
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		return checkRelationsError(baseline, err)
	}

	return err
}

func (r *estimationRepositoryPostgres) GetBaseline(ctx context.Context, baselineID string) (*domain.Baseline, error) {
	baselineModel, err := r.queries.FindBaselineById(ctx, baselineID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, common.NewNotFoundError(errors.New("baseline not found"))
		}
		return nil, err
	}

	props := domain.RestoreBaselineProps{
		BaselineID:  baselineModel.BaselineID,
		Code:        baselineModel.Code,
		Review:      baselineModel.Review,
		Title:       baselineModel.Title,
		Description: baselineModel.Description.String,
		StartDate:   baselineModel.StartDate.Time,
		Duration:    baselineModel.Duration,
		ManagerID:   baselineModel.ManagerID,
		EstimatorID: baselineModel.EstimatorID,
		CreatedAt:   baselineModel.CreatedAt.Time,
		UpdatedAt:   baselineModel.UpdatedAt.Time,
	}

	baseline := domain.RestoreBaseline(props)
	err = baseline.Validate()
	if err != nil {
		return nil, err
	}

	return baseline, nil
}

func (r *estimationRepositoryPostgres) UpdateBaseline(ctx context.Context, baseline *domain.Baseline) error {
	err := r.queries.UpdateBaseline(ctx, db.UpdateBaselineParams{
		BaselineID:  baseline.BaselineID,
		Code:        baseline.Code,
		Review:      baseline.Review,
		Title:       baseline.Title,
		Description: pgtype.Text{String: baseline.Description, Valid: true},
		StartDate:   pgtype.Date{Time: baseline.StartDate, Valid: true},
		Duration:    baseline.Duration,
		ManagerID:   baseline.ManagerID,
		EstimatorID: baseline.EstimatorID,
		UpdatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		return checkRelationsError(baseline, err)
	}

	return err
}

func (r *estimationRepositoryPostgres) DeleteBaseline(ctx context.Context, baselineID string) error {
	_, err := r.queries.DeleteBaseline(ctx, baselineID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return common.NewNotFoundError(errors.New("baseline not found"))
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23503" {
				if pgErr.ConstraintName == "costs_baseline_id_fkey" {
					return common.NewConflictError(fmt.Errorf("cannot delete baseline id %s with costs", baselineID))
				}
				return common.NewConflictError(fmt.Errorf("cannot delete baseline id %s with relations: %w", baselineID, err))
			}
		}
		return common.NewConflictError(fmt.Errorf("cannot delete baseline id %s: %w", baselineID, err))
	}
	return err
}

func checkRelationsError(baseline *domain.Baseline, err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			return common.NewConflictError(fmt.Errorf("baseline code %s with review %d already exists", baseline.Code, baseline.Review))
		}
		if pgErr.Code == "23503" {
			if pgErr.ConstraintName == "baselines_manager_id_fkey" {
				return common.NewConflictError(fmt.Errorf("manager id %s does not exist", baseline.ManagerID))
			}
			if pgErr.ConstraintName == "baselines_estimator_id_fkey" {
				return common.NewConflictError(fmt.Errorf("estimator id %s does not exist", baseline.EstimatorID))
			}
		}
		return common.NewConflictError(err)
	}

	return err
}
