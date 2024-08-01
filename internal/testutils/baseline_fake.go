package testutils

import (
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/google/uuid"
)

type BaselineFakeBuilder struct {
	BaselineID  string
	Code        string
	Review      int32
	Title       string
	Description string
	StartDate   time.Time
	Duration    int32
	ManagerID   string
	EstimatorID string
	CreatedAt   time.Time
	updatedAt   time.Time
}

func NewBaselineFakeBuilder() *BaselineFakeBuilder {
	return &BaselineFakeBuilder{
		BaselineID:  uuid.New().String(),
		Code:        randomdata.Alphanumeric(20),
		Review:      int32(1),
		Title:       randomdata.SillyName(),
		Description: randomdata.Paragraph(),
		StartDate:   time.Date(randomdata.Number(2020, 2030), time.Month(randomdata.Number(1, 12)), 1, 0, 0, 0, 0, time.UTC),
		Duration:    int32(randomdata.Number(1, 60)),
		ManagerID:   uuid.New().String(),
		EstimatorID: uuid.New().String(),
		CreatedAt:   time.Now(),
		updatedAt:   time.Now(),
	}
}

func (b *BaselineFakeBuilder) WithBaselineID(baselineID string) *BaselineFakeBuilder {
	b.BaselineID = baselineID
	return b
}

func (b *BaselineFakeBuilder) WithCode(code string) *BaselineFakeBuilder {
	b.Code = code
	return b
}

func (b *BaselineFakeBuilder) WithReview(review int32) *BaselineFakeBuilder {
	b.Review = review
	return b
}

func (b *BaselineFakeBuilder) WithTitle(title string) *BaselineFakeBuilder {
	b.Title = title
	return b
}

func (b *BaselineFakeBuilder) WithDescription(description string) *BaselineFakeBuilder {
	b.Description = description
	return b
}

func (b *BaselineFakeBuilder) WithStartDate(startDate time.Time) *BaselineFakeBuilder {
	b.StartDate = startDate
	return b
}

func (b *BaselineFakeBuilder) WithDuration(duration int32) *BaselineFakeBuilder {
	b.Duration = duration
	return b
}

func (b *BaselineFakeBuilder) WithManagerID(managerID string) *BaselineFakeBuilder {
	b.ManagerID = managerID
	return b
}

func (b *BaselineFakeBuilder) WithEstimatorID(estimatorID string) *BaselineFakeBuilder {
	b.EstimatorID = estimatorID
	return b
}

func (b *BaselineFakeBuilder) WithCreatedAt(createdAt time.Time) *BaselineFakeBuilder {
	b.CreatedAt = createdAt
	return b
}

func (b *BaselineFakeBuilder) WithUpdatedAt(updatedAt time.Time) *BaselineFakeBuilder {
	b.updatedAt = updatedAt
	return b
}

func (b *BaselineFakeBuilder) Build() *domain.Baseline {
	props := domain.RestoreBaselineProps{}

	props.BaselineID = b.BaselineID
	props.Review = b.Review
	props.Code = b.Code
	props.Title = b.Title
	props.Description = b.Description
	props.StartDate = b.StartDate
	props.Duration = b.Duration
	props.ManagerID = b.ManagerID
	props.EstimatorID = b.EstimatorID
	props.CreatedAt = b.CreatedAt
	props.UpdatedAt = b.updatedAt

	baseline := domain.RestoreBaseline(props)
	err := baseline.Validate()
	if err != nil {
		panic(err)
	}
	return baseline
}
