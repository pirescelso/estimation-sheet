package domain

import (
	"fmt"
	"slices"
	"time"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/google/uuid"
)

type Effort struct {
	EffortID          string             `validate:"required,uuid"`
	BaselineID        string             `validate:"required,uuid"`
	CompetenceID      string             `validate:"required,uuid"`
	Comment           string             `validate:"-"`
	Hours             int                `validate:"required,gte=1"`
	EffortAllocations []EffortAllocation `validate:"required"`
	CreatedAt         time.Time          `validate:"-"`
	UpdatedAt         time.Time          `validate:"-"`
}

type EffortAllocation struct {
	AllocationDate time.Time
	Hours          int
}

type RestoreEffortProps Effort

type EffortAllocationProps struct {
	Year  int
	Month time.Month
	Hours int
}

type NewEffortProps struct {
	BaselineID        string
	CompetenceID      string
	Comment           string
	Hours             int
	EffortAllocations []EffortAllocationProps
}

func NewEffort(props NewEffortProps) *Effort {
	effortAllocations := createEffortAllocations(props.EffortAllocations)
	return &Effort{
		EffortID:          uuid.New().String(),
		BaselineID:        props.BaselineID,
		CompetenceID:      props.CompetenceID,
		Comment:           props.Comment,
		Hours:             props.Hours,
		EffortAllocations: effortAllocations,
	}
}

func newEffortAllocation(year int, month time.Month, hours int) EffortAllocation {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return EffortAllocation{
		AllocationDate: date,
		Hours:          hours,
	}
}

func createEffortAllocations(params []EffortAllocationProps) []EffortAllocation {
	effortAllocations := make([]EffortAllocation, len(params))

	for i, v := range params {
		effortAllocations[i] = newEffortAllocation(v.Year, v.Month, v.Hours)
	}

	slices.SortStableFunc(effortAllocations, func(a, b EffortAllocation) int {
		return a.AllocationDate.Compare(b.AllocationDate)
	})

	return effortAllocations
}

func RestoreEffort(props RestoreEffortProps) *Effort {
	return &Effort{
		EffortID:          props.EffortID,
		BaselineID:        props.BaselineID,
		CompetenceID:      props.CompetenceID,
		Comment:           props.Comment,
		Hours:             props.Hours,
		EffortAllocations: props.EffortAllocations,
		CreatedAt:         props.CreatedAt,
		UpdatedAt:         props.UpdatedAt,
	}
}

func (e *Effort) Validate() error {
	err := common.Validate.Struct(e)
	if err != nil {
		return common.NewDomainValidationError(fmt.Errorf("effort domain validation failed: %w", err))
	}

	total := 0
	for _, v := range e.EffortAllocations {
		total += v.Hours
	}

	if total != e.Hours {
		return common.NewDomainValidationError(fmt.Errorf("effort allocation total %d is not equal to effort hours %d", total, e.Hours))
	}

	return nil
}

func (e *Effort) ChangeComment(commentStr *string) {
	if commentStr == nil {
		return
	}
	e.Comment = *commentStr
}

func (e *Effort) ChangeHours(hours *int) {
	if hours == nil {
		return
	}
	e.Hours = *hours
}

func (e *Effort) ChangeEffortAllocations(effortAllocations []EffortAllocationProps) {
	if effortAllocations == nil {
		return
	}
	e.EffortAllocations = createEffortAllocations(effortAllocations)
}
