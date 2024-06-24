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

func (r *estimationRepositoryPostgres) CreatePlan(ctx context.Context, plan *domain.Plan) error {
	err := r.queries.InsertPlan(ctx, db.InsertPlanParams{
		PlanID:      plan.PlanID,
		Code:        plan.Code,
		Name:        plan.Name,
		Assumptions: plan.Assumptions,
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return common.NewConflictError(fmt.Errorf("plan code %s already exists", plan.Code))
			}
			return common.NewConflictError(err)
		}
	}

	return err
}

func (r *estimationRepositoryPostgres) GetPlan(ctx context.Context, planID string) (*domain.Plan, error) {
	planModel, err := r.queries.FindPlanById(ctx, planID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, common.NewNotFoundError(fmt.Errorf("plan with id %s not found", planID))
		}
		return nil, err
	}

	props := domain.RestorePlanProps{
		PlanID:      planModel.PlanID,
		Code:        planModel.Code,
		Name:        planModel.Name,
		Assumptions: planModel.Assumptions,
		CreatedAt:   planModel.CreatedAt.Time,
		UpdatedAt:   planModel.UpdatedAt.Time,
	}

	plan := domain.RestorePlan(props)
	err = plan.Validate()
	if err != nil {
		return nil, err
	}
	return plan, nil
}

func (r *estimationRepositoryPostgres) UpdatePlan(ctx context.Context, plan *domain.Plan) error {
	_, err := r.queries.UpdatePlan(ctx, db.UpdatePlanParams{
		PlanID:      plan.PlanID,
		Code:        plan.Code,
		Name:        plan.Name,
		Assumptions: plan.Assumptions,
		UpdatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return common.NewNotFoundError(fmt.Errorf("plan with id %s not found", plan.PlanID))
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return common.NewConflictError(fmt.Errorf("plan code %s already exists", plan.Code))
			}
			return common.NewConflictError(err)
		}
		return err
	}

	return err
}

func (r *estimationRepositoryPostgres) DeletePlan(ctx context.Context, planID string) error {
	_, err := r.queries.DeletePlan(ctx, planID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return common.NewNotFoundError(fmt.Errorf("plan with id  %s not found", planID))
		}
		return err
	}
	return nil
}

func (r *estimationRepositoryPostgres) ValidatePlan(ctx context.Context, planID string) error {
	_, err := r.queries.FindPlanById(ctx, planID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return common.NewNotFoundError(fmt.Errorf("plan with id %s not found", planID))
		}
		return err
	}
	return nil
}
