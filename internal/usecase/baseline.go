package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/mapper"
)

// CreateBaselineUseCase is responsible for creating a new baseline in the system
type CreateBaselineUseCase struct {
	repository domain.EstimationRepository
}

type CreateBaselineInputDTO struct {
	Code        string  `json:"code" validate:"required,max=20"`
	Review      int32   `json:"review" validate:"required,gte=1"`
	Title       string  `json:"title" validate:"required"`
	Description *string `json:"description" validate:"-"`
	StartMonth  int     `json:"start_month" validate:"gte=1,lte=12"`
	StartYear   int     `json:"start_year" validate:"required"`
	Duration    int32   `json:"duration" validate:"gt=0,lte=60"`
	ManagerID   string  `json:"manager_id" validate:"required,uuid4"`
	EstimatorID string  `json:"estimator_id" validate:"required,uuid4"`
}

type CreateBaselineOutputDTO struct {
	mapper.BaselineOutput
}

func NewCreateBaselineUseCase(repo domain.EstimationRepository) *CreateBaselineUseCase {
	return &CreateBaselineUseCase{repo}
}

func (uc *CreateBaselineUseCase) Execute(ctx context.Context, input CreateBaselineInputDTO) (*CreateBaselineOutputDTO, error) {
	startDate := time.Date(input.StartYear, time.Month(input.StartMonth), 1, 0, 0, 0, 0, time.UTC)

	description := ""
	if input.Description != nil {
		description = *input.Description
	}

	baseline := domain.NewBaseline(
		input.Code,
		input.Review,
		input.Title,
		description,
		startDate,
		input.Duration,
		input.ManagerID,
		input.EstimatorID,
	)
	if err := uc.repository.CreateBaseline(ctx, baseline); err != nil {
		return nil, err
	}

	createdBaseline, err := uc.repository.GetBaseline(ctx, baseline.BaselineID)
	if err != nil {
		return nil, err
	}

	output := mapper.BaselineOutput{
		BaselineID:  createdBaseline.BaselineID,
		Code:        createdBaseline.Code,
		Review:      createdBaseline.Review,
		Title:       createdBaseline.Title,
		Description: createdBaseline.Description,
		StartDate:   createdBaseline.StartDate,
		Duration:    createdBaseline.Duration,
		ManagerID:   createdBaseline.ManagerID,
		EstimatorID: createdBaseline.EstimatorID,
		CreatedAt:   createdBaseline.CreatedAt,
	}

	return &CreateBaselineOutputDTO{output}, nil
}

// UpdateBaselineUseCase is responsible for updating an existing baseline in the system
type UpdateBaselineUseCase struct {
	repository domain.EstimationRepository
}

type UpdateBaselineInputDTO struct {
	BaselineID  string  `json:"baseline_id" validate:"required,uuid4"`
	Code        *string `json:"code" validate:"omitempty,max=20"`
	Review      *int32  `json:"review" validate:"omitempty,gte=1"`
	Title       *string `json:"title" validate:"omitempty"`
	Description *string `json:"description" validate:"omitempty"`
	StartMonth  *int    `json:"start_month" validate:"omitempty,gte=1,lte=12"`
	StartYear   *int    `json:"start_year" validate:"omitempty"`
	Duration    *int32  `json:"duration" validate:"omitempty,gt=0,lte=60"`
	ManagerID   *string `json:"manager_id" validate:"omitempty,uuid4"`
	EstimatorID *string `json:"estimator_id" validate:"omitempty,uuid4"`
}

type UpdateBaselineOutputDTO struct {
	mapper.BaselineOutput
}

func NewUpdateBaselineUseCase(repo domain.EstimationRepository) *UpdateBaselineUseCase {
	return &UpdateBaselineUseCase{repo}
}

func (uc *UpdateBaselineUseCase) Execute(ctx context.Context, input UpdateBaselineInputDTO) (*UpdateBaselineOutputDTO, error) {
	baseline, err := uc.repository.GetBaseline(ctx, input.BaselineID)
	if err != nil {
		return nil, err
	}

	baseline.ChangeCode(input.Code)
	baseline.ChangeReview(input.Review)
	baseline.ChangeTitle(input.Title)
	baseline.ChangeDescription(input.Description)
	baseline.ChangeStartDate(input.StartYear, input.StartMonth)
	baseline.ChangeDuration(input.Duration)
	baseline.ChangeManagerID(input.ManagerID)
	baseline.ChangeEstimatorID(input.EstimatorID)

	err = baseline.Validate()
	if err != nil {
		return nil, err
	}

	count, err := uc.repository.CountPortfoliosByBaselineId(ctx, baseline.BaselineID)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, common.NewConflictError(fmt.Errorf("baseline %s has %d portfolio(s)", baseline.BaselineID, count))
	}

	err = uc.repository.UpdateBaseline(ctx, baseline)
	if err != nil {
		return nil, err
	}

	updated, err := uc.repository.GetBaseline(ctx, baseline.BaselineID)
	if err != nil {
		return nil, err
	}

	output := mapper.BaselineOutput{
		BaselineID:  updated.BaselineID,
		Code:        updated.Code,
		Review:      updated.Review,
		Title:       updated.Title,
		Description: updated.Description,
		StartDate:   updated.StartDate,
		Duration:    updated.Duration,
		ManagerID:   updated.ManagerID,
		EstimatorID: updated.EstimatorID,
		CreatedAt:   updated.CreatedAt,
		UpdatedAt:   updated.UpdatedAt,
	}

	return &UpdateBaselineOutputDTO{output}, nil
}

// DeleteBaselineUseCase is responsible for deleting an existing baseline in the system
type DeleteBaselineUseCase struct {
	repository domain.EstimationRepository
}

type DeleteBaselineInputDTO struct {
	BaselineID string `json:"baseline_id" validate:"required"`
}

type DeleteBaselineOutputDTO struct{}

func NewDeleteBaselineUseCase(repo domain.EstimationRepository) *DeleteBaselineUseCase {
	return &DeleteBaselineUseCase{repo}
}

func (uc *DeleteBaselineUseCase) Execute(ctx context.Context, input DeleteBaselineInputDTO) (*DeleteBaselineOutputDTO, error) {
	err := uc.repository.DeleteBaseline(ctx, input.BaselineID)
	if err != nil {
		return nil, err
	}
	return &DeleteBaselineOutputDTO{}, nil
}
