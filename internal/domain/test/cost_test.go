package domain_test

import (
	"testing"
	"time"

	domain "github.com/celsopires1999/estimation/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCost(t *testing.T) {
	t.Run("should create cost with valid values", func(t *testing.T) {
		props := domain.NewCostProps{
			ProjectID:   "2eb1f2e6-bceb-45b2-a938-3d76b18b020f",
			CostType:    domain.OneTimeCost,
			Description: "Mão de obra do PMO",
			Comment:     "estimativa do Ferraz",
			Amount:      100.,
			Currency:    domain.EUR,
			Installments: []domain.NewInstallmentProps{
				{Year: 2020, Month: time.January, Amount: 60.},
				{Year: 2020, Month: time.August, Amount: 40.},
			},
		}

		cost := domain.NewCost(props)
		assert.NotEmpty(t, cost.CostID)
		_, err := uuid.Parse(cost.CostID)
		assert.NoError(t, err)
		assert.Equal(t, props.ProjectID, cost.ProjectID)
		assert.Equal(t, props.CostType, cost.CostType)
		assert.Equal(t, props.Description, cost.Description)
		assert.Equal(t, props.Comment, cost.Comment)
		assert.Equal(t, props.Amount, cost.Amount)
		assert.Equal(t, props.Currency, cost.Currency)
		assert.Empty(t, cost.CreatedAt)
		assert.Empty(t, cost.UpdatedAt)
	})

	t.Run("should add months", func(t *testing.T) {
		cost := arrangeValidCost()

		cost.AddMonths(5)

		assert.Equal(t, 2, len(cost.Installments))
		assert.Equal(t, time.Date(2020, time.June, 1, 0, 0, 0, 0, time.UTC), cost.Installments[0].PaymentDate)
		assert.Equal(t, 60., cost.Installments[0].Amount)
		assert.Equal(t, time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC), cost.Installments[1].PaymentDate)
		assert.Equal(t, 40., cost.Installments[1].Amount)
	})

	t.Run("should return error when installments is greater than cost amount", func(t *testing.T) {
		cost := arrangeInValidCost()

		err := cost.Validate()

		assert.EqualError(t, err, "installment total 110.00 is not equal to cost amount 100.00")
	})
}
