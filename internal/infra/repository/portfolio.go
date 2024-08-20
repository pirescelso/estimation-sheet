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

func (r *estimationRepositoryPostgres) CreatePortfolio(ctx context.Context, portfolio *domain.Portfolio) error {
	err := r.queries.InsertPortfolio(ctx, db.InsertPortfolioParams{
		PortfolioID: portfolio.PortfolioID,
		BaselineID:  portfolio.BaselineID,
		PlanID:      portfolio.PlanID,
		StartDate:   pgtype.Date{Time: portfolio.StartDate, Valid: true},
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return common.NewConflictError(fmt.Errorf("portfolio for baseline id %s and plan id %s already exists", portfolio.BaselineID, portfolio.PlanID))
			}
			return common.NewConflictError(err)
		}
	}

	return nil
}

func (r *estimationRepositoryPostgres) GetPortfolio(ctx context.Context, portfolioID string) (*domain.Portfolio, error) {
	portfolioModel, err := r.queries.FindPortfolioById(ctx, portfolioID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, common.NewNotFoundError(fmt.Errorf("portfolio with id %s not found", portfolioID))
		}
		return nil, err
	}

	props := domain.RestorePortfolioProps{
		PortfolioID: portfolioModel.PortfolioID,
		BaselineID:  portfolioModel.BaselineID,
		PlanID:      portfolioModel.PlanID,
		StartDate:   portfolioModel.StartDate.Time,
		CreatedAt:   portfolioModel.CreatedAt.Time,
		UpdatedAt:   portfolioModel.UpdatedAt.Time,
	}

	portfolio := domain.RestorePortfolio(props)
	err = portfolio.Validate()
	if err != nil {
		return nil, err
	}

	return portfolio, nil
}

func (r *estimationRepositoryPostgres) ValidatePortfolioUniqueBaselineByPlan(ctx context.Context, planID string, baselineCode string) error {
	_, err := r.queries.FindPortfolioByPlanIdAndBaselineCode(ctx,
		db.FindPortfolioByPlanIdAndBaselineCodeParams{
			PlanID: planID,
			Code:   baselineCode,
		},
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}

	return common.NewConflictError(fmt.Errorf("portfolio for plan id %s and baseline code %s already exists", planID, baselineCode))
}

func (r *estimationRepositoryPostgres) UpdatePortfolio(ctx context.Context, portfolio *domain.Portfolio) error {
	err := r.queries.UpdatePortfolio(ctx, db.UpdatePortfolioParams{
		PortfolioID: portfolio.PortfolioID,
		BaselineID:  portfolio.BaselineID,
		PlanID:      portfolio.PlanID,
		StartDate:   pgtype.Date{Time: portfolio.StartDate, Valid: true},
		UpdatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return common.NewConflictError(fmt.Errorf("portfolio for baseline id %s and plan id %s already exists", portfolio.BaselineID, portfolio.PlanID))
			}
			return common.NewConflictError(err)
		}
	}

	return nil
}

func (r *estimationRepositoryPostgres) DeletePortfolio(ctx context.Context, portfolioID string) error {
	_, err := r.queries.DeletePortfolio(ctx, portfolioID)
	return err
}

func (r *estimationRepositoryPostgres) CountPortfoliosByPlanId(ctx context.Context, planID string) (int64, error) {
	return r.queries.CountPortfoliosByPlanId(ctx, planID)
}

func (r *estimationRepositoryPostgres) CountPortfoliosByBaselineId(ctx context.Context, baselineID string) (int64, error) {
	return r.queries.CountPortfoliosByBaselineId(ctx, baselineID)
}
