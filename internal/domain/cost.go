package domain

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"time"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/google/uuid"
)

type CostType string

const (
	OneTimeCost CostType = "one_time"
	RunningCost CostType = "running"
	Investment  CostType = "investment"
)

type Cost struct {
	CostID          string           `validate:"required,uuid"`
	BaselineID      string           `validate:"required,uuid"`
	CostType        CostType         `validate:"required"`
	Description     string           `validate:"required"`
	Comment         string           `validate:"-"`
	Amount          float64          `validate:"required"`
	Currency        Currency         `validate:"required"`
	Tax             float64          `validate:"gte=0"`
	CostAllocations []CostAllocation `validate:"required"`
	CreatedAt       time.Time        `validate:"-"`
	UpdatedAt       time.Time        `validate:"-"`
}

type RestoreCostProps Cost
type CostAllocationProps struct {
	Year   int
	Month  time.Month
	Amount float64
}

type NewCostProps struct {
	BaselineID      string
	CostType        CostType
	Description     string
	Comment         string
	Amount          float64
	Currency        Currency
	Tax             float64
	CostAllocations []CostAllocationProps
}

var ErrCostDomainValidation = errors.New("cost domain validation failed")

func NewCost(props NewCostProps) *Cost {
	costAllocations := createCostAllocation(props.CostAllocations)
	return &Cost{
		CostID:          uuid.New().String(),
		BaselineID:      props.BaselineID,
		CostType:        props.CostType,
		Description:     props.Description,
		Comment:         props.Comment,
		Amount:          props.Amount,
		Currency:        props.Currency,
		Tax:             props.Tax,
		CostAllocations: costAllocations,
	}
}

func RestoreCost(props RestoreCostProps) *Cost {
	return &Cost{
		CostID:          props.CostID,
		BaselineID:      props.BaselineID,
		CostType:        props.CostType,
		Description:     props.Description,
		Comment:         props.Comment,
		Amount:          props.Amount,
		Currency:        props.Currency,
		Tax:             props.Tax,
		CostAllocations: props.CostAllocations,
		CreatedAt:       props.CreatedAt,
		UpdatedAt:       props.UpdatedAt,
	}
}

func (c *Cost) Validate() error {
	err := common.Validate.Struct(c)
	if err != nil {
		log.Printf("%v\nstruct: %+v\n", err, c)
		return ErrCostDomainValidation
	}

	if c.Amount <= 0 {
		return fmt.Errorf("invalid cost amount %.2f", c.Amount)
	}

	total := 0.
	for _, v := range c.CostAllocations {
		total += v.Amount
	}
	if total != c.Amount {
		return fmt.Errorf("cost allocation total %.2f is not equal to cost amount %.2f", total, c.Amount)
	}

	if c.Tax < 0 {
		return fmt.Errorf("invalid tax %.2f", c.Tax)
	}
	return nil
}

func (c *Cost) GetCostAllocation() []CostAllocation {
	return c.CostAllocations
}

type CostAllocation struct {
	AllocationDate time.Time
	Amount         float64
}

func NewCostAllocation(year int, month time.Month, amount float64) CostAllocation {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return CostAllocation{
		AllocationDate: date,
		Amount:         amount,
	}
}

func (c *Cost) ChangeCostType(costTypeStr *string) {
	if costTypeStr == nil {
		return
	}
	c.CostType = CostType(*costTypeStr)
}

func (c *Cost) ChangeDescription(description *string) {
	if description == nil {
		return
	}
	c.Description = *description
}

func (c *Cost) ChangeComment(comment *string) {
	if comment == nil {
		return
	}
	c.Comment = *comment
}

func (c *Cost) ChangeAmount(amount *float64) {
	if amount == nil {
		return
	}
	c.Amount = *amount
}

func (c *Cost) ChangeCurrency(currencyStr *string) {
	if currencyStr == nil {
		return
	}
	c.Currency = Currency(*currencyStr)
}

func (c *Cost) ChangeTax(tax *float64) {
	if tax == nil {
		return
	}
	c.Tax = *tax
}

func (c *Cost) ChangeCostAllocations(costAllocationProps []CostAllocationProps) {
	costAllocations := createCostAllocation(costAllocationProps)
	c.CostAllocations = costAllocations
}

func createCostAllocation(params []CostAllocationProps) []CostAllocation {
	costAllocations := make([]CostAllocation, len(params))

	for i, v := range params {
		costAllocations[i] = NewCostAllocation(v.Year, v.Month, v.Amount)
	}

	slices.SortStableFunc(costAllocations, func(a, b CostAllocation) int {
		return a.AllocationDate.Compare(b.AllocationDate)
	})

	return costAllocations
}
