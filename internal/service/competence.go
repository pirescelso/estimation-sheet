package service

import (
	"context"

	"github.com/celsopires1999/estimation/internal/mapper"
)

func (s *EstimationService) ListCompetences(ctx context.Context, imput ListCompetencesInputDTO) (*ListCompetencesOutputDTO, error) {
	competences, err := s.queries.FindAllCompetences(ctx)
	if err != nil {
		return nil, err
	}

	competencesOutput := make([]mapper.CompetenceOutput, len(competences))
	for i, competence := range competences {
		competencesOutput[i] = mapper.CompetenceOutputFromDb(competence)
	}

	return &ListCompetencesOutputDTO{Competences: competencesOutput}, nil
}

type ListCompetencesInputDTO struct{}
type ListCompetencesOutputDTO struct {
	Competences []mapper.CompetenceOutput `json:"competences"`
}
