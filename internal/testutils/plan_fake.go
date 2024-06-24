package testutils

import (
	"time"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/google/uuid"
)

type PlanFakeBuilder struct {
	Code        string
	Name        string
	assumptions domain.Assumptions
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewPlanFakeBuilder() *PlanFakeBuilder {
	assumptions := domain.Assumptions{
		domain.Assumption{
			Year:      2022,
			Inflation: 3.10,
			Currencies: []domain.CurrencyAssumption{{
				Currency: domain.USD,
				Exchange: 4.15,
			}, {
				Currency: domain.EUR,
				Exchange: 5.36,
			}},
		},
		domain.Assumption{
			Year:      2023,
			Inflation: 3.50,
			Currencies: []domain.CurrencyAssumption{{
				Currency: domain.USD,
				Exchange: 4.25,
			}, {
				Currency: domain.EUR,
				Exchange: 5.56,
			}},
		},
		domain.Assumption{
			Year:      2024,
			Inflation: 4.00,
			Currencies: []domain.CurrencyAssumption{{
				Currency: domain.USD,
				Exchange: 4.50,
			}, {
				Currency: domain.EUR,
				Exchange: 5.50,
			}},
		},
		domain.Assumption{
			Year:      2025,
			Inflation: 5.20,
			Currencies: []domain.CurrencyAssumption{{
				Currency: domain.USD,
				Exchange: 5.00,
			}, {
				Currency: domain.EUR,
				Exchange: 6.00,
			}},
		},
		domain.Assumption{
			Year:      2026,
			Inflation: 5.26,
			Currencies: []domain.CurrencyAssumption{{
				Currency: domain.USD,
				Exchange: 5.55,
			}, {
				Currency: domain.EUR,
				Exchange: 6.66,
			}},
		},
		domain.Assumption{
			Year:      2027,
			Inflation: 5.30,
			Currencies: []domain.CurrencyAssumption{{
				Currency: domain.USD,
				Exchange: 5.77,
			}, {
				Currency: domain.EUR,
				Exchange: 6.88,
			}},
		},
	}

	return &PlanFakeBuilder{
		Code:        "BP 2026",
		Name:        "Business Plan 2026",
		assumptions: assumptions,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (b *PlanFakeBuilder) Build() *domain.Plan {
	plan := &domain.Plan{
		PlanID:      uuid.New().String(),
		Code:        b.Code,
		Name:        b.Name,
		Assumptions: b.assumptions,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}

	err := plan.Validate()
	if err != nil {
		panic(err)
	}
	return plan
}
