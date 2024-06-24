package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/celsopires1999/estimation/internal/common"
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

	portfolioOutput := mapper.PortfolioOutput{
		PortfolioID: portfolio.PortfolioID,
		PlanCode:    portfolio.PlanCode,
		Code:        portfolio.Code,
		Title:       portfolio.Title,
		Description: portfolio.Description.String,
		StartDate:   portfolio.StartDate.Time,
		Manager:     portfolio.Manager,
		Estimator:   portfolio.Estimator,
		CreatedAt:   portfolio.CreatedAt.Time,
		UpdatedAt:   portfolio.UpdatedAt.Time,
	}

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

		allocs := make([]mapper.BudgetAllocationOutput, len(allocations))
		for j, alloc := range allocations {
			allocs[j] = mapper.BudgetAllocationOutput{
				Year:   alloc.AllocationDate.Time.Year(),
				Month:  int(alloc.AllocationDate.Time.Month()),
				Amount: alloc.Amount,
			}
		}

		budgetsOutput[i] = mapper.BudgetOutput{
			BudgetID:          budget.BudgetID,
			PortfolioID:       budget.PortfolioID,
			CostType:          budget.CostType,
			Description:       budget.Description,
			Comment:           budget.Comment.String,
			CostAmount:        budget.CostAmount,
			CostCurrency:      budget.CostCurrency,
			CostTax:           budget.CostTax,
			Amount:            budget.Amount,
			BudgetAllocations: allocs,
			CreatedAt:         budget.CreatedAt.Time,
			UpdatedAt:         budget.UpdatedAt.Time,
		}
	}
	return &GetPortfolioOutputDTO{portfolioOutput, budgetsOutput}, nil
}

type GetPortfolioInputDTO struct {
	PortfolioID string `json:"portfolio_id"`
}

type GetPortfolioOutputDTO struct {
	Portfolio mapper.PortfolioOutput `json:"portfolio"`
	Budgets   []mapper.BudgetOutput  `json:"budgets"`
}
