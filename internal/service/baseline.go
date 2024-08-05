package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/celsopires1999/estimation/internal/mapper"
	"github.com/jackc/pgx/v5"
)

func (s *EstimationService) GetBaseline(ctx context.Context, input GetBaselineInputDTO) (*GetBaselineOutputDTO, error) {
	baseline, err := s.queries.FindBaselineByIdWithRelations(ctx, input.BaselineID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, common.NewNotFoundError(fmt.Errorf("baseline with id %s not found", input.BaselineID))
		}
		return nil, err
	}
	output := mapper.BaselineOutputFromDb(db.BaselineRow(baseline))

	return &GetBaselineOutputDTO{output}, nil
}

type GetBaselineInputDTO struct {
	BaselineID string
}

type GetBaselineOutputDTO struct {
	mapper.BaselineOutput
}

func (s *EstimationService) ListBaselines(ctx context.Context, input ListBaselinesInputDTO) (*ListBaselinesOutputDTO, error) {
	baselines, err := s.queries.FindAllBaselines(ctx)
	if err != nil {
		return nil, err
	}

	baselinesOutput := make([]mapper.BaselineOutput, len(baselines))
	for i, baseline := range baselines {
		baselinesOutput[i] = mapper.BaselineOutputFromDb(db.BaselineRow(baseline))
	}

	return &ListBaselinesOutputDTO{baselinesOutput}, nil
}

type ListBaselinesInputDTO struct{}

type ListBaselinesOutputDTO struct {
	Baselines []mapper.BaselineOutput `json:"baselines"`
}
