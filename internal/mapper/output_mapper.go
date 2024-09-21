package mapper

import (
	"encoding/json"
	"time"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
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

func UserOutputFromDomain(user domain.User) UserOutput {
	return UserOutput{
		UserID:    user.UserID,
		Email:     user.Email,
		UserName:  user.UserName,
		Name:      user.Name,
		UserType:  user.UserType.String(),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserOutputFromDb(user db.User) UserOutput {
	return UserOutput{
		UserID:    user.UserID,
		Email:     user.Email,
		UserName:  user.UserName,
		Name:      user.Name,
		UserType:  user.UserType,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}
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

func PlanOutputFromDomain(plan domain.Plan) PlanOutput {
	return PlanOutput{
		PlanID:      plan.PlanID,
		Code:        plan.Code,
		Name:        plan.Name,
		Assumptions: plan.Assumptions,
		CreatedAt:   plan.CreatedAt,
		UpdatedAt:   plan.UpdatedAt,
	}
}

func PlanOutputFromDb(plan db.Plan) PlanOutput {
	return PlanOutput{
		PlanID:      plan.PlanID,
		Code:        plan.Code,
		Name:        plan.Name,
		Assumptions: plan.Assumptions,
		CreatedAt:   plan.CreatedAt.Time,
		UpdatedAt:   plan.UpdatedAt.Time,
	}
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

func BaselineOutputFromDomain(b domain.Baseline) BaselineOutput {
	return BaselineOutput{
		BaselineID:  b.BaselineID,
		Code:        b.Code,
		Review:      b.Review,
		Title:       b.Title,
		Description: b.Description,
		StartDate:   b.StartDate,
		Duration:    b.Duration,
		ManagerID:   b.ManagerID,
		EstimatorID: b.EstimatorID,
		CreatedAt:   b.CreatedAt,
	}
}

func BaselineOutputFromDb(b db.BaselineRow) BaselineOutput {
	return BaselineOutput{
		BaselineID:  b.BaselineID,
		Code:        b.Code,
		Review:      b.Review,
		Title:       b.Title,
		Description: b.Description.String,
		StartDate:   b.StartDate.Time,
		Duration:    b.Duration,
		ManagerID:   b.ManagerID,
		Mananger:    b.Manager,
		EstimatorID: b.EstimatorID,
		Estimator:   b.Estimator,
		CreatedAt:   b.CreatedAt.Time,
		UpdatedAt:   b.UpdatedAt.Time,
	}
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
	ApplyInflation  bool                   `json:"apply_inflation"`
	CostAllocations []costAllocationOutput `json:"cost_allocations"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
}

func CostOutputFromDomain(cost domain.Cost) CostOutput {
	allocs := make([]costAllocationOutput, len(cost.CostAllocations))

	for i := range cost.CostAllocations {
		allocs[i] = costAllocationOutput{
			Year:   cost.CostAllocations[i].AllocationDate.Year(),
			Month:  int(cost.CostAllocations[i].AllocationDate.Month()),
			Amount: cost.CostAllocations[i].Amount,
		}
	}

	return CostOutput{
		CostID:          cost.CostID,
		BaselineID:      cost.BaselineID,
		CostType:        cost.CostType.String(),
		Description:     cost.Description,
		Comment:         cost.Comment,
		Amount:          cost.Amount,
		Currency:        cost.Currency.String(),
		Tax:             cost.Tax,
		ApplyInflation:  cost.ApplyInflation,
		CostAllocations: allocs,
		CreatedAt:       cost.CreatedAt,
	}
}

type costAllocationOutput struct {
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

type CompetenceOutput struct {
	CompetenceID string    `json:"competence_id"`
	Code         string    `json:"code"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func CompetenceOutputFromDomain(c domain.Competence) CompetenceOutput {
	return CompetenceOutput{
		CompetenceID: c.CompetenceID,
		Code:         c.Code,
		Name:         c.Name,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}
}

func CompetenceOutputFromDb(c db.Competence) CompetenceOutput {
	return CompetenceOutput{
		CompetenceID: c.CompetenceID,
		Code:         c.Code,
		Name:         c.Name,
		CreatedAt:    c.CreatedAt.Time,
		UpdatedAt:    c.UpdatedAt.Time,
	}
}

func (o CompetenceOutput) MarshalJSON() ([]byte, error) {
	type Dup CompetenceOutput

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

type EffortOutput struct {
	EffortID          string                   `json:"effort_id"`
	BaselineID        string                   `json:"baseline_id"`
	CompetenceID      string                   `json:"competence_id"`
	Comment           string                   `json:"comment"`
	Hours             int                      `json:"hours"`
	EffortAllocations []effortAllocationOutput `json:"effort_allocations"`
	CreatedAt         time.Time                `json:"created_at"`
	UpdatedAt         time.Time                `json:"updated_at"`
}

func EffortOutputFromDomain(effort domain.Effort) EffortOutput {
	allocs := make([]effortAllocationOutput, len(effort.EffortAllocations))

	for i := range effort.EffortAllocations {
		allocs[i] = effortAllocationOutput{
			Year:  effort.EffortAllocations[i].AllocationDate.Year(),
			Month: int(effort.EffortAllocations[i].AllocationDate.Month()),
			Hours: int(effort.EffortAllocations[i].Hours),
		}
	}

	return EffortOutput{
		EffortID:          effort.EffortID,
		BaselineID:        effort.BaselineID,
		CompetenceID:      effort.CompetenceID,
		Comment:           effort.Comment,
		Hours:             effort.Hours,
		EffortAllocations: allocs,
		CreatedAt:         effort.CreatedAt,
		UpdatedAt:         effort.UpdatedAt,
	}
}

type effortAllocationOutput struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Hours int `json:"hours"`
}

func (o EffortOutput) MarshalJSON() ([]byte, error) {
	type Dup EffortOutput

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
	PortfolioID string           `json:"portfolio_id"`
	Code        string           `json:"code"`
	Review      int32            `json:"review"`
	PlanCode    string           `json:"plan_code"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	StartDate   time.Time        `json:"start_date"`
	Duration    int32            `json:"duration"`
	Manager     string           `json:"manager"`
	Estimator   string           `json:"estimator"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	Budgets     []BudgetOutput   `json:"budgets,omitempty"`
	Workloads   []WorkloadOutput `json:"workloads,omitempty"`
}

func PortfolioOutputFromDb(p db.PortfolioRow) PortfolioOutput {
	return PortfolioOutput{
		PortfolioID: p.PortfolioID,
		PlanCode:    p.PlanCode,
		Code:        p.Code,
		Review:      p.Review,
		Title:       p.Title,
		Description: p.Description.String,
		StartDate:   p.StartDate.Time,
		Duration:    p.Duration,
		Manager:     p.Manager,
		Estimator:   p.Estimator,
		CreatedAt:   p.CreatedAt.Time,
		UpdatedAt:   p.UpdatedAt.Time,
	}
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
	BudgetID           string                   `json:"budget_id"`
	PortfolioID        string                   `json:"portfolio_id"`
	CostType           string                   `json:"cost_type"`
	Description        string                   `json:"description"`
	Comment            string                   `json:"comment"`
	CostAmount         float64                  `json:"cost_amount"`
	CostCurrency       string                   `json:"cost_currency"`
	CostTax            float64                  `json:"cost_tax"`
	CostApplyInflation bool                     `json:"cost_apply_inflation"`
	Amount             float64                  `json:"amount"`
	BudgetAllocations  []budgetAllocationOutput `json:"budget_allocations"`
	CreatedAt          time.Time                `json:"created_at"`
	UpdatedAt          time.Time                `json:"updated_at"`
}

func BudgetOutputFromDb(budget db.BudgetRow, allocations []db.BudgetAllocation) BudgetOutput {
	allocs := make([]budgetAllocationOutput, len(allocations))
	for i, alloc := range allocations {
		allocs[i] = budgetAllocationOutput{
			Year:   alloc.AllocationDate.Time.Year(),
			Month:  int(alloc.AllocationDate.Time.Month()),
			Amount: alloc.Amount,
		}
	}

	return BudgetOutput{
		BudgetID:           budget.BudgetID,
		PortfolioID:        budget.PortfolioID,
		CostType:           budget.CostType,
		Description:        budget.Description,
		Comment:            budget.Comment.String,
		CostAmount:         budget.CostAmount,
		CostCurrency:       budget.CostCurrency,
		CostTax:            budget.CostTax,
		CostApplyInflation: budget.CostApplyInflation,
		Amount:             budget.Amount,
		BudgetAllocations:  allocs,
		CreatedAt:          budget.CreatedAt.Time,
		UpdatedAt:          budget.UpdatedAt.Time,
	}
}

type budgetAllocationOutput struct {
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

type WorkloadOutput struct {
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

func WorkloadOutputFromDb(workload db.WorkloadRow, allocations []db.WorkloadAllocation) WorkloadOutput {
	allocs := make([]workloadAllocationOutput, len(allocations))
	for i, alloc := range allocations {
		allocs[i] = workloadAllocationOutput{
			Year:  alloc.AllocationDate.Time.Year(),
			Month: int(alloc.AllocationDate.Time.Month()),
			Hours: int(alloc.Hours),
		}
	}

	return WorkloadOutput{
		WorkloadID:          workload.WorkloadID,
		PortfolioID:         workload.PortfolioID,
		CompetenceCode:      workload.CompetenceCode,
		CompetenceName:      workload.CompetenceName,
		Comment:             workload.Comment.String,
		Hours:               int(workload.Hours),
		WorkloadAllocations: allocs,
		CreatedAt:           workload.CreatedAt.Time,
		UpdatedAt:           workload.UpdatedAt.Time,
	}
}

type workloadAllocationOutput struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Hours int `json:"hours"`
}

func (o WorkloadOutput) MarshalJSON() ([]byte, error) {
	type Dup WorkloadOutput

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
