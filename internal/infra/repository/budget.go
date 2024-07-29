package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
)

func (r *estimationRepositoryPostgres) CreateBudget(ctx context.Context, budget *domain.Budget) error {
	err := r.queries.InsertBudget(ctx, db.InsertBudgetParams{
		BudgetID:    budget.BudgetID,
		PortfolioID: budget.PortfolioID,
		CostID:      budget.CostID,
		Amount:      budget.Amount,
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		return err
	}

	budgetAllocations := db.BulkInsertBudgetAllocationParams{}
	for _, allocation := range budget.BudgetAllocations {
		budgetAllocations.Column1 = append(budgetAllocations.Column1, uuid.New().String())
		budgetAllocations.Column2 = append(budgetAllocations.Column2, budget.BudgetID)
		budgetAllocations.Column3 = append(budgetAllocations.Column3, pgtype.Date{Time: allocation.AllocationDate, Valid: true})
		budgetAllocations.Column4 = append(budgetAllocations.Column4, allocation.Amount)
		budgetAllocations.Column5 = append(budgetAllocations.Column5, pgtype.Timestamp{Time: time.Now(), Valid: true})
	}
	err = r.queries.BulkInsertBudgetAllocation(ctx, budgetAllocations)

	if err != nil {
		return err
	}

	return nil
}

func (r *estimationRepositoryPostgres) CreateBudgetMany(ctx context.Context, budgets []*domain.Budget) error {
	budgetsParams := db.BulkInsertBudgetParams{}
	budgetAllocations := db.BulkInsertBudgetAllocationParams{}

	for _, budget := range budgets {
		budgetsParams.Column1 = append(budgetsParams.Column1, budget.BudgetID)
		budgetsParams.Column2 = append(budgetsParams.Column2, budget.PortfolioID)
		budgetsParams.Column3 = append(budgetsParams.Column3, budget.CostID)
		budgetsParams.Column4 = append(budgetsParams.Column4, budget.Amount)
		budgetsParams.Column5 = append(budgetsParams.Column5, pgtype.Timestamp{Time: time.Now(), Valid: true})
		for _, allocation := range budget.BudgetAllocations {
			budgetAllocations.Column1 = append(budgetAllocations.Column1, uuid.New().String())
			budgetAllocations.Column2 = append(budgetAllocations.Column2, budget.BudgetID)
			budgetAllocations.Column3 = append(budgetAllocations.Column3, pgtype.Date{Time: allocation.AllocationDate, Valid: true})
			budgetAllocations.Column4 = append(budgetAllocations.Column4, allocation.Amount)
			budgetAllocations.Column5 = append(budgetAllocations.Column5, pgtype.Timestamp{Time: time.Now(), Valid: true})
		}
	}
	err := r.queries.BulkInsertBudget(ctx, budgetsParams)
	if err != nil {
		return err
	}
	err = r.queries.BulkInsertBudgetAllocation(ctx, budgetAllocations)
	if err != nil {
		return err
	}

	return nil
}

func (r *estimationRepositoryPostgres) GetBudget(ctx context.Context, budgetID string) (*domain.Budget, error) {
	budgetModel, err := r.queries.FindBudgetById(ctx, budgetID)
	if err != nil {
		return nil, err
	}

	allocationModels, err := r.queries.FindBudgetAllocations(ctx, budgetID)
	if err != nil {
		return nil, err
	}

	allocations := make([]domain.BudgetAllocation, len(allocationModels))
	for i, allocation := range allocationModels {
		allocations[i] = domain.BudgetAllocation{
			AllocationDate: allocation.AllocationDate.Time,
			Amount:         allocation.Amount,
		}
	}

	props := domain.RestoreBudgetProps{
		BudgetID:          budgetModel.BudgetID,
		PortfolioID:       budgetModel.PortfolioID,
		CostID:            budgetModel.CostID,
		Amount:            budgetModel.Amount,
		BudgetAllocations: allocations,
		CreatedAt:         budgetModel.CreatedAt.Time,
		UpdatedAt:         budgetModel.UpdatedAt.Time,
	}

	budget := domain.RestoreBudget(props)
	err = budget.Validate()
	if err != nil {
		return nil, err
	}

	return budget, nil
}

func (r *estimationRepositoryPostgres) UpdateBudget(ctx context.Context, budget *domain.Budget) error {
	err := r.queries.UpdateBudget(ctx, db.UpdateBudgetParams{
		BudgetID:    budget.BudgetID,
		PortfolioID: budget.PortfolioID,
		CostID:      budget.CostID,
		Amount:      budget.Amount,
		UpdatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		return err
	}

	_, err = r.queries.DeleteBudgetAllocations(ctx, budget.CostID)

	if err != nil {
		return err
	}

	budgetAllocations := db.BulkInsertBudgetAllocationParams{}
	for _, allocation := range budget.BudgetAllocations {
		budgetAllocations.Column1 = append(budgetAllocations.Column1, uuid.New().String())
		budgetAllocations.Column2 = append(budgetAllocations.Column2, budget.CostID)
		budgetAllocations.Column3 = append(budgetAllocations.Column3, pgtype.Date{Time: allocation.AllocationDate, Valid: true})
		budgetAllocations.Column4 = append(budgetAllocations.Column4, allocation.Amount)
		budgetAllocations.Column5 = append(budgetAllocations.Column5, pgtype.Timestamp{Time: time.Now(), Valid: true})
	}
	err = r.queries.BulkInsertBudgetAllocation(ctx, budgetAllocations)

	return err
}

func (r *estimationRepositoryPostgres) DeleteBudget(ctx context.Context, budgetID string) error {
	_, err := r.queries.DeleteBudgetAllocations(ctx, budgetID)
	if err != nil {
		return err
	}

	_, err = r.queries.DeleteBudget(ctx, budgetID)

	return err
}

func (r *estimationRepositoryPostgres) DeleteBudgetsByPortfolioID(ctx context.Context, portfolioID string) error {
	budgetModels, err := r.queries.FindBudgetsByPortfolioId(ctx, portfolioID)
	if err != nil {
		return err
	}
	for _, budgetModel := range budgetModels {
		_, err = r.queries.DeleteBudgetAllocations(ctx, budgetModel.BudgetID)
		if err != nil {
			return err
		}
	}
	_, err = r.queries.DeleteBudgetsByPortfolioId(ctx, portfolioID)
	if err != nil {
		return err
	}

	return nil
	// _, err = r.queries.DeletePortfolio(ctx, portfolioID)
	// return err
}

func (r *estimationRepositoryPostgres) GetBudgetManyByPortfolioID(ctx context.Context, portfolioID string) ([]*domain.Budget, error) {
	budgetModels, err := r.queries.FindBudgetsByPortfolioId(ctx, portfolioID)
	if err != nil {
		return nil, err
	}

	budgets := make([]*domain.Budget, len(budgetModels))
	for i, budgetModel := range budgetModels {
		allocations, err := r.queries.FindBudgetAllocations(ctx, budgetModel.BudgetID)
		if err != nil {
			return nil, err
		}

		allocs := make([]domain.BudgetAllocation, len(allocations))
		for j, allocation := range allocations {
			allocs[j] = domain.BudgetAllocation{
				AllocationDate: allocation.AllocationDate.Time,
				Amount:         allocation.Amount,
			}
		}

		props := domain.RestoreBudgetProps{
			BudgetID:          budgetModel.BudgetID,
			PortfolioID:       budgetModel.PortfolioID,
			CostID:            budgetModel.CostID,
			Amount:            budgetModel.Amount,
			BudgetAllocations: allocs,
			CreatedAt:         budgetModel.CreatedAt.Time,
			UpdatedAt:         budgetModel.UpdatedAt.Time,
		}

		budgets[i] = domain.RestoreBudget(props)
		err = budgets[i].Validate()
		if err != nil {
			return nil, err
		}
	}

	return budgets, nil
}
