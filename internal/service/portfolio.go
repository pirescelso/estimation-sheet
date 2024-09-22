package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/celsopires1999/estimation/internal/mapper"
	"github.com/jackc/pgx/v5"
)

func (s *EstimationService) GetPortfolio(ctx context.Context, input GetPortfolioInputDTO) (*GetPortfolioOutputDTO, error) {
	portfolio, err := s.queries.FindPortfolioByIdWithRelations(ctx, input.PortfolioID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, common.NewNotFoundError(fmt.Errorf("portfolio with id %s not found", input.PortfolioID))
		}
		return nil, err
	}

	portfolioOutput := mapper.PortfolioOutputFromDb(db.PortfolioRow(portfolio))

	budgets, err := s.queries.FindBudgetsByPortfolioIdWithRelations(ctx, input.PortfolioID)
	if err != nil {
		return nil, err
	}

	budgetsOutput := make([]mapper.BudgetOutput, len(budgets))

	for i, budget := range budgets {
		allocations, err := s.queries.FindBudgetAllocations(ctx, budget.BudgetID)
		if err != nil {
			return nil, err
		}

		yearlyAllocations, err := s.queries.FindBudgetAllocationsGroupedByYear(ctx, budget.BudgetID)
		if err != nil {
			return nil, err
		}

		budgetsOutput[i] = mapper.BudgetOutputFromDb(db.BudgetRow(budget), allocations, yearlyAllocations)
	}
	portfolioOutput.Budgets = budgetsOutput

	workloads, err := s.queries.FindWorkloadsByPortfolioIdWithRelations(ctx, input.PortfolioID)
	if err != nil {
		return nil, err
	}

	workloadsOutput := make([]mapper.WorkloadOutput, len(workloads))

	for i, workload := range workloads {
		allocations, err := s.queries.FindWorkloadAllocations(ctx, workload.WorkloadID)
		if err != nil {
			return nil, err
		}

		yearlyAllocations, err := s.queries.FindWorkloadAllocationsGroupedByYear(ctx, workload.WorkloadID)
		if err != nil {
			return nil, err
		}

		workloadsOutput[i] = mapper.WorkloadOutputFromDb(db.WorkloadRow(workload), allocations, yearlyAllocations)
	}
	portfolioOutput.Workloads = workloadsOutput

	return &GetPortfolioOutputDTO{portfolioOutput}, nil
}

type GetPortfolioInputDTO struct {
	PortfolioID string `json:"portfolio_id"`
}

type GetPortfolioOutputDTO struct {
	mapper.PortfolioOutput
}

func (s *EstimationService) ListPortfoliosByPlanID(ctx context.Context, input ListPortfoliosInputDTO) (*ListPortfoliosOutputDTO, error) {
	portfolios, err := s.queries.FindAllPortfoliosByPlanIdWithRelations(ctx, input.PlanID)
	if err != nil {
		return nil, err
	}

	portfoliosOutput := make([]mapper.PortfolioOutput, len(portfolios))
	for i, portfolio := range portfolios {
		portfoliosOutput[i] = mapper.PortfolioOutput{
			PortfolioID: portfolio.PortfolioID,
			PlanCode:    portfolio.PlanCode,
			Code:        portfolio.Code,
			Review:      portfolio.Review,
			Title:       portfolio.Title,
			Description: portfolio.Description.String,
			StartDate:   portfolio.StartDate.Time,
			Duration:    portfolio.Duration,
			Manager:     portfolio.Manager,
			Estimator:   portfolio.Estimator,
			CreatedAt:   portfolio.CreatedAt.Time,
			UpdatedAt:   portfolio.UpdatedAt.Time,
		}
	}

	return &ListPortfoliosOutputDTO{portfoliosOutput}, nil
}

func (s *EstimationService) ListAllPortfolios(ctx context.Context) (*ListPortfoliosOutputDTO, error) {
	portfolios, err := s.queries.FindAllPortfoliosWithRelations(ctx)
	if err != nil {
		return nil, err
	}

	portfoliosOutput := make([]mapper.PortfolioOutput, len(portfolios))
	for i, portfolio := range portfolios {
		portfoliosOutput[i] = mapper.PortfolioOutputFromDb(db.PortfolioRow(portfolio))
	}

	return &ListPortfoliosOutputDTO{portfoliosOutput}, nil
}

type ListPortfoliosInputDTO struct {
	PlanID string `json:"plan_id"`
}

type ListPortfoliosOutputDTO struct {
	Portfolios []mapper.PortfolioOutput `json:"portfolios"`
}
