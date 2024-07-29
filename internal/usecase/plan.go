package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/mapper"
	"github.com/jackc/pgx/v5/pgconn"
)

// CreatePlanUseCase is responsible for creating a new plan in the system
type CreatePlanUseCase struct {
	repository domain.EstimationRepository
}

type CreatePlanInputDTO struct {
	Code        string             `json:"code" validate:"required,max=10"`
	Name        string             `json:"name" validate:"required,max=50"`
	Assumptions domain.Assumptions `json:"assumptions" validate:"required,dive"`
}

type CreatePlanOutputDTO struct {
	mapper.PlanOutput
}

func NewCreatePlanUseCase(repo domain.EstimationRepository) *CreatePlanUseCase {
	return &CreatePlanUseCase{repo}
}

func (uc *CreatePlanUseCase) Execute(ctx context.Context, input CreatePlanInputDTO) (*CreatePlanOutputDTO, error) {
	plan := domain.NewPlan(input.Code, input.Name, input.Assumptions)
	if err := plan.Validate(); err != nil {
		return nil, err
	}

	err := uc.repository.CreatePlan(ctx, plan)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil, common.NewConflictError(fmt.Errorf("plan with code %s already exists", input.Code))
			}
			return nil, common.NewConflictError(err)
		}
		return nil, err
	}

	createdPlan, err := uc.repository.GetPlan(ctx, plan.PlanID)
	if err != nil {
		return nil, err
	}

	output := mapper.PlanOutput{
		PlanID:      createdPlan.PlanID,
		Code:        createdPlan.Code,
		Name:        createdPlan.Name,
		Assumptions: createdPlan.Assumptions,
		CreatedAt:   createdPlan.CreatedAt,
	}

	return &CreatePlanOutputDTO{output}, nil
}

// GetPlanUseCase is a use case to get a plan by its ID
type GetPlanUseCase struct {
	repository domain.EstimationRepository
}

type GetPlanInputDTO struct {
	PlanID string `json:"plan_id" validate:"required"`
}

type GetPlanOutputDTO struct {
	mapper.PlanOutput
}

func NewGetPlanUseCase(repository domain.EstimationRepository) *GetPlanUseCase {
	return &GetPlanUseCase{repository}
}

func (uc *GetPlanUseCase) Execute(ctx context.Context, input GetPlanInputDTO) (*GetPlanOutputDTO, error) {
	plan, err := uc.repository.GetPlan(ctx, input.PlanID)
	if err != nil {
		return nil, err
	}

	output := mapper.PlanOutput{
		PlanID:      plan.PlanID,
		Code:        plan.Code,
		Name:        plan.Name,
		Assumptions: plan.Assumptions,
		CreatedAt:   plan.CreatedAt,
		UpdatedAt:   plan.UpdatedAt,
	}

	return &GetPlanOutputDTO{output}, nil

}

// UpdatePlanUseCase is a use case to update a plan
type UpdatePlanUseCase struct {
	repository domain.EstimationRepository
}

type UpdatePlanInputDTO struct {
	PlanID      string              `json:"plan_id" validate:"required"`
	Code        *string             `json:"code" validate:"omitempty,max=10"`
	Name        *string             `json:"name" validate:"omitempty,max=50"`
	Assumptions *domain.Assumptions `json:"assumptions" validate:"omitempty,required,dive"`
}

type UpdatePlanOutputDTO struct {
	mapper.PlanOutput
}

func NewUpdatePlanUseCase(repository domain.EstimationRepository) *UpdatePlanUseCase {
	return &UpdatePlanUseCase{repository}
}

func (uc *UpdatePlanUseCase) Execute(ctx context.Context, input UpdatePlanInputDTO) (*UpdatePlanOutputDTO, error) {
	plan, err := uc.repository.GetPlan(ctx, input.PlanID)
	if err != nil {
		return nil, err
	}

	if input.Code != nil {
		plan.ChangeCode(*input.Code)
	}
	if input.Name != nil {
		plan.ChangeName(*input.Name)
	}
	if input.Assumptions != nil {
		plan.ChangeAssumptions(*input.Assumptions)
	}
	if err := plan.Validate(); err != nil {
		return nil, err
	}

	count, err := uc.repository.CountPortfoliosByPlanId(ctx, plan.PlanID)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, common.NewConflictError(fmt.Errorf("plan %s has %d portfolio(s)", plan.Code, count))
	}

	if err := uc.repository.UpdatePlan(ctx, plan); err != nil {
		return nil, err
	}

	updated, err := uc.repository.GetPlan(ctx, plan.PlanID)
	if err != nil {
		return nil, err
	}

	output := mapper.PlanOutput{
		PlanID:      updated.PlanID,
		Code:        updated.Code,
		Name:        updated.Name,
		Assumptions: updated.Assumptions,
		CreatedAt:   updated.CreatedAt,
		UpdatedAt:   updated.UpdatedAt,
	}

	return &UpdatePlanOutputDTO{output}, nil
}

// DeletePlanUseCase is a use case to delete a plan
type DeletePlanUseCase struct {
	repository domain.EstimationRepository
}

type DeletePlanInputDTO struct {
	PlanID string `json:"plan_id" validate:"required"`
}

type DeletePlanOutputDTO struct{}

func NewDeletePlanUseCase(repository domain.EstimationRepository) *DeletePlanUseCase {
	return &DeletePlanUseCase{repository}
}

func (uc *DeletePlanUseCase) Execute(ctx context.Context, planID string) (*DeletePlanOutputDTO, error) {
	count, err := uc.repository.CountPortfoliosByPlanId(ctx, planID)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, common.NewConflictError(fmt.Errorf("plan %s has %d portfolio(s)", planID, count))
	}
	return &DeletePlanOutputDTO{}, uc.repository.DeletePlan(ctx, planID)
}
