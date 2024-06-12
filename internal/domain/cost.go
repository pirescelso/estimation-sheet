package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CostType string

const (
	OneTimeCost CostType = "one_time"
	RunningCost CostType = "running"
	Investment  CostType = "investment"
)

type Cost struct {
	CostID       string
	ProjectID    string
	CostType     CostType
	Description  string
	Comment      string
	Amount       float64
	Currency     Currency
	Installments []Installment
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type RestoreCostProps Cost
type NewInstallmentProps struct {
	Year   int
	Month  time.Month
	Amount float64
}

type NewCostProps struct {
	ProjectID    string
	CostType     CostType
	Description  string
	Comment      string
	Amount       float64
	Currency     Currency
	Installments []NewInstallmentProps
}

func NewCost(props NewCostProps) *Cost {
	installments := createInstallment(props.Installments)
	return &Cost{
		CostID:       uuid.New().String(),
		ProjectID:    props.ProjectID,
		CostType:     props.CostType,
		Description:  props.Description,
		Comment:      props.Comment,
		Amount:       props.Amount,
		Currency:     props.Currency,
		Installments: installments,
	}
}

func createInstallment(params []NewInstallmentProps) []Installment {
	installments := make([]Installment, len(params))

	for i, v := range params {
		installments[i] = NewInstallment(v.Year, v.Month, v.Amount)
	}

	return installments
}

func RestoreCost(input RestoreCostProps) *Cost {
	return &Cost{
		CostID:       input.CostID,
		ProjectID:    input.ProjectID,
		CostType:     input.CostType,
		Description:  input.Description,
		Comment:      input.Comment,
		Amount:       input.Amount,
		Currency:     input.Currency,
		Installments: input.Installments,
		CreatedAt:    input.CreatedAt,
		UpdatedAt:    input.UpdatedAt,
	}
}

func (c *Cost) Validate() error {
	if c.Amount <= 0 {
		return fmt.Errorf("invalid cost amount %.2f", c.Amount)
	}

	total := 0.
	for _, v := range c.Installments {
		total += v.Amount
	}
	if total != c.Amount {
		return fmt.Errorf("installment total %.2f is not equal to cost amount %.2f", total, c.Amount)
	}

	return nil
}

func (c *Cost) AddMonths(months int) error {

	for j := range c.Installments {
		if err := c.Installments[j].AddMonths(months); err != nil {
			return err
		}
	}
	return nil
}

func (c *Cost) GetInstalmments() []Installment {
	return c.Installments
}

type Installment struct {
	PaymentDate time.Time
	Amount      float64
}

func NewInstallment(year int, month time.Month, amount float64) Installment {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return Installment{
		PaymentDate: date,
		Amount:      amount,
	}
}

func (i *Installment) AddMonths(months int) error {
	i.PaymentDate = i.PaymentDate.AddDate(0, months, 0)
	return nil
}
