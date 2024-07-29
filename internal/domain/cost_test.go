package domain_test

import (
	"testing"
	"time"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUnitCost(t *testing.T) {
	t.Run("should create cost with valid values", func(t *testing.T) {
		props := domain.NewCostProps{
			BaselineID:  "2eb1f2e6-bceb-45b2-a938-3d76b18b020f",
			CostType:    domain.OneTimeCost,
			Description: "Mão de obra do PMO",
			Comment:     "estimativa do Ferraz",
			Amount:      100.00,
			Tax:         0.00,
			Currency:    domain.EUR,
			CostAllocations: []domain.CostAllocationProps{
				{Year: 2020, Month: time.January, Amount: 60.},
				{Year: 2020, Month: time.August, Amount: 40.},
			},
		}

		cost := domain.NewCost(props)
		err := cost.Validate()
		assert.NoError(t, err)
		assert.NotEmpty(t, cost.CostID)
		_, err = uuid.Parse(cost.CostID)
		assert.NoError(t, err)
		assert.Equal(t, props.BaselineID, cost.BaselineID)
		assert.Equal(t, props.CostType, cost.CostType)
		assert.Equal(t, props.Description, cost.Description)
		assert.Equal(t, props.Comment, cost.Comment)
		assert.Equal(t, props.Amount, cost.Amount)
		assert.Equal(t, props.Currency, cost.Currency)
		assert.Empty(t, cost.CreatedAt)
		assert.Empty(t, cost.UpdatedAt)
	})

	t.Run("should return error when cost allocation is greater than cost amount", func(t *testing.T) {
		cost := arrangeInValidCost()

		err := cost.Validate()

		assert.EqualError(t, err, "cost allocation total 110.00 is not equal to cost amount 100.00")
	})
}

func arrangeInValidCost() *domain.Cost {
	props := domain.NewCostProps{
		BaselineID:  "2eb1f2e6-bceb-45b2-a938-3d76b18b020f",
		CostType:    domain.OneTimeCost,
		Description: "Mão de obra do PMO",
		Comment:     "estimativa do Ferraz",
		Amount:      100.00,
		Currency:    domain.EUR,
		Tax:         45.00,
		CostAllocations: []domain.CostAllocationProps{
			{Year: 2020, Month: time.January, Amount: 60.00},
			{Year: 2020, Month: time.August, Amount: 50.00},
		},
	}

	return domain.NewCost(props)
}
