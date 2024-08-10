package domain

import (
	"fmt"
	"time"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/google/uuid"
)

type Workload struct {
	WorkloadID          string
	PortfolioID         string
	EffortID            string
	Hours               int
	WorkloadAllocations []WorkloadAllocation
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type RestoreWorkloadProps Workload
type NewWorkloadAllocationProps struct {
	Year  int
	Month time.Month
	Hours int
}

type NewWorkloadProps struct {
	PortfolioID         string
	EffortID            string
	Hours               int
	WorkloadAllocations []NewWorkloadAllocationProps
}

func NewWorkload(props NewWorkloadProps) *Workload {
	workloadAllocations := createWorkloadAllocation(props.WorkloadAllocations)
	return &Workload{
		WorkloadID:          uuid.NewString(),
		PortfolioID:         props.PortfolioID,
		EffortID:            props.EffortID,
		Hours:               props.Hours,
		WorkloadAllocations: workloadAllocations,
	}
}

func createWorkloadAllocation(params []NewWorkloadAllocationProps) []WorkloadAllocation {
	workloadAllocations := make([]WorkloadAllocation, len(params))

	for i, v := range params {
		workloadAllocations[i] = NewWorkloadAllocation(v.Year, v.Month, v.Hours)
	}

	return workloadAllocations
}

func RestoreWorkload(props RestoreWorkloadProps) *Workload {
	return &Workload{
		WorkloadID:          props.WorkloadID,
		PortfolioID:         props.PortfolioID,
		EffortID:            props.EffortID,
		Hours:               props.Hours,
		WorkloadAllocations: props.WorkloadAllocations,
		CreatedAt:           props.CreatedAt,
		UpdatedAt:           props.UpdatedAt,
	}
}

func (w *Workload) Validate() error {
	if w.Hours <= 0 {
		return common.NewDomainValidationError(fmt.Errorf("invalid workload hours %d", w.Hours))
	}

	total := 0
	for _, v := range w.WorkloadAllocations {
		total += v.Hours
	}
	if total != w.Hours {
		return common.NewDomainValidationError(fmt.Errorf("workload allocation total %d is not equal to workload hours %d", total, w.Hours))
	}

	return nil
}

func (w *Workload) GetWorloadAllocation() []WorkloadAllocation {
	return w.WorkloadAllocations
}

type WorkloadAllocation struct {
	AllocationDate time.Time
	Hours          int
}

func NewWorkloadAllocation(year int, month time.Month, hours int) WorkloadAllocation {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return WorkloadAllocation{
		AllocationDate: date,
		Hours:          hours,
	}
}
