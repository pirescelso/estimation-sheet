package testutils

import (
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/google/uuid"
)

type EffortFakeBuilder struct {
	EffortID               string
	BaselineID             string
	CompetenceID           string
	Comment                string
	Hours                  int
	EffortAllocationsProps []domain.EffortAllocationProps
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

func NewEffortFakeBuilder() *EffortFakeBuilder {
	return &EffortFakeBuilder{
		EffortID:     uuid.NewString(),
		BaselineID:   uuid.NewString(),
		CompetenceID: uuid.NewString(),
		Comment:      randomdata.SillyName(),
		Hours:        160,
		EffortAllocationsProps: []domain.EffortAllocationProps{
			{Year: 2020, Month: time.January, Hours: 60},
			{Year: 2020, Month: time.August, Hours: 55},
			{Year: 2020, Month: time.December, Hours: 45},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (b *EffortFakeBuilder) WithEffortID(id string) *EffortFakeBuilder {
	b.EffortID = id
	return b
}

func (b *EffortFakeBuilder) WithBaselineID(id string) *EffortFakeBuilder {
	b.BaselineID = id
	return b
}

func (b *EffortFakeBuilder) WithCompetenceID(id string) *EffortFakeBuilder {
	b.CompetenceID = id
	return b
}

func (b *EffortFakeBuilder) WithComment(comment string) *EffortFakeBuilder {
	b.Comment = comment
	return b
}

func (b *EffortFakeBuilder) WithHours(hours int) *EffortFakeBuilder {
	b.Hours = hours
	return b
}

func (b *EffortFakeBuilder) WithEffortAllocationsProps(props []domain.EffortAllocationProps) *EffortFakeBuilder {
	b.EffortAllocationsProps = props
	return b
}

func (b *EffortFakeBuilder) WithCreatedAt(createdAt time.Time) *EffortFakeBuilder {
	b.CreatedAt = createdAt
	return b
}

func (b *EffortFakeBuilder) WithUpdatedAt(updatedAt time.Time) *EffortFakeBuilder {
	b.UpdatedAt = updatedAt
	return b
}

func (b *EffortFakeBuilder) Build() *domain.Effort {
	allocations := make([]domain.EffortAllocation, len(b.EffortAllocationsProps))

	for i, props := range b.EffortAllocationsProps {
		allocations[i] = newEffortAllocation(props.Year, props.Month, props.Hours)
	}

	props := domain.RestoreEffortProps{
		EffortID:          b.EffortID,
		BaselineID:        b.BaselineID,
		CompetenceID:      b.CompetenceID,
		Comment:           b.Comment,
		Hours:             b.Hours,
		EffortAllocations: allocations,
		CreatedAt:         b.CreatedAt,
		UpdatedAt:         b.UpdatedAt,
	}

	effort := domain.RestoreEffort(props)
	err := effort.Validate()
	if err != nil {
		panic(err)
	}
	return effort
}

func newEffortAllocation(year int, month time.Month, hours int) domain.EffortAllocation {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return domain.EffortAllocation{
		AllocationDate: date,
		Hours:          hours,
	}
}
