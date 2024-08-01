package domain

import (
	"fmt"
	"time"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/google/uuid"
)

type Budget struct {
	BudgetID          string
	PortfolioID       string
	CostID            string
	Amount            float64
	BudgetAllocations []BudgetAllocation
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type RestoreBudgetProps Budget
type NewBudgetAllocationProps struct {
	Year   int
	Month  time.Month
	Amount float64
}

type NewBudgetProps struct {
	PortfolioID       string
	CostID            string
	Amount            float64
	BudgetAllocations []NewBudgetAllocationProps
}

func NewBudget(props NewBudgetProps) *Budget {
	budgetAllocations := createBudgetAllocation(props.BudgetAllocations)
	return &Budget{
		BudgetID:          uuid.New().String(),
		PortfolioID:       props.PortfolioID,
		CostID:            props.CostID,
		Amount:            props.Amount,
		BudgetAllocations: budgetAllocations,
	}
}

func createBudgetAllocation(params []NewBudgetAllocationProps) []BudgetAllocation {
	budgetAllocations := make([]BudgetAllocation, len(params))

	for i, v := range params {
		budgetAllocations[i] = NewBudgetAllocation(v.Year, v.Month, v.Amount)
	}

	return budgetAllocations
}

func RestoreBudget(props RestoreBudgetProps) *Budget {
	return &Budget{
		BudgetID:          props.BudgetID,
		PortfolioID:       props.PortfolioID,
		CostID:            props.CostID,
		Amount:            props.Amount,
		BudgetAllocations: props.BudgetAllocations,
		CreatedAt:         props.CreatedAt,
		UpdatedAt:         props.UpdatedAt,
	}
}

func (b *Budget) Validate() error {
	if b.Amount <= 0 {
		return common.NewDomainValidationError(fmt.Errorf("invalid budget amount %.2f", b.Amount))
	}

	total := 0.
	for _, v := range b.BudgetAllocations {
		total += v.Amount
	}
	total = common.RoundToTwoDecimals(total)
	if total != b.Amount {
		return common.NewDomainValidationError(fmt.Errorf("budget allocation total %.2f is not equal to budget amount %.2f", total, b.Amount))
	}

	return nil
}

func (b *Budget) GetBudgetAllocation() []BudgetAllocation {
	return b.BudgetAllocations
}

type BudgetAllocation struct {
	AllocationDate time.Time
	Amount         float64
}

func NewBudgetAllocation(year int, month time.Month, amount float64) BudgetAllocation {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return BudgetAllocation{
		AllocationDate: date,
		Amount:         amount,
	}
}
