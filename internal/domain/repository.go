package domain

import "context"

type EstimationRepository interface {
	BaselineRepository
	CostRepository
	PlanRepository
	PortfolioRepository
	BudgetRepository
	Validator
}

type BaselineRepository interface {
	CreateBaseline(ctx context.Context, baseline *Baseline) error
	GetBaseline(ctx context.Context, baselineID string) (*Baseline, error)
	UpdateBaseline(ctx context.Context, baseline *Baseline) error
	DeleteBaseline(ctx context.Context, baselineID string) error
}

type CostRepository interface {
	CreateCost(ctx context.Context, cost *Cost) error
	CreateCostMany(ctx context.Context, costs []*Cost) error
	GetCost(ctx context.Context, costID string) (*Cost, error)
	UpdateCost(ctx context.Context, cost *Cost) error
	DeleteCost(ctx context.Context, costID string) error
	GetCostManyByBaselineID(ctx context.Context, baselineID string) ([]*Cost, error)
}

type PlanRepository interface {
	CreatePlan(ctx context.Context, plan *Plan) error
	GetPlan(ctx context.Context, planID string) (*Plan, error)
	UpdatePlan(ctx context.Context, plan *Plan) error
	DeletePlan(ctx context.Context, planID string) error
}

type PortfolioRepository interface {
	CreatePortfolio(ctx context.Context, portfolio *Portfolio) error
	GetPortfolio(ctx context.Context, portfolioID string) (*Portfolio, error)
	UpdatePortfolio(ctx context.Context, portfolio *Portfolio) error
	DeletePortfolio(ctx context.Context, portfolioID string) error
	CountPortfoliosByPlanId(ctx context.Context, planID string) (int64, error)
	CountPortfoliosByBaselineId(ctx context.Context, baselineID string) (int64, error)
}

type BudgetRepository interface {
	CreateBudget(ctx context.Context, budget *Budget) error
	CreateBudgetMany(ctx context.Context, budgets []*Budget) error
	GetBudget(ctx context.Context, budgetID string) (*Budget, error)
	UpdateBudget(ctx context.Context, budget *Budget) error
	DeleteBudget(ctx context.Context, budgetID string) error
	DeleteBudgetsByPortfolioID(ctx context.Context, portfolioID string) error
	GetBudgetManyByPortfolioID(ctx context.Context, portfolioID string) ([]*Budget, error)
}

type Validator interface {
	ValidateUser(ctx context.Context, userID string) (isManager bool, err error)
	ValidatePlan(ctx context.Context, planID string) error
}
