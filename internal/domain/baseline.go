package domain

import (
	"errors"
	"log"
	"time"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/google/uuid"
)

type Baseline struct {
	BaselineID  string    `validate:"required,uuid4"`
	Code        string    `validate:"required,max=20"`
	Review      int32     `validate:"gt=0"`
	Title       string    `validate:"required"`
	Description string    `validate:"-"`
	StartDate   time.Time `validate:"required"`
	Duration    int32     `validate:"gt=0,lte=60"`
	ManagerID   string    `validate:"required"`
	EstimatorID string    `validate:"required"`
	CreatedAt   time.Time `validate:"-"`
	UpdatedAt   time.Time `validate:"-"`
}

type RestoreBaselineProps Baseline

var ErrBaselineDomainValidation = errors.New("baseline domain validation failed")

func NewBaseline(
	code string,
	review int32,
	title string,
	description string,
	startDate time.Time,
	duration int32,
	managerID string,
	estimatorID string,
) *Baseline {
	return &Baseline{
		BaselineID:  uuid.New().String(),
		Code:        code,
		Review:      review,
		Title:       title,
		Description: description,
		StartDate:   startDate,
		Duration:    duration,
		ManagerID:   managerID,
		EstimatorID: estimatorID,
	}
}

func RestoreBaseline(props RestoreBaselineProps) *Baseline {
	return &Baseline{
		BaselineID:  props.BaselineID,
		Code:        props.Code,
		Review:      props.Review,
		Title:       props.Title,
		Description: props.Description,
		StartDate:   props.StartDate,
		Duration:    props.Duration,
		ManagerID:   props.ManagerID,
		EstimatorID: props.EstimatorID,
		CreatedAt:   props.CreatedAt,
		UpdatedAt:   props.UpdatedAt,
	}

}

func (b *Baseline) ChangeCode(code *string) {
	if code == nil {
		return
	}
	b.Code = *code
}

func (b *Baseline) ChangeReview(review *int32) {
	if review == nil {
		return
	}
	b.Review = *review
}

func (b *Baseline) ChangeTitle(title *string) {
	if title == nil {
		return
	}
	b.Title = *title
}

func (b *Baseline) ChangeDescription(description *string) {
	if description == nil {
		return
	}
	b.Description = *description
}

func (b *Baseline) ChangeStartDate(startYear, startMonth *int) {
	if startYear == nil && startMonth == nil {
		return
	}

	tmpStartYear := b.StartDate.Year()
	tmpStartMonth := b.StartDate.Month()

	if startYear != nil {
		tmpStartYear = *startYear
	}
	if startMonth != nil {
		tmpStartMonth = time.Month(*startMonth)
	}

	b.StartDate = time.Date(tmpStartYear, tmpStartMonth, 1, 0, 0, 0, 0, time.UTC)
}

func (b *Baseline) ChangeDuration(duration *int32) {
	if duration == nil {
		return
	}
	b.Duration = *duration
}

func (b *Baseline) ChangeManagerID(managerID *string) {
	if managerID == nil {
		return
	}
	b.ManagerID = *managerID
}

func (b *Baseline) ChangeEstimatorID(estimatorID *string) {
	if estimatorID == nil {
		return
	}
	b.EstimatorID = *estimatorID
}

func (b *Baseline) Validate() error {
	err := common.Validate.Struct(b)
	if err != nil {
		log.Printf("baseline domain validation failed: %v\nbaseline: %+v", err, b)
		return ErrBaselineDomainValidation
	}
	return nil
}
