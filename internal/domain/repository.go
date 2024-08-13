package domain

import "context"

type EstimationRepository interface {
	UserRepository
	BaselineRepository
	CostRepository
	CompetenceRepository
	EffortRepository
	PlanRepository
	PortfolioRepository
	BudgetRepository
	WorkloadRepository
	Validator
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, userID string) (*User, error)
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, userID string) error
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

type CompetenceRepository interface {
	CreateCompetence(ctx context.Context, competence *Competence) error
	GetCompetence(ctx context.Context, competenceID string) (*Competence, error)
	UpdateCompetence(ctx context.Context, competence *Competence) error
	DeleteCompetence(ctx context.Context, competenceID string) error
}

type EffortRepository interface {
	CreateEffort(ctx context.Context, effort *Effort) error
	CreateEffortMany(ctx context.Context, efforts []*Effort) error
	GetEffort(ctx context.Context, effortID string) (*Effort, error)
	UpdateEffort(ctx context.Context, effort *Effort) error
	DeleteEffort(ctx context.Context, effortID string) error
	GetEffortManyByBaselineID(ctx context.Context, baselineID string) ([]*Effort, error)
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

type WorkloadRepository interface {
	CreateWorkload(ctx context.Context, workload *Workload) error
	CreateWorkloadMany(ctx context.Context, workloads []*Workload) error
	GetWorkload(ctx context.Context, workloadID string) (*Workload, error)
	UpdateWorkload(ctx context.Context, workload *Workload) error
	DeleteWorkload(ctx context.Context, workloadID string) error
	DeleteWorkloadsByPortfolioID(ctx context.Context, portfolioID string) error
	GetWorkloadManyByPortfolioID(ctx context.Context, portfolioID string) ([]*Workload, error)
}

type Validator interface {
	ValidatePlan(ctx context.Context, planID string) error
}
