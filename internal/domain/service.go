package domain

import (
	"time"

	"github.com/celsopires1999/estimation/internal/common"
)

type PortfolioService struct {
	planID      string
	baseline    *Baseline
	costs       []*Cost
	inflation   *inflation
	exchange    *exchange
	shiftMonths int
}

func NewPortfolioService(
	planID string,
	baseline *Baseline,
	costs []*Cost,
	inflation *inflation,
	exchange *exchange,
	shiftMonths int,
) *PortfolioService {
	return &PortfolioService{
		planID,
		baseline,
		costs,
		inflation,
		exchange,
		shiftMonths,
	}
}

func (s *PortfolioService) GeneratePortfolio() (*Portfolio, []*Budget, error) {
	startDate := s.baseline.StartDate.AddDate(0, s.shiftMonths, 0)
	portfolio := NewPortfolio(s.baseline.BaselineID, s.planID, startDate)
	err := portfolio.Validate()
	if err != nil {
		return nil, nil, err
	}

	budgets := make([]*Budget, len(s.costs))

	for i, cost := range s.costs {
		budgetProps := NewBudgetProps{}
		budgetProps.PortfolioID = portfolio.PortfolioID
		budgetProps.CostID = cost.CostID
		budgetProps.Amount = 0.0
		for _, costAllocation := range cost.CostAllocations {
			newAllocationDate := costAllocation.AllocationDate.AddDate(0, s.shiftMonths, 0)
			amount, err := s.calculateBudgetAllocation(cost, costAllocation, newAllocationDate)
			if err != nil {
				return nil, nil, err
			}
			budgetProps.Amount += amount
			budgetProps.BudgetAllocations = append(budgetProps.BudgetAllocations, NewBudgetAllocationProps{
				Year:   newAllocationDate.Year(),
				Month:  newAllocationDate.Month(),
				Amount: amount,
			})
		}

		budgetProps.Amount = common.RoundToTwoDecimals(budgetProps.Amount)
		budgets[i] = NewBudget(budgetProps)
		err := budgets[i].Validate()
		if err != nil {
			return nil, nil, err
		}
	}

	return portfolio, budgets, nil

}

func (s *PortfolioService) calculateBudgetAllocation(cost *Cost, costAllocation CostAllocation, budgetAllocationDate time.Time) (float64, error) {
	if cost.Currency.IsBRL() {
		amount, err := s.inflation.ApplyInflation(s.applyTax(cost, costAllocation), s.baseline.StartDate.Year(), budgetAllocationDate.Year())
		if err != nil {
			return 0.0, err
		}
		return amount, nil
	}

	amount, err := s.exchange.ConvertToBRL(s.applyTax(cost, costAllocation), cost.Currency, budgetAllocationDate.Year())
	if err != nil {
		return 0.0, err
	}

	return amount, nil
}

func (s *PortfolioService) applyTax(cost *Cost, costAllocation CostAllocation) float64 {
	if cost.Tax == 0 {
		return costAllocation.Amount
	}
	return (1 + cost.Tax/100) * costAllocation.Amount
}
