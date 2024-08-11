package testutils

import (
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/google/uuid"
)

type CompetenceFakeBuilder struct {
	CompetenceID string
	Code         string
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewCompetenceFakeBuilder() *CompetenceFakeBuilder {
	return &CompetenceFakeBuilder{
		CompetenceID: uuid.NewString(),
		Code:         randomdata.Digits(20),
		Name:         randomdata.SillyName(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (b *CompetenceFakeBuilder) WithCode(code string) *CompetenceFakeBuilder {
	b.Code = code
	return b
}

func (b *CompetenceFakeBuilder) WithName(name string) *CompetenceFakeBuilder {
	b.Name = name
	return b
}

func (b *CompetenceFakeBuilder) WithCreatedAt(createdAt time.Time) *CompetenceFakeBuilder {
	b.CreatedAt = createdAt
	return b
}

func (b *CompetenceFakeBuilder) WithUpdatedAt(updatedAt time.Time) *CompetenceFakeBuilder {
	b.UpdatedAt = updatedAt
	return b
}

func (b *CompetenceFakeBuilder) Build() *domain.Competence {
	return &domain.Competence{
		CompetenceID: b.CompetenceID,
		Code:         b.Code,
		Name:         b.Name,
		CreatedAt:    b.CreatedAt,
		UpdatedAt:    b.UpdatedAt,
	}
}
