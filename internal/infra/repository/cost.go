package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
)

func (r *estimationRepositoryPostgres) CreateCost(ctx context.Context, cost *domain.Cost) error {
	err := r.queries.InsertCost(ctx, db.InsertCostParams{
		CostID:      cost.CostID,
		BaselineID:  cost.BaselineID,
		CostType:    string(cost.CostType),
		Description: cost.Description,
		Comment:     pgtype.Text{String: cost.Comment, Valid: true},
		Amount:      cost.Amount,
		Currency:    string(cost.Currency),
		Tax:         cost.Tax,
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		return costCheckRelationsError(cost, err)
	}

	costAllocations := db.BulkInsertCostAllocationParams{}
	for _, allocation := range cost.CostAllocations {
		costAllocations.Column1 = append(costAllocations.Column1, uuid.New().String())
		costAllocations.Column2 = append(costAllocations.Column2, cost.CostID)
		costAllocations.Column3 = append(costAllocations.Column3, pgtype.Date{Time: allocation.AllocationDate, Valid: true})
		costAllocations.Column4 = append(costAllocations.Column4, allocation.Amount)
		costAllocations.Column5 = append(costAllocations.Column5, pgtype.Timestamp{Time: time.Now(), Valid: true})
	}
	err = r.queries.BulkInsertCostAllocation(ctx, costAllocations)

	if err != nil {
		return err
	}

	return nil
}

func (r *estimationRepositoryPostgres) CreateCostMany(ctx context.Context, costs []*domain.Cost) error {
	costsParams := db.BulkInsertCostParams{}
	costAllocations := db.BulkInsertCostAllocationParams{}

	for _, cost := range costs {
		costsParams.Column1 = append(costsParams.Column1, cost.CostID)
		costsParams.Column2 = append(costsParams.Column2, cost.BaselineID)
		costsParams.Column3 = append(costsParams.Column3, string(cost.CostType))
		costsParams.Column4 = append(costsParams.Column4, cost.Description)
		costsParams.Column5 = append(costsParams.Column5, cost.Comment)
		costsParams.Column6 = append(costsParams.Column6, cost.Amount)
		costsParams.Column7 = append(costsParams.Column7, string(cost.Currency))
		costsParams.Column8 = append(costsParams.Column8, cost.Tax)
		costsParams.Column9 = append(costsParams.Column9, pgtype.Timestamp{Time: time.Now(), Valid: true})

		for _, allocation := range cost.CostAllocations {
			costAllocations.Column1 = append(costAllocations.Column1, uuid.New().String())
			costAllocations.Column2 = append(costAllocations.Column2, cost.CostID)
			costAllocations.Column3 = append(costAllocations.Column3, pgtype.Date{Time: allocation.AllocationDate, Valid: true})
			costAllocations.Column4 = append(costAllocations.Column4, allocation.Amount)
			costAllocations.Column5 = append(costAllocations.Column5, pgtype.Timestamp{Time: time.Now(), Valid: true})
		}
	}
	err := r.queries.BulkInsertCost(ctx, costsParams)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return common.NewConflictError(fmt.Errorf("duplicated cost on creating many costs: %w", err))
			}
			return common.NewConflictError(err)
		}

		return err
	}
	err = r.queries.BulkInsertCostAllocation(ctx, costAllocations)
	if err != nil {
		return err
	}

	return nil
}

func (r *estimationRepositoryPostgres) GetCost(ctx context.Context, costID string) (*domain.Cost, error) {
	costModel, err := r.queries.FindCostById(ctx, costID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, common.NewNotFoundError(errors.New("cost not found"))
		}
		return nil, common.NewNotFoundError(fmt.Errorf("cannot find cost with id %s: %w", costID, err))
	}

	allocationModels, err := r.queries.FindCostAllocationsByCostId(ctx, costID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, common.NewNotFoundError(errors.New("cost allocations not found"))
		}
		return nil, common.NewNotFoundError(fmt.Errorf("cannot find cost allocations with id %s: %w", costID, err))
	}

	allocations := make([]domain.CostAllocation, len(allocationModels))
	for i, allocation := range allocationModels {
		allocations[i] = domain.CostAllocation{
			AllocationDate: allocation.AllocationDate.Time,
			Amount:         allocation.Amount,
		}
	}

	props := domain.RestoreCostProps{
		CostID:          costModel.CostID,
		BaselineID:      costModel.BaselineID,
		CostType:        domain.CostType(costModel.CostType),
		Description:     costModel.Description,
		Comment:         costModel.Comment.String,
		Amount:          costModel.Amount,
		Currency:        domain.Currency(costModel.Currency),
		Tax:             costModel.Tax,
		CostAllocations: allocations,
		CreatedAt:       costModel.CreatedAt.Time,
		UpdatedAt:       costModel.UpdatedAt.Time,
	}

	cost := domain.RestoreCost(props)
	err = cost.Validate()
	if err != nil {
		return nil, err
	}

	return cost, nil
}

