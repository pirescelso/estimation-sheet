package domain_test

import (
	"time"

	domain "github.com/celsopires1999/estimation/internal/domain"
)

func arrangeValidCost() *domain.Cost {
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

	return domain.NewCost(props)
}

func arrangeInValidCost() *domain.Cost {
	props := domain.NewCostProps{
		ProjectID:   "2eb1f2e6-bceb-45b2-a938-3d76b18b020f",
		CostType:    domain.OneTimeCost,
		Description: "Mão de obra do PMO",
		Comment:     "estimativa do Ferraz",
		Amount:      100.,
		Currency:    domain.EUR,
		Installments: []domain.NewInstallmentProps{
			{Year: 2020, Month: time.January, Amount: 60.},
			{Year: 2020, Month: time.August, Amount: 50.},
		},
	}

	return domain.NewCost(props)
}
