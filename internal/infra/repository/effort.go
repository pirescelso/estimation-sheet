package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *estimationRepositoryPostgres) CreateEffort(ctx context.Context, effort *domain.Effort) error {
	err := r.queries.InsertEffort(ctx, db.InsertEffortParams{
		EffortID:     effort.EffortID,
		BaselineID:   effort.BaselineID,
		CompetenceID: effort.CompetenceID,
		Comment:      pgtype.Text{String: effort.Comment, Valid: true},
		Hours:        int32(effort.Hours),
		CreatedAt:    pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		return effortCheckRelationsError(effort, err)
	}

	effortAllocations := db.BulkInsertEffortAllocationParams{}

	for _, allocation := range effort.EffortAllocations {
		effortAllocations.Column1 = append(effortAllocations.Column1, uuid.New().String())
		effortAllocations.Column2 = append(effortAllocations.Column2, effort.EffortID)
		effortAllocations.Column3 = append(effortAllocations.Column3, pgtype.Date{Time: allocation.AllocationDate, Valid: true})
		effortAllocations.Column4 = append(effortAllocations.Column4, int32(allocation.Hours))
		effortAllocations.Column5 = append(effortAllocations.Column5, pgtype.Timestamp{Time: time.Now(), Valid: true})
	}

	err = r.queries.BulkInsertEffortAllocation(ctx, effortAllocations)
	if err != nil {
		return err
	}

	return nil
}

func effortCheckRelationsError(effort *domain.Effort, err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			return common.NewConflictError(fmt.Errorf("competence: '%s' already exists in the baseline id: '%s'", effort.CompetenceID, effort.BaselineID))
		}
		return common.NewConflictError(err)
	}

	return err
}

func (r *estimationRepositoryPostgres) CreateEffortMany(ctx context.Context, efforts []*domain.Effort) error {
	effortsParams := db.BulkInsertEffortParams{}
	effortAllocations := db.BulkInsertEffortAllocationParams{}

	for _, effort := range efforts {
		effortsParams.Column1 = append(effortsParams.Column1, effort.EffortID)
		effortsParams.Column2 = append(effortsParams.Column2, effort.BaselineID)
		effortsParams.Column3 = append(effortsParams.Column3, effort.CompetenceID)
		effortsParams.Column4 = append(effortsParams.Column4, effort.Comment)
		effortsParams.Column5 = append(effortsParams.Column5, int32(effort.Hours))
		effortsParams.Column6 = append(effortsParams.Column6, pgtype.Timestamp{Time: time.Now(), Valid: true})

		for _, allocation := range effort.EffortAllocations {
			effortAllocations.Column1 = append(effortAllocations.Column1, uuid.New().String())
			effortAllocations.Column2 = append(effortAllocations.Column2, effort.EffortID)
			effortAllocations.Column3 = append(effortAllocations.Column3, pgtype.Date{Time: allocation.AllocationDate, Valid: true})
			effortAllocations.Column4 = append(effortAllocations.Column4, int32(allocation.Hours))
			effortAllocations.Column5 = append(effortAllocations.Column5, pgtype.Timestamp{Time: time.Now(), Valid: true})
		}
	}

	err := r.queries.BulkInsertEffort(ctx, effortsParams)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return common.NewConflictError(fmt.Errorf("duplicated effort on creating many efforts: %w", err))
			}
			return common.NewConflictError(err)
		}

		return err
	}

	err = r.queries.BulkInsertEffortAllocation(ctx, effortAllocations)
	if err != nil {
		return err
	}

	return nil
}

func (r *estimationRepositoryPostgres) GetEffort(ctx context.Context, effortID string) (*domain.Effort, error) {
	effortModel, err := r.queries.FindEffortById(ctx, effortID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, common.NewNotFoundError(errors.New("effort not found"))
		}
		return nil, common.NewNotFoundError(fmt.Errorf("cannot find effort with id %s: %w", effortID, err))
	}

	allocationModels, err := r.queries.FindEffortAllocationsByEffortId(ctx, effortID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, common.NewNotFoundError(errors.New("effort allocations not found"))
		}
		return nil, common.NewNotFoundError(fmt.Errorf("cannot find effort allocations with id %s: %w", effortID, err))
	}

	allocations := make([]domain.EffortAllocation, len(allocationModels))
	for i, allocationModel := range allocationModels {
		allocations[i] = domain.EffortAllocation{
			AllocationDate: allocationModel.AllocationDate.Time,
			Hours:          int(allocationModel.Hours),
		}
	}

	props := domain.RestoreEffortProps{
		EffortID:          effortModel.EffortID,
		BaselineID:        effortModel.BaselineID,
		CompetenceID:      effortModel.CompetenceID,
		Comment:           effortModel.Comment.String,
		Hours:             int(effortModel.Hours),
		EffortAllocations: allocations,
		CreatedAt:         effortModel.CreatedAt.Time,
		UpdatedAt:         effortModel.UpdatedAt.Time,
	}

	effort := domain.RestoreEffort(props)
	err = effort.Validate()
	if err != nil {
		return nil, err
	}
	return effort, nil
}

