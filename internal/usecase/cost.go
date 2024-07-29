package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/celsopires1999/estimation/internal/mapper"
)

var (
	ErrCostAllocationDateIsInvalid = errors.New("cost allocation date is invalid")
)

type CreateCostUseCase struct {
	txm db.TransactionManagerInterface
}

type CreateCostInputDTO struct {
	BaselineID      string                `json:"baseline_id" validate:"required"`
	CostType        string                `json:"cost_type" validate:"required,oneof=one_time running investment" errmsg:"Cost type must be one of: one_time, running, investment"`
	Description     string                `json:"description" validate:"required"`
	Comment         string                `json:"comment"`
	Amount          float64               `json:"amount" validate:"required,twodecimals"`
	Currency        string                `json:"currency" validate:"required,oneof=BRL USD EUR"`
	Tax             float64               `json:"tax" validate:"gte=0,twodecimals"`
	CostAllocations []CostAllocationInput `json:"cost_allocations" validate:"required,dive"`
}

type CreateCostOutputDTO struct {
	mapper.CostOutput
}

type CostAllocationInput struct {
	Year   int     `json:"year" validate:"required"`
	Month  int     `json:"month" validate:"gte=1,lte=12"`
	Amount float64 `json:"amount" validate:"required"`
}

func NewCreateCostUseCase(txm db.TransactionManagerInterface) *CreateCostUseCase {
	return &CreateCostUseCase{txm}
}

