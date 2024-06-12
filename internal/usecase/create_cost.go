package usecase

import (
	"context"
	"time"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
)

type CreateCostInputDTO struct {
	ProjectID    string        `json:"project_id" validate:"required"`
	CostType     string        `json:"cost_type" validate:"required,oneof=one_time running investment" errmsg:"Cost type must be one of: one_time, running, investment"`
	Description  string        `json:"description" validate:"required"`
	Comment      string        `json:"comment"`
	Amount       float64       `json:"amount" validate:"required"`
	Currency     string        `json:"currency" validate:"required,oneof=BRL USD EUR"`
	Installments []Installment `json:"installments" validate:"required,dive,required"`
}

type CreateCostOutputDTO struct {
	CostID       string        `json:"cost_id"`
	ProjectID    string        `json:"project_id"`
	CostType     string        `json:"cost_type"`
	Description  string        `json:"description"`
	Comment      string        `json:"comment"`
	Amount       float64       `json:"amount"`
	Currency     string        `json:"currency"`
	Installments []Installment `json:"installments"`
	CreatedAt    time.Time     `json:"created_at"`
}

type Installment struct {
	Year   int     `json:"year" validate:"required"`
	Month  int     `json:"month" validate:"gte=1,lte=12"`
	Amount float64 `json:"amount" validate:"required"`
}

type CreateCostUseCase struct {
	txm *db.TransactionManager
}

func NewCreateCostUseCase(txm *db.TransactionManager) *CreateCostUseCase {
	return &CreateCostUseCase{
		txm: txm,
	}
}

func (uc *CreateCostUseCase) Execute(ctx context.Context, input CreateCostInputDTO) (*CreateCostOutputDTO, error) {
	installments := make([]domain.NewInstallmentProps, len(input.Installments))
	for i, installment := range input.Installments {
		installments[i] = domain.NewInstallmentProps{
			Year:   installment.Year,
			Month:  time.Month(installment.Month),
			Amount: installment.Amount,
		}
	}
	cost := domain.NewCost(domain.NewCostProps{
		ProjectID:    input.ProjectID,
		CostType:     domain.CostType(input.CostType),
		Description:  input.Description,
		Comment:      input.Comment,
		Amount:       input.Amount,
		Currency:     domain.Currency(input.Currency),
		Installments: installments,
	})

	var createdCost *domain.Cost

	err := uc.txm.Do(ctx, func() error {
		costRepo := uc.getCostRepo(ctx)
		err := costRepo.CreateCost(ctx, cost)
		if err != nil {
			return err
		}

		createdCost, err = costRepo.GetCost(ctx, cost.CostID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	outputInstallments := make([]Installment, len(createdCost.Installments))

	for i := range createdCost.Installments {
		outputInstallments[i] = Installment{
			Year:   createdCost.Installments[i].PaymentDate.Year(),
			Month:  int(createdCost.Installments[i].PaymentDate.Month()),
			Amount: createdCost.Installments[i].Amount,
		}
	}

	return &CreateCostOutputDTO{
		CostID:       createdCost.CostID,
		ProjectID:    createdCost.ProjectID,
		CostType:     string(createdCost.CostType),
		Description:  createdCost.Description,
		Comment:      createdCost.Comment,
		Amount:       createdCost.Amount,
		Currency:     string(createdCost.Currency),
		Installments: outputInstallments,
		CreatedAt:    createdCost.CreatedAt,
	}, nil
}

func (uc *CreateCostUseCase) getCostRepo(ctx context.Context) domain.CostRepository {
	repo, err := uc.txm.GetRepository(ctx, "CostRepo")
	if err != nil {
		panic(err)
	}
	return repo.(domain.CostRepository)
}