func (r *estimationRepositoryPostgres) UpdateEffort(ctx context.Context, effort *domain.Effort) error {
	err := r.queries.UpdateEffort(ctx, db.UpdateEffortParams{
		EffortID:     effort.EffortID,
		BaselineID:   effort.BaselineID,
		CompetenceID: effort.CompetenceID,
		Comment:      pgtype.Text{String: effort.Comment, Valid: true},
		Hours:        int32(effort.Hours),
		UpdatedAt:    pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		return effortCheckRelationsError(effort, err)
	}

	_, err = r.queries.DeleteEffortAllocations(ctx, effort.EffortID)

	if err != nil {
		return err
	}

	effortAllocations := db.BulkInsertEffortAllocationParams{}

	for _, allocation := range effort.EffortAllocations {
		effortAllocations.Column1 = append(effortAllocations.Column1, uuid.New().String())
		effortAllocations.Column2 = append(effortAllocations.Column2, effort.EffortID)
		effortAllocations.Column3 = append(effortAllocations.Column3, pgtype.Date{Time: allocation.AllocationDate, Valid: true})
		effortAllocations.Column4 = append(effortAllocations.Column4, int32(allocation.Hours))
		effortAllocations.Column5 = append(effortAllocations.Column5, pgtype.Timestamp{Time: time.Now(), Valid: true})
	}

	err = r.queries.BulkInsertEffortAllocation(ctx, effortAllocations)

	return err
}

func (r *estimationRepositoryPostgres) DeleteEffort(ctx context.Context, effortID string) error {
	_, err := r.queries.DeleteEffortAllocations(ctx, effortID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return common.NewNotFoundError(errors.New("effort allocations not found"))
		}
		return common.NewConflictError(fmt.Errorf("cannot delete effort allocations wiht effort id %s: %w", effortID, err))
	}
	_, err = r.queries.DeleteEffort(ctx, effortID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return common.NewNotFoundError(errors.New("effort not found"))
		}
		return common.NewConflictError(fmt.Errorf("cannot delete effort id %s: %w", effortID, err))
	}

	return err
}

func (r *estimationRepositoryPostgres) GetEffortManyByBaselineID(ctx context.Context, baselineID string) ([]*domain.Effort, error) {
	effortModels, err := r.queries.FindEffortsByBaselineIdWithRelations(ctx, baselineID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, common.NewNotFoundError(errors.New("efforts not found"))
		}
		return nil, common.NewNotFoundError(fmt.Errorf("cannot find efforts with baseline id %s: %w", baselineID, err))
	}

	efforts := make([]*domain.Effort, len(effortModels))
	for i, effortModel := range effortModels {
		allocations, err := r.queries.FindEffortAllocationsByEffortId(ctx, effortModel.EffortID)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return nil, common.NewNotFoundError(errors.New("effort allocations not found"))
			}
			return nil, common.NewNotFoundError(fmt.Errorf("cannot find effort allocations with id %s: %w", effortModel.EffortID, err))
		}

		allocs := make([]domain.EffortAllocation, len(allocations))
		for j, allocation := range allocations {
			allocs[j] = domain.EffortAllocation{
				AllocationDate: allocation.AllocationDate.Time,
				Hours:          int(allocation.Hours),
			}
		}

		props := domain.RestoreEffortProps{
			EffortID:          effortModel.EffortID,
			BaselineID:        effortModel.BaselineID,
			CompetenceID:      effortModel.CompetenceID,
			Comment:           effortModel.Comment.String,
			Hours:             int(effortModel.Hours),
			EffortAllocations: allocs,
			CreatedAt:         effortModel.CreatedAt.Time,
			UpdatedAt:         effortModel.UpdatedAt.Time,
		}

		efforts[i] = domain.RestoreEffort(props)
		err = efforts[i].Validate()

		if err != nil {
			return nil, err
		}
	}

	return efforts, nil
}
