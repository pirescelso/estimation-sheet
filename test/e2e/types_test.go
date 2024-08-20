package e2e_test

import (
	"time"

	"github.com/celsopires1999/estimation/internal/domain"
)

type userInput struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Name     string `json:"name"`
	UserType string `json:"user_type"`
}

type userOutput struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	UserName  string    `json:"user_name"`
	Name      string    `json:"name"`
	UserType  string    `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type competenceInput struct {
	Code string `json:"code" validate:"required,max=20"`
	Name string `json:"name" validate:"required,max=50"`
}

type competenceOutput struct {
	CompetenceID string    `json:"competence_id"`
	Code         string    `json:"code"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type baselineInput struct {
	Code        string `json:"code"`
	Review      int    `json:"review"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartMonth  int    `json:"start_month"`
	StartYear   int    `json:"start_year"`
	Duration    int    `json:"duration"`
	ManagerID   string `json:"manager_id"`
	EstimatorID string `json:"estimator_id"`
}

type baselineOutput struct {
	BaselineID  string    `json:"baseline_id"`
	Code        string    `json:"code"`
	Review      int       `json:"review"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   string    `json:"start_date" layout:"2006-01-02"`
	Duration    int       `json:"duration"`
	ManagerID   string    `json:"manager_id"`
	Mananger    string    `json:"manager"`
	EstimatorID string    `json:"estimator_id"`
	Estimator   string    `json:"estimator"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type costInput struct {
	BaselineID      string                `json:"baseline_id"`
	CostType        string                `json:"cost_type"`
	Description     string                `json:"description"`
	Comment         string                `json:"comment"`
	Amount          float64               `json:"amount"`
	Currency        string                `json:"currency"`
	Tax             float64               `json:"tax"`
	ApplyInflation  bool                  `json:"apply_inflation"`
	CostAllocations []costAllocationInput `json:"cost_allocations"`
}

type costAllocationInput struct {
	Year   int     `json:"year"`
	Month  int     `json:"month"`
	Amount float64 `json:"amount"`
}

type effortInput struct {
	BaselineID        string                  `json:"baseline_id" validate:"required,uuid4"`
	CompetenceID      string                  `json:"competence_id" validate:"required,uuid4"`
	Comment           string                  `json:"comment"`
	Hours             int                     `json:"hours" validate:"required,gte=1,lte=160_000"`
	EffortAllocations []effortAllocationInput `json:"effort_allocations" validate:"required,dive"`
}

type effortAllocationInput struct {
	Year  int `json:"year" validate:"required"`
	Month int `json:"month" validate:"gte=1,lte=12"`
	Hours int `json:"hours" validate:"required,gte=1,lte=8_000"`
}

type effortOutput struct {
	EffortID          string                   `json:"effort_id"`
	BaselineID        string                   `json:"baseline_id"`
	CompetenceID      string                   `json:"competence_id"`
	Comment           string                   `json:"comment"`
	Hours             int                      `json:"hours"`
	EffortAllocations []effortAllocationOutput `json:"effort_allocations"`
	CreatedAt         time.Time                `json:"created_at"`
	UpdatedAt         time.Time                `json:"updated_at"`
}

type effortAllocationOutput struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Hours int `json:"hours"`
}

type costOutput struct {
	CostID          string                 `json:"cost_id"`
	BaselineID      string                 `json:"baseline_id"`
	CostType        string                 `json:"cost_type"`
	Description     string                 `json:"description"`
	Comment         string                 `json:"comment"`
	Amount          float64                `json:"amount"`
	Currency        string                 `json:"currency"`
	Tax             float64                `json:"tax"`
	ApplyInflation  bool                   `json:"apply_inflation"`
	CostAllocations []costAllocationOutput `json:"cost_allocations"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
}

type costAllocationOutput struct {
	Year   int     `json:"year"`
	Month  int     `json:"month"`
	Amount float64 `json:"amount"`
}

type planInput struct {
	Code        string             `json:"code"`
	Name        string             `json:"name"`
	Assumptions domain.Assumptions `json:"assumptions"`
}

type planOutput struct {
	PlanID      string             `json:"plan_id"`
	Code        string             `json:"code"`
	Name        string             `json:"name"`
	Assumptions domain.Assumptions `json:"assumptions,omitempty"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type portfolioInput struct {
	BaselineID  string `json:"baseline_id"`
	PlanID      string `json:"plan_id"`
	ShiftMonths int    `json:"shift_months"`
}

type portfolioIDOutput struct {
	PortfolioID string `json:"portfolio_id"`
}

type portfolioOutput struct {
	PortfolioID string           `json:"portfolio_id"`
	Code        string           `json:"code"`
	Review      int32            `json:"review"`
	PlanCode    string           `json:"plan_code"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	StartDate   string           `json:"start_date" layout:"2006-01-02"`
	Duration    int32            `json:"duration"`
	Manager     string           `json:"manager"`
	Estimator   string           `json:"estimator"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	Budgets     []budgetOutput   `json:"budgets,omitempty"`
	Workloads   []workloadOutput `json:"workloads,omitempty"`
}

type budgetOutput struct {
	BudgetID          string                   `json:"budget_id"`
	PortfolioID       string                   `json:"portfolio_id"`
	CostType          string                   `json:"cost_type"`
	Description       string                   `json:"description"`
	Comment           string                   `json:"comment"`
	CostAmount        float64                  `json:"cost_amount"`
	CostCurrency      string                   `json:"cost_currency"`
	CostTax           float64                  `json:"cost_tax"`
	ApplyInflation    bool                     `json:"apply_inflation"`
	Amount            float64                  `json:"amount"`
	BudgetAllocations []budgetAllocationOutput `json:"budget_allocations"`
	CreatedAt         time.Time                `json:"created_at"`
	UpdatedAt         time.Time                `json:"updated_at"`
}

type budgetAllocationOutput struct {
	Year   int     `json:"year"`
	Amount float64 `json:"amount"`
}

type workloadOutput struct {
	WorkloadID          string                     `json:"workload_id"`
	PortfolioID         string                     `json:"portfolio_id"`
	CompetenceCode      string                     `json:"competence_code"`
	CompetenceName      string                     `json:"competence_name"`
	Comment             string                     `json:"comment"`
	Hours               int                        `json:"hours"`
	WorkloadAllocations []workloadAllocationOutput `json:"workload_allocations"`
	CreatedAt           time.Time                  `json:"created_at"`
	UpdatedAt           time.Time                  `json:"updated_at"`
}

type workloadAllocationOutput struct {
	Year  int `json:"year"`
	Hours int `json:"hours"`
}
