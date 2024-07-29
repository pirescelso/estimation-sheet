package mapper

import (
	"encoding/json"
	"time"

	"github.com/celsopires1999/estimation/internal/domain"
)

type UserOutput struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	UserName  string    `json:"user_name"`
	Name      string    `json:"name"`
	UserType  string    `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (o UserOutput) MarshalJSON() ([]byte, error) {
	type Dup UserOutput

	tmp := struct {
		Dup
		CreatedAt *string `json:"created_at"`
		UpdatedAt *string `json:"updated_at"`
	}{
		Dup: (Dup)(o),
	}

	tmp.CreatedAt, tmp.UpdatedAt = fmtRFC3339Time(o.CreatedAt, o.UpdatedAt)

	b, err := json.Marshal(tmp)
	return b, err
}

type PlanOutput struct {
	PlanID      string             `json:"plan_id"`
	Code        string             `json:"code"`
	Name        string             `json:"name"`
	Assumptions domain.Assumptions `json:"assumptions,omitempty"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

func (o PlanOutput) MarshalJSON() ([]byte, error) {
	type Dup PlanOutput

	tmp := struct {
		Dup
		CreatedAt *string `json:"created_at"`
		UpdatedAt *string `json:"updated_at"`
	}{
		Dup: (Dup)(o),
	}

	tmp.CreatedAt, tmp.UpdatedAt = fmtRFC3339Time(o.CreatedAt, o.UpdatedAt)

	b, err := json.Marshal(tmp)
	return b, err
}

type BaselineOutput struct {
	BaselineID  string    `json:"baseline_id"`
	Code        string    `json:"code"`
	Review      int32     `json:"review"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	Duration    int32     `json:"duration"`
	ManagerID   string    `json:"manager_id,omitempty"`
	Mananger    string    `json:"manager,omitempty"`
	EstimatorID string    `json:"estimator_id,omitempty"`
	Estimator   string    `json:"estimator,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (o BaselineOutput) MarshalJSON() ([]byte, error) {
	type Dup BaselineOutput

	tmp := struct {
		Dup
		StartDate string  `json:"start_date"`
		CreatedAt *string `json:"created_at"`
		UpdatedAt *string `json:"updated_at"`
	}{
		Dup:       (Dup)(o),
		StartDate: o.StartDate.Format("2006-01-02"),
	}

	tmp.CreatedAt, tmp.UpdatedAt = fmtRFC3339Time(o.CreatedAt, o.UpdatedAt)

	b, err := json.Marshal(tmp)
	return b, err
}

type CostOutput struct {
	CostID          string                 `json:"cost_id"`
	BaselineID      string                 `json:"baseline_id"`
	CostType        string                 `json:"cost_type"`
	Description     string                 `json:"description"`
	Comment         string                 `json:"comment"`
	Amount          float64                `json:"amount"`
	Currency        string                 `json:"currency"`
	Tax             float64                `json:"tax"`
	CostAllocations []CostAllocationOutput `json:"cost_allocations"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
}

type CostAllocationOutput struct {
	Year   int     `json:"year"`
	Month  int     `json:"month"`
	Amount float64 `json:"amount"`
}

func (o CostOutput) MarshalJSON() ([]byte, error) {
	type Dup CostOutput

	tmp := struct {
		Dup
		CreatedAt *string `json:"created_at"`
		UpdatedAt *string `json:"updated_at"`
	}{
		Dup: (Dup)(o),
	}

	tmp.CreatedAt, tmp.UpdatedAt = fmtRFC3339Time(o.CreatedAt, o.UpdatedAt)

	b, err := json.Marshal(tmp)
	return b, err
}

type PortfolioOutput struct {
	PortfolioID string         `json:"portfolio_id"`
	Code        string         `json:"code"`
	Review      int32          `json:"review"`
	PlanCode    string         `json:"plan_code"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	StartDate   time.Time      `json:"start_date"`
	Duration    int32          `json:"duration"`
	Manager     string         `json:"manager"`
	Estimator   string         `json:"estimator"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Budgets     []BudgetOutput `json:"budgets,omitempty"`
}

func (o PortfolioOutput) MarshalJSON() ([]byte, error) {
	type Dup PortfolioOutput

	tmp := struct {
		Dup
		StartDate string  `json:"start_date"`
		CreatedAt *string `json:"created_at"`
		UpdatedAt *string `json:"updated_at"`
	}{
		Dup:       (Dup)(o),
		StartDate: o.StartDate.Format("2006-01-02"),
	}

	tmp.CreatedAt, tmp.UpdatedAt = fmtRFC3339Time(o.CreatedAt, o.UpdatedAt)

	b, err := json.Marshal(tmp)
	return b, err
}

type BudgetOutput struct {
	BudgetID          string                   `json:"budget_id"`
	PortfolioID       string                   `json:"portfolio_id"`
	CostType          string                   `json:"cost_type"`
	Description       string                   `json:"description"`
	Comment           string                   `json:"comment"`
	CostAmount        float64                  `json:"cost_amount"`
	CostCurrency      string                   `json:"cost_currency"`
	CostTax           float64                  `json:"cost_tax"`
	Amount            float64                  `json:"amount"`
	BudgetAllocations []BudgetAllocationOutput `json:"budget_allocations"`
	CreatedAt         time.Time                `json:"created_at"`
	UpdatedAt         time.Time                `json:"updated_at"`
}

type BudgetAllocationOutput struct {
	Year   int     `json:"year"`
	Month  int     `json:"month"`
	Amount float64 `json:"amount"`
}

func (o BudgetOutput) MarshalJSON() ([]byte, error) {
	type Dup BudgetOutput

	tmp := struct {
		Dup
		CreatedAt *string `json:"created_at"`
		UpdatedAt *string `json:"updated_at"`
	}{
		Dup: (Dup)(o),
	}

	tmp.CreatedAt, tmp.UpdatedAt = fmtRFC3339Time(o.CreatedAt, o.UpdatedAt)

	b, err := json.Marshal(tmp)
	return b, err
}

func fmtRFC3339Time(createdAt, updatedAt time.Time) (createdAtStr *string, updatedAtStr *string) {
	if !createdAt.IsZero() {
		tmp := createdAt.Format(time.RFC3339)
		createdAtStr = &tmp
	}

	if !updatedAt.IsZero() {
		tmp := updatedAt.Format(time.RFC3339)
		updatedAtStr = &tmp
	}
	return
}
