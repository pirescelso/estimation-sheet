package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/celsopires1999/estimation/internal/mapper"
)

var (
	ErrEffortAllocationDateIsInvalid = errors.New("effort allocation date is invalid")
	ErrEffortBaselineMismatch        = errors.New("effort baseline mismatch")
)

type CreateEffortUseCase struct {
	txm db.TransactionManagerInterface
}

type CreateEffortInputDTO struct {
	BaselineID        string                  `json:"baseline_id" validate:"required,uuid4"`
	CompetenceID      string                  `json:"competence_id" validate:"required,uuid4"`
	Comment           string                  `json:"comment"`
	Hours             int                     `json:"hours" validate:"required,gte=1,lte=160_000"`
	EffortAllocations []EffortAllocationInput `json:"effort_allocations" validate:"required,dive"`
}

type EffortAllocationInput struct {
	Year  int `json:"year" validate:"required"`
	Month int `json:"month" validate:"gte=1,lte=12"`
	Hours int `json:"hours" validate:"required,gte=1,lte=8_000"`
}

type CreateEffortOutputDTO struct {
	mapper.EffortOutput
}

func NewCreateEffortUseCase(txm db.TransactionManagerInterface) *CreateEffortUseCase {
	return &CreateEffortUseCase{txm}
}

func (uc *CreateEffortUseCase) Execute(ctx context.Context, input CreateEffortInputDTO) (*CreateEffortOutputDTO, error) {
	var createdEffort *domain.Effort

	err := uc.txm.Do(ctx, func(ctx context.Context, tx db.TransactionInterface) error {
		repository, err := db.GetAs[domain.EstimationRepository](tx, "EstimationRepository")
		if err != nil {
			return err
		}

		baseline, err := repository.GetBaseline(ctx, input.BaselineID)
		if err != nil {
			return err
		}

		count, err := repository.CountPortfoliosByBaselineId(ctx, input.BaselineID)
		if err != nil {
			return err
		}
		if count > 0 {
			return common.NewConflictError(fmt.Errorf("baseline %s has %d portfolio(s)", baseline.BaselineID, count))
		}

		effortAllocations := make([]domain.EffortAllocationProps, len(input.EffortAllocations))
		for i, allocation := range input.EffortAllocations {
			effortAllocations[i] = domain.EffortAllocationProps{
				Year:  allocation.Year,
				Month: time.Month(allocation.Month),
				Hours: allocation.Hours,
			}
		}
		effort := domain.NewEffort(domain.NewEffortProps{
			BaselineID:        input.BaselineID,
			CompetenceID:      input.CompetenceID,
			Comment:           input.Comment,
			Hours:             input.Hours,
			EffortAllocations: effortAllocations,
		})

		if err := effort.Validate(); err != nil {
			return err
		}

		for _, a := range effort.EffortAllocations {
			if baseline.StartDate.After(a.AllocationDate) {
				return ErrEffortAllocationDateIsInvalid
			}
		}

		err = repository.CreateEffort(ctx, effort)
		if err != nil {
			return err
		}

		createdEffort, err = repository.GetEffort(ctx, effort.EffortID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	output := mapper.EffortOutputFromDomain(*createdEffort)

	return &CreateEffortOutputDTO{output}, nil
}

type UpdateEffortUseCase struct {
	txm db.TransactionManagerInterface
}

type UpdateEffortInputDTO struct {
	EffortID          string                   `json:"effort_id" validate:"required,uuid4"`
	BaselineID        string                   `json:"baseline_id" validate:"required,uuid4"`
	CompetenceID      *string                  `json:"competence_id" validate:"omitempty,required,uuid4"`
	Comment           *string                  `json:"comment"`
	Hours             *int                     `json:"hours" validate:"omitempty,required,gte=1,lte=160_000"`
	EffortAllocations []*EffortAllocationInput `json:"effort_allocations" validate:"omitempty,required,dive"`
}

type UpdateEffortOutputDTO struct {
	mapper.EffortOutput
}

func NewUpdateEfforttUseCase(txm db.TransactionManagerInterface) *UpdateEffortUseCase {
	return &UpdateEffortUseCase{txm}
}

func (uc *UpdateEffortUseCase) Execute(ctx context.Context, input UpdateEffortInputDTO) (*UpdateEffortOutputDTO, error) {
	var updatedEffort *domain.Effort

	err := uc.txm.Do(ctx, func(ctx context.Context, tx db.TransactionInterface) error {
		repository, err := db.GetAs[domain.EstimationRepository](tx, "EstimationRepository")
		if err != nil {
			return err
		}

		effort, err := repository.GetEffort(ctx, input.EffortID)
		if err != nil {
			return err
		}

		if effort.BaselineID != input.BaselineID {
			return ErrEffortBaselineMismatch
		}

		count, err := repository.CountPortfoliosByBaselineId(ctx, input.BaselineID)
		if err != nil {
			return err
		}
		if count > 0 {
			return common.NewConflictError(fmt.Errorf("baseline %s has %d portfolio(s)", input.BaselineID, count))
		}

		effort.ChangeComment(input.Comment)
		effort.ChangeHours(input.Hours)

		if input.EffortAllocations != nil {
			effortAllocations := make([]domain.EffortAllocationProps, len(input.EffortAllocations))
			for i, allocation := range input.EffortAllocations {
				effortAllocations[i] = domain.EffortAllocationProps{
					Year:  allocation.Year,
					Month: time.Month(allocation.Month),
					Hours: allocation.Hours,
				}
			}
			effort.ChangeEffortAllocations(effortAllocations)
		}

		if err := effort.Validate(); err != nil {
			return err
		}

		err = repository.UpdateEffort(ctx, effort)
		if err != nil {
			return err
		}

		updatedEffort, err = repository.GetEffort(ctx, effort.EffortID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	output := mapper.EffortOutputFromDomain(*updatedEffort)

	return &UpdateEffortOutputDTO{output}, nil
}

type DeleteEffortUseCase struct {
	txm db.TransactionManagerInterface
}

type DeleteEffortInputDTO struct {
	EffortID   string `json:"effort_id" validate:"required,uuid4"`
	BaselineID string `json:"baseline_id" validate:"required,uuid4"`
}

type DeleteEffortOutputDTO struct{}

func NewDeleteEffortUseCase(txm db.TransactionManagerInterface) *DeleteEffortUseCase {
	return &DeleteEffortUseCase{txm}
}

func (uc *DeleteEffortUseCase) Execute(ctx context.Context, input DeleteEffortInputDTO) (*DeleteEffortOutputDTO, error) {
	err := uc.txm.Do(ctx, func(ctx context.Context, tx db.TransactionInterface) error {
		repository, err := db.GetAs[domain.EstimationRepository](tx, "EstimationRepository")
		if err != nil {
			return err
		}

		effort, err := repository.GetEffort(ctx, input.EffortID)
		if err != nil {
			return err
		}

		if effort.BaselineID != input.BaselineID {
			return ErrEffortBaselineMismatch
		}

		count, err := repository.CountPortfoliosByBaselineId(ctx, input.BaselineID)
		if err != nil {
			return err
		}
		if count > 0 {
			return common.NewConflictError(fmt.Errorf("baseline %s has %d portfolio(s)", input.BaselineID, count))
		}

		err = repository.DeleteEffort(ctx, input.EffortID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &DeleteEffortOutputDTO{}, nil
}

type GetEffortsByBaselineIDUseCase struct {
	repository domain.EstimationRepository
}

type GetEffortsByBaselineIDInputDTO struct {
	BaselineID string `json:"baseline_id" validate:"required,uuid4"`
}

type GetEffortsByBaselineIDOutputDTO struct {
	Efforts []mapper.EffortOutput `json:"efforts"`
}

func NewGetEffortsByBaselineIDUseCase(repository domain.EstimationRepository) *GetEffortsByBaselineIDUseCase {
	return &GetEffortsByBaselineIDUseCase{repository}
}

func (uc *GetEffortsByBaselineIDUseCase) Execute(ctx context.Context, input GetEffortsByBaselineIDInputDTO) (*GetEffortsByBaselineIDOutputDTO, error) {
	_, err := uc.repository.GetBaseline(ctx, input.BaselineID)
	if err != nil {
		return nil, err
	}

	efforts, err := uc.repository.GetEffortManyByBaselineID(ctx, input.BaselineID)
	if err != nil {
		return nil, err
	}

	output := make([]mapper.EffortOutput, len(efforts))

	for i, effort := range efforts {
		output[i] = mapper.EffortOutputFromDomain(*effort)
	}

	return &GetEffortsByBaselineIDOutputDTO{Efforts: output}, nil
}