func (uc *CreateCostUseCase) Execute(ctx context.Context, input CreateCostInputDTO) (*CreateCostOutputDTO, error) {
	var createdCost *domain.Cost

	err := uc.txm.Do(ctx, func(ctx context.Context, tx db.TransactionInterface) error {
		repository, err := db.GetAs[domain.EstimationRepository](tx, "EstimationRepository")
		if err != nil {
			return err
		}

		baseline, err := repository.GetBaseline(ctx, input.BaselineID)
		if err != nil {
			return err
		}

		costAllocations := make([]domain.CostAllocationProps, len(input.CostAllocations))
		for i, allocation := range input.CostAllocations {
			costAllocations[i] = domain.CostAllocationProps{
				Year:   allocation.Year,
				Month:  time.Month(allocation.Month),
				Amount: allocation.Amount,
			}
		}
		cost := domain.NewCost(domain.NewCostProps{
			BaselineID:      input.BaselineID,
			CostType:        domain.CostType(input.CostType),
			Description:     input.Description,
			Comment:         input.Comment,
			Amount:          input.Amount,
			Currency:        domain.Currency(input.Currency),
			Tax:             input.Tax,
			CostAllocations: costAllocations,
		})

		if err := cost.Validate(); err != nil {
			return err
		}

		for _, a := range cost.CostAllocations {
			if baseline.StartDate.After(a.AllocationDate) {
				return ErrCostAllocationDateIsInvalid
			}
		}

		err = repository.CreateCost(ctx, cost)
		if err != nil {
			return err
		}

		createdCost, err = repository.GetCost(ctx, cost.CostID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	outputCostAllocations := make([]mapper.CostAllocationOutput, len(createdCost.CostAllocations))

	for i := range createdCost.CostAllocations {
		outputCostAllocations[i] = mapper.CostAllocationOutput{
			Year:   createdCost.CostAllocations[i].AllocationDate.Year(),
			Month:  int(createdCost.CostAllocations[i].AllocationDate.Month()),
			Amount: createdCost.CostAllocations[i].Amount,
		}
	}

	output := mapper.CostOutput{
		CostID:          createdCost.CostID,
		BaselineID:      createdCost.BaselineID,
		CostType:        string(createdCost.CostType),
		Description:     createdCost.Description,
		Comment:         createdCost.Comment,
		Amount:          createdCost.Amount,
		Currency:        string(createdCost.Currency),
		Tax:             createdCost.Tax,
		CostAllocations: outputCostAllocations,
		CreatedAt:       createdCost.CreatedAt,
	}

	return &CreateCostOutputDTO{output}, nil
}

// UpdateCostUseCase represents the use case for updating a cost
type UpdateCostUseCase struct {
	txm db.TransactionManagerInterface
}

type UpdateCostInputDTO struct {
	CostID          string                 `json:"cost_id" validate:"required"`
	CostType        *string                `json:"cost_type" validate:"omitempty,required,oneof=one_time running investment" errmsg:"Cost type must be one of: one_time, running, investment"`
	Description     *string                `json:"description" validate:"omitempty,required"`
	Comment         *string                `json:"comment"`
	Amount          *float64               `json:"amount" validate:"omitempty,required,twodecimals"`
	Currency        *string                `json:"currency" validate:"omitempty,required,oneof=BRL USD EUR"`
	Tax             *float64               `json:"tax" validate:"omitempty,gte=0,twodecimals"`
	CostAllocations []*CostAllocationInput `json:"cost_allocations" validate:"omitempty,required,dive"`
}

type UpdateCostOutputDTO struct {
	mapper.CostOutput
}

// NewUpdateCostUseCase creates a new instance of UpdateCostUseCase
func NewUpdateCostUseCase(txm db.TransactionManagerInterface) *UpdateCostUseCase {
	return &UpdateCostUseCase{txm}
}

// Execute updates a cost
func (uc *UpdateCostUseCase) Execute(ctx context.Context, input UpdateCostInputDTO) (*UpdateCostOutputDTO, error) {
	var updatedCost *domain.Cost

	err := uc.txm.Do(ctx, func(ctx context.Context, tx db.TransactionInterface) error {
		repository, err := db.GetAs[domain.EstimationRepository](tx, "EstimationRepository")
		if err != nil {
			return err
		}

		cost, err := repository.GetCost(ctx, input.CostID)
		if err != nil {
			return err
		}

		cost.ChangeCostType(input.CostType)
		cost.ChangeDescription(input.Description)
		cost.ChangeComment(input.Comment)
		cost.ChangeAmount(input.Amount)
		cost.ChangeCurrency(input.Currency)
		cost.ChangeTax(input.Tax)

		if input.CostAllocations != nil {
			costAllocations := make([]domain.CostAllocationProps, len(input.CostAllocations))
			for i, allocation := range input.CostAllocations {
				costAllocations[i] = domain.CostAllocationProps{
					Year:   allocation.Year,
					Month:  time.Month(allocation.Month),
					Amount: allocation.Amount,
				}
			}
			cost.ChangeCostAllocations(costAllocations)
		}

		if err := cost.Validate(); err != nil {
			return err
		}

		err = repository.UpdateCost(ctx, cost)
		if err != nil {
			return err
		}

		updatedCost, err = repository.GetCost(ctx, cost.CostID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	outputCostAllocations := make([]mapper.CostAllocationOutput, len(updatedCost.CostAllocations))

	for i := range updatedCost.CostAllocations {
		outputCostAllocations[i] = mapper.CostAllocationOutput{
			Year:   updatedCost.CostAllocations[i].AllocationDate.Year(),
			Month:  int(updatedCost.CostAllocations[i].AllocationDate.Month()),
			Amount: updatedCost.CostAllocations[i].Amount,
		}
	}

	output := mapper.CostOutput{
		CostID:          updatedCost.CostID,
		BaselineID:      updatedCost.BaselineID,
		CostType:        string(updatedCost.CostType),
		Description:     updatedCost.Description,
		Comment:         updatedCost.Comment,
		Amount:          updatedCost.Amount,
		Currency:        string(updatedCost.Currency),
		Tax:             updatedCost.Tax,
		CostAllocations: outputCostAllocations,
		CreatedAt:       updatedCost.CreatedAt,
		UpdatedAt:       updatedCost.UpdatedAt,
	}

	return &UpdateCostOutputDTO{output}, nil
}

// DeleteCostUseCase represents the use case for deleting a cost
type DeleteCostUseCase struct {
	txm db.TransactionManagerInterface
}

type DeleteCostInputDTO struct {
	CostID string `json:"cost_id" validate:"required"`
}

type DeleteCostOutputDTO struct{}

// NewDeleteCostUseCase creates a new instance of DeleteCostUseCase
func NewDeleteCostUseCase(txm db.TransactionManagerInterface) *DeleteCostUseCase {
	return &DeleteCostUseCase{txm}
}

// Execute deletes a cost
func (uc *DeleteCostUseCase) Execute(ctx context.Context, input DeleteCostInputDTO) (*DeleteCostOutputDTO, error) {
	err := uc.txm.Do(ctx, func(ctx context.Context, tx db.TransactionInterface) error {
		repository, err := db.GetAs[domain.EstimationRepository](tx, "EstimationRepository")
		if err != nil {
			return err
		}

		err = repository.DeleteCost(ctx, input.CostID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &DeleteCostOutputDTO{}, nil
}
