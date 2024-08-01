package domain

import (
	"fmt"
	"time"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Portfolio struct {
	PortfolioID string    `validate:"required"`
	BaselineID  string    `validate:"required"`
	PlanID      string    `validate:"required"`
	StartDate   time.Time `validate:"required"`
	CreatedAt   time.Time `validate:"-"`
	UpdatedAt   time.Time `validate:"-"`
}

type RestorePortfolioProps Portfolio

func NewPortfolio(baselineID string, planID string, startDate time.Time) *Portfolio {
	return &Portfolio{
		PortfolioID: uuid.New().String(),
		BaselineID:  baselineID,
		PlanID:      planID,
		StartDate:   startDate,
	}
}

func RestorePortfolio(props RestorePortfolioProps) *Portfolio {
	return &Portfolio{
		PortfolioID: props.PortfolioID,
		BaselineID:  props.BaselineID,
		PlanID:      props.PlanID,
		StartDate:   props.StartDate,
		CreatedAt:   props.CreatedAt,
		UpdatedAt:   props.UpdatedAt,
	}
}

func (p *Portfolio) Validate() error {
	validate := validator.New()
	err := validate.Struct(p)
	if err != nil {
		return common.NewDomainValidationError(fmt.Errorf("portfolio domain validation failed: %w", err))
	}
	return nil
}
