package testutils

import (
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/google/uuid"
)

type CostFakeBuilder struct {
	CostID              string
	BaselineID          string
	CostType            domain.CostType
	Description         string
	Comment             string
	Amount              float64
	Currency            domain.Currency
	Tax                 float64
	ApplyInflation      bool
	CostAllocationProps []domain.CostAllocationProps
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func NewCostFakeBuilder() *CostFakeBuilder {
	return &CostFakeBuilder{
		CostID:         uuid.New().String(),
		BaselineID:     uuid.New().String(),
		CostType:       domain.CostType(randomdata.StringSample("one_time", "running", "investment")),
		Description:    randomdata.Paragraph(),
		Comment:        randomdata.SillyName(),
		Amount:         100.0,
		Currency:       domain.Currency(randomdata.StringSample("BRL", "USD", "EUR")),
		Tax:            randomdata.Decimal(0.00, 45.00),
		ApplyInflation: randomdata.Boolean(),
		CostAllocationProps: []domain.CostAllocationProps{
			{Year: 2020, Month: time.January, Amount: 60.00},
			{Year: 2020, Month: time.August, Amount: 40.00},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (b *CostFakeBuilder) WithCostID(costID string) *CostFakeBuilder {
	b.CostID = costID
	return b
}

func (b *CostFakeBuilder) WithBaselineID(baselineID string) *CostFakeBuilder {
	b.BaselineID = baselineID
	return b
}

func (b *CostFakeBuilder) WithCostType(costType domain.CostType) *CostFakeBuilder {
	b.CostType = costType
	return b
}

func (b *CostFakeBuilder) WithDescription(description string) *CostFakeBuilder {
	b.Description = description
	return b
}

func (b *CostFakeBuilder) WithComment(comment string) *CostFakeBuilder {
	b.Comment = comment
	return b
}

func (b *CostFakeBuilder) WithAmount(amount float64) *CostFakeBuilder {
	b.Amount = amount
	return b
}

func (b *CostFakeBuilder) WithCurrency(currency domain.Currency) *CostFakeBuilder {
	b.Currency = currency
	return b
}

func (b *CostFakeBuilder) WithTax(tax float64) *CostFakeBuilder {
	b.Tax = tax
	return b
}

func (b *CostFakeBuilder) WithApplyInflation(applyInflation bool) *CostFakeBuilder {
	b.ApplyInflation = applyInflation
	return b
}

func (b *CostFakeBuilder) WithCostAllocationProps(allocations []domain.CostAllocationProps) *CostFakeBuilder {
	b.CostAllocationProps = allocations
	return b
}

func (b *CostFakeBuilder) WithCreatedAt(createdAt time.Time) *CostFakeBuilder {
	b.CreatedAt = createdAt
	return b
}

func (b *CostFakeBuilder) WithUpdatedAt(updatedAt time.Time) *CostFakeBuilder {
	b.UpdatedAt = updatedAt
	return b
}

func (b *CostFakeBuilder) Build() *domain.Cost {
	allocations := make([]domain.CostAllocation, len(b.CostAllocationProps))

	for i, v := range b.CostAllocationProps {
		allocations[i] = newCostAllocation(v.Year, v.Month, v.Amount)
	}

	if len(b.Description) > 255 {
		b.Description = b.Description[:255]
	}

	props := domain.RestoreCostProps{
		CostID:          b.CostID,
		BaselineID:      b.BaselineID,
		CostType:        b.CostType,
		Description:     b.Description,
		Comment:         b.Comment,
		Amount:          b.Amount,
		Currency:        b.Currency,
		Tax:             b.Tax,
		ApplyInflation:  b.ApplyInflation,
		CostAllocations: allocations,
	}

	cost := domain.RestoreCost(props)
	err := cost.Validate()
	if err != nil {
		panic(err)
	}

	return cost
}

func newCostAllocation(year int, month time.Month, amount float64) domain.CostAllocation {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return domain.CostAllocation{
		AllocationDate: date,
		Amount:         amount,
	}
}