func (r *estimationRepositoryPostgres) UpdateCost(ctx context.Context, cost *domain.Cost) error {
	err := r.queries.UpdateCost(ctx, db.UpdateCostParams{
		CostID:      cost.CostID,
		BaselineID:  cost.BaselineID,
		CostType:    string(cost.CostType),
		Description: cost.Description,
		Comment:     pgtype.Text{String: cost.Comment, Valid: true},
		Amount:      cost.Amount,
		Currency:    string(cost.Currency),
		Tax:         cost.Tax,
		UpdatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		return costCheckRelationsError(cost, err)
	}

	_, err = r.queries.DeleteCostAllocations(ctx, cost.CostID)

	if err != nil {
		return err
	}

	costAllocations := db.BulkInsertCostAllocationParams{}
	for _, allocation := range cost.CostAllocations {
		costAllocations.Column1 = append(costAllocations.Column1, uuid.New().String())
		costAllocations.Column2 = append(costAllocations.Column2, cost.CostID)
		costAllocations.Column3 = append(costAllocations.Column3, pgtype.Date{Time: allocation.AllocationDate, Valid: true})
		costAllocations.Column4 = append(costAllocations.Column4, allocation.Amount)
		costAllocations.Column5 = append(costAllocations.Column5, pgtype.Timestamp{Time: time.Now(), Valid: true})
	}
	err = r.queries.BulkInsertCostAllocation(ctx, costAllocations)

	return err
}

func (r *estimationRepositoryPostgres) DeleteCost(ctx context.Context, costID string) error {
	_, err := r.queries.DeleteCostAllocations(ctx, costID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return common.NewNotFoundError(errors.New("cost allocations not found"))
		}
		return common.NewConflictError(fmt.Errorf("cannot delete cost allocations wiht cost id %s: %w", costID, err))
	}
	_, err = r.queries.DeleteCost(ctx, costID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return common.NewNotFoundError(errors.New("cost not found"))
		}
		return common.NewConflictError(fmt.Errorf("cannot delete cost id %s: %w", costID, err))
	}

	return err
}

func (r *estimationRepositoryPostgres) GetCostManyByBaselineID(ctx context.Context, baselineID string) ([]*domain.Cost, error) {
	costModels, err := r.queries.FindCostsByBaselineId(ctx, baselineID)
	if err != nil {
		return nil, err
	}

	costs := make([]*domain.Cost, len(costModels))
	for i, costModel := range costModels {
		allocations, err := r.queries.FindCostAllocationsByCostId(ctx, costModel.CostID)
		if err != nil {
			return nil, err
		}

		allocs := make([]domain.CostAllocation, len(allocations))
		for j, allocation := range allocations {
			allocs[j] = domain.CostAllocation{
				AllocationDate: allocation.AllocationDate.Time,
				Amount:         allocation.Amount,
			}
		}

		props := domain.RestoreCostProps{
			CostID:          costModel.CostID,
			BaselineID:      costModel.BaselineID,
			CostType:        domain.CostType(costModel.CostType),
			Description:     costModel.Description,
			Comment:         costModel.Comment.String,
			Amount:          costModel.Amount,
			Currency:        domain.Currency(costModel.Currency),
			Tax:             costModel.Tax,
			CostAllocations: allocs,
			CreatedAt:       costModel.CreatedAt.Time,
			UpdatedAt:       costModel.UpdatedAt.Time,
		}

		costs[i] = domain.RestoreCost(props)
		err = costs[i].Validate()
		if err != nil {
			return nil, err
		}
	}

	return costs, nil
}

func costCheckRelationsError(cost *domain.Cost, err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			return common.NewConflictError(fmt.Errorf("cost type: '%s' description: '%s' already exists in the baseline id: '%s'", string(cost.CostType), cost.Description, cost.BaselineID))
		}
		return common.NewConflictError(err)
	}

	return err
}
