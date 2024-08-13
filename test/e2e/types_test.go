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
	PortfolioID string         `json:"portfolio_id"`
	Code        string         `json:"code"`
	Review      int32          `json:"review"`
	PlanCode    string         `json:"plan_code"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	StartDate   string         `json:"start_date" layout:"2006-01-02"`
	Duration    int32          `json:"duration"`
	Manager     string         `json:"manager"`
	Estimator   string         `json:"estimator"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Budgets     []budgetOutput `json:"budgets,omitempty"`
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
	Amount            float64                  `json:"amount"`
	BudgetAllocations []budgetAllocationOutput `json:"budget_allocations"`
	CreatedAt         time.Time                `json:"created_at"`
	UpdatedAt         time.Time                `json:"updated_at"`
}

type budgetAllocationOutput struct {
	Year   int     `json:"year"`
	Month  int     `json:"month"`
	Amount float64 `json:"amount"`
}
