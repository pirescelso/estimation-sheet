package domain

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/google/uuid"
)

type Plan struct {
	PlanID      string      `validate:"required,uuid"`
	Code        string      `validate:"required,max=10"`
	Name        string      `validate:"required,max=50"`
	Assumptions Assumptions `validate:"required,dive"`
	CreatedAt   time.Time   `validate:"-"`
	UpdatedAt   time.Time   `validate:"-"`
}

type Assumptions []Assumption

type Assumption struct {
	Year       int                  `json:"year" validate:"required"`
	Inflation  float64              `json:"inflation" validate:"gte=0,twodecimals"`
	Currencies []CurrencyAssumption `json:"currencies" validate:"required,dive"`
}

type CurrencyAssumption struct {
	Currency Currency `json:"currency" validate:"required,oneof=USD EUR"`
	Exchange float64  `json:"exchange" validate:"gte=0,twodecimals"`
}

type RestorePlanProps Plan

var (
	ErrPlanValidation             = errors.New("plan domain validation failed")
	ErrAssumptionConsecutiveYears = errors.New("assumptions must have consecutive years")
	ErrAssumptionCurrencyEUR      = errors.New("one EUR currency per year is required")
	ErrAssumptionCurrencyUSD      = errors.New("one USD currency per year is required")
)

func NewPlan(
	code string,
	name string,
	assumptions Assumptions,
) *Plan {
	plan := &Plan{
		PlanID:      uuid.NewString(),
		Code:        code,
		Name:        name,
		Assumptions: assumptions,
	}

	plan.sortAssumptions()
	return plan
}

func RestorePlan(props RestorePlanProps) *Plan {
	plan := &Plan{
		PlanID:      props.PlanID,
		Code:        props.Code,
		Name:        props.Name,
		Assumptions: props.Assumptions,
		CreatedAt:   props.CreatedAt,
		UpdatedAt:   props.UpdatedAt,
	}
	plan.sortAssumptions()
	return plan
}

func (p *Plan) ChangeCode(code string) {
	p.Code = code
}

func (p *Plan) ChangeName(name string) {
	p.Name = name
}

func (p *Plan) ChangeAssumptions(assumptions Assumptions) {
	p.Assumptions = assumptions
	p.sortAssumptions()
}

func (p *Plan) Validate() error {
	err := common.Validate.Struct(p)
	if err != nil {
		return common.NewDomainValidationError(ErrPlanValidation)
	}

	err = p.validateAssumptions()
	if err != nil {
		return common.NewDomainValidationError(fmt.Errorf("baseline domain validation failed: %w", err))
	}
	return nil
}

func (p *Plan) validateAssumptions() error {
	var previousYear int

	for _, assumption := range p.Assumptions {
		if previousYear > 0 && assumption.Year != previousYear+1 {
			return common.NewDomainValidationError(ErrAssumptionConsecutiveYears)
		}
		previousYear = assumption.Year
		var numEUR int
		var numUSD int
		for _, currencyAssumption := range assumption.Currencies {
			if currencyAssumption.Currency.IsEUR() {
				numEUR++
			}
			if currencyAssumption.Currency.IsUSD() {
				numUSD++
			}
		}
		if numEUR != 1 {
			return common.NewDomainValidationError(ErrAssumptionCurrencyEUR)
		}
		if numUSD != 1 {
			return common.NewDomainValidationError(ErrAssumptionCurrencyUSD)
		}
	}
	return nil
}

func (p *Plan) sortAssumptions() {
	slices.SortStableFunc(p.Assumptions, func(a, b Assumption) int {
		return cmp.Compare(a.Year, b.Year)
	})
}

func (p *Plan) GetInflation() *inflation {
	factors := make([]factor, len(p.Assumptions))

	for i, assumption := range p.Assumptions {
		factors[i] = factor{
			year:      assumption.Year,
			inflation: assumption.Inflation,
		}
	}
	return NewInflation(factors)
}

func (p *Plan) GetExchange() *exchange {
	rates := make([]rate, 0)
	for _, assumption := range p.Assumptions {
		for i := range assumption.Currencies {
			rate := rate{
				year:     assumption.Year,
				currency: assumption.Currencies[i].Currency,
				rate:     assumption.Currencies[i].Exchange,
			}
			rates = append(rates, rate)
		}
	}

	return NewExchange(rates)
}
