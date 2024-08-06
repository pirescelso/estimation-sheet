package service

import (
	"context"

	"github.com/celsopires1999/estimation/internal/mapper"
)

func (s *EstimationService) ListPlans(ctx context.Context, input ListPlansInputDTO) (*ListPlansOutputDTO, error) {
	plans, err := s.queries.FindAllPlans(ctx)
	if err != nil {
		return nil, err
	}

	plansOutput := make([]mapper.PlanOutput, len(plans))
	for i, plan := range plans {
		plansOutput[i] = mapper.PlanOutputFromDb(plan)
	}

	return &ListPlansOutputDTO{plansOutput}, nil
}

type ListPlansInputDTO struct{}

type ListPlansOutputDTO struct {
	Plans []mapper.PlanOutput `json:"plans"`
}
