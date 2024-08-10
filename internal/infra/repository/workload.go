package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
)

func (r *estimationRepositoryPostgres) CreateWorkload(ctx context.Context, workload *domain.Workload) error {
	err := r.queries.InsertWorkload(ctx, db.InsertWorkloadParams{
		WorkloadID:  workload.WorkloadID,
		PortfolioID: workload.PortfolioID,
		EffortID:    workload.EffortID,
		Hours:       int32(workload.Hours),
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		return err
	}

	workloadAllocations := db.BulkInsertWorkloadAllocationParams{}
	for _, allocation := range workload.WorkloadAllocations {
		workloadAllocations.Column1 = append(workloadAllocations.Column1, uuid.NewString())
		workloadAllocations.Column2 = append(workloadAllocations.Column2, workload.WorkloadID)
		workloadAllocations.Column3 = append(workloadAllocations.Column3, pgtype.Date{Time: allocation.AllocationDate, Valid: true})
		workloadAllocations.Column4 = append(workloadAllocations.Column4, int32(allocation.Hours))
		workloadAllocations.Column5 = append(workloadAllocations.Column5, pgtype.Timestamp{Time: time.Now(), Valid: true})
	}
	err = r.queries.BulkInsertWorkloadAllocation(ctx, workloadAllocations)

	if err != nil {
		return err
	}

	return nil
}

func (r *estimationRepositoryPostgres) CreateWorkloadMany(ctx context.Context, workloads []*domain.Workload) error {
	workloadsParams := db.BulkInsertWorkloadParams{}
	workloadAllocations := db.BulkInsertWorkloadAllocationParams{}

	for _, workload := range workloads {
		workloadsParams.Column1 = append(workloadsParams.Column1, workload.WorkloadID)
		workloadsParams.Column2 = append(workloadsParams.Column2, workload.PortfolioID)
		workloadsParams.Column3 = append(workloadsParams.Column3, workload.EffortID)
		workloadsParams.Column4 = append(workloadsParams.Column4, int32(workload.Hours))
		workloadsParams.Column5 = append(workloadsParams.Column5, pgtype.Timestamp{Time: time.Now(), Valid: true})
		for _, allocation := range workload.WorkloadAllocations {
			workloadAllocations.Column1 = append(workloadAllocations.Column1, uuid.NewString())
			workloadAllocations.Column2 = append(workloadAllocations.Column2, workload.WorkloadID)
			workloadAllocations.Column3 = append(workloadAllocations.Column3, pgtype.Date{Time: allocation.AllocationDate, Valid: true})
			workloadAllocations.Column4 = append(workloadAllocations.Column4, int32(allocation.Hours))
			workloadAllocations.Column5 = append(workloadAllocations.Column5, pgtype.Timestamp{Time: time.Now(), Valid: true})
		}
	}
	err := r.queries.BulkInsertWorkload(ctx, workloadsParams)
	if err != nil {
		return err
	}
	err = r.queries.BulkInsertWorkloadAllocation(ctx, workloadAllocations)
	if err != nil {
		return err
	}

	return nil
}

func (r *estimationRepositoryPostgres) GetWorkload(ctx context.Context, workloadID string) (*domain.Workload, error) {
	workloadModel, err := r.queries.FindWorkloadById(ctx, workloadID)
	if err != nil {
		return nil, err
	}

	allocationModels, err := r.queries.FindWorkloadAllocations(ctx, workloadID)
	if err != nil {
		return nil, err
	}

	allocations := make([]domain.WorkloadAllocation, len(allocationModels))
	for i, allocation := range allocationModels {
		allocations[i] = domain.WorkloadAllocation{
			AllocationDate: allocation.AllocationDate.Time,
			Hours:          int(allocation.Hours),
		}
	}

	props := domain.RestoreWorkloadProps{
		WorkloadID:          workloadModel.WorkloadID,
		PortfolioID:         workloadModel.PortfolioID,
		EffortID:            workloadModel.EffortID,
		Hours:               int(workloadModel.Hours),
		WorkloadAllocations: allocations,
		CreatedAt:           workloadModel.CreatedAt.Time,
		UpdatedAt:           workloadModel.UpdatedAt.Time,
	}

	workload := domain.RestoreWorkload(props)
	err = workload.Validate()
	if err != nil {
		return nil, err
	}

	return workload, nil
}

func (r *estimationRepositoryPostgres) UpdateWorkload(ctx context.Context, workload *domain.Workload) error {
	err := r.queries.UpdateWorkload(ctx, db.UpdateWorkloadParams{
		WorkloadID:  workload.WorkloadID,
		PortfolioID: workload.PortfolioID,
		EffortID:    workload.EffortID,
		Hours:       int32(workload.Hours),
		UpdatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		return err
	}

	_, err = r.queries.DeleteWorkloadAllocations(ctx, workload.WorkloadID)

	if err != nil {
		return err
	}

	workloadAllocations := db.BulkInsertWorkloadAllocationParams{}
	for _, allocation := range workload.WorkloadAllocations {
		workloadAllocations.Column1 = append(workloadAllocations.Column1, uuid.NewString())
		workloadAllocations.Column2 = append(workloadAllocations.Column2, workload.WorkloadID)
		workloadAllocations.Column3 = append(workloadAllocations.Column3, pgtype.Date{Time: allocation.AllocationDate, Valid: true})
		workloadAllocations.Column4 = append(workloadAllocations.Column4, int32(allocation.Hours))
		workloadAllocations.Column5 = append(workloadAllocations.Column5, pgtype.Timestamp{Time: time.Now(), Valid: true})
	}
	err = r.queries.BulkInsertWorkloadAllocation(ctx, workloadAllocations)

	return err
}

func (r *estimationRepositoryPostgres) DeleteWorkload(ctx context.Context, workloadID string) error {
	_, err := r.queries.DeleteWorkloadAllocations(ctx, workloadID)
	if err != nil {
		return err
	}

	_, err = r.queries.DeleteWorkload(ctx, workloadID)

	return err
}

func (r *estimationRepositoryPostgres) DeleteWorkloadsByPortfolioID(ctx context.Context, portfolioID string) error {
	workloadModels, err := r.queries.FindWorkloadsByPortfolioId(ctx, portfolioID)
	if err != nil {
		return err
	}
	for _, workloadModel := range workloadModels {
		_, err = r.queries.DeleteWorkloadAllocations(ctx, workloadModel.WorkloadID)
		if err != nil {
			return err
		}
	}
	_, err = r.queries.DeleteWorkloadsByPortfolioId(ctx, portfolioID)
	if err != nil {
		return err
	}

	return nil
}

func (r *estimationRepositoryPostgres) GetWorkloadManyByPortfolioID(ctx context.Context, portfolioID string) ([]*domain.Workload, error) {
	workloadModels, err := r.queries.FindWorkloadsByPortfolioId(ctx, portfolioID)
	if err != nil {
		return nil, err
	}

	workloads := make([]*domain.Workload, len(workloadModels))
	for i, workloadModel := range workloadModels {
		allocations, err := r.queries.FindWorkloadAllocations(ctx, workloadModel.WorkloadID)
		if err != nil {
			return nil, err
		}

		allocs := make([]domain.WorkloadAllocation, len(allocations))
		for j, allocation := range allocations {
			allocs[j] = domain.WorkloadAllocation{
				AllocationDate: allocation.AllocationDate.Time,
				Hours:          int(allocation.Hours),
			}
		}

		props := domain.RestoreWorkloadProps{
			WorkloadID:          workloadModel.WorkloadID,
			PortfolioID:         workloadModel.PortfolioID,
			EffortID:            workloadModel.EffortID,
			Hours:               int(workloadModel.Hours),
			WorkloadAllocations: allocs,
			CreatedAt:           workloadModel.CreatedAt.Time,
			UpdatedAt:           workloadModel.UpdatedAt.Time,
		}

		workloads[i] = domain.RestoreWorkload(props)
		err = workloads[i].Validate()
		if err != nil {
			return nil, err
		}
	}

	return workloads, nil
}
