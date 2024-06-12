package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
)

type CostRepositoryPostgres struct {
	queries *db.Queries
}

func NewCostRepositoryPostgres(queries *db.Queries) *CostRepositoryPostgres {
	return &CostRepositoryPostgres{
		queries: queries,
	}
}

func (r *CostRepositoryPostgres) CreateCost(ctx context.Context, cost *domain.Cost) error {
	err := r.queries.CreateCost(ctx, db.CreateCostParams{
		CostID:      cost.CostID,
		ProjectID:   cost.ProjectID,
		CostType:    string(cost.CostType),
		Description: cost.Description,
		Comment:     pgtype.Text{String: cost.Comment, Valid: true},
		Amount:      cost.Amount,
		Currency:    string(cost.Currency),
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		return err
	}

	installments := db.CreateInstallmentsParams{}
	for _, installment := range cost.Installments {
		installments.Column1 = append(installments.Column1, uuid.New().String())
		installments.Column2 = append(installments.Column2, cost.CostID)
		installments.Column3 = append(installments.Column3, pgtype.Date{Time: installment.PaymentDate, Valid: true})
		installments.Column4 = append(installments.Column4, installment.Amount)
		installments.Column5 = append(installments.Column5, pgtype.Timestamp{Time: time.Now(), Valid: true})
	}
	err = r.queries.CreateInstallments(ctx, installments)

	if err != nil {
		return err
	}

	return nil
}

func (r *CostRepositoryPostgres) GetCost(ctx context.Context, costID string) (*domain.Cost, error) {
	costModel, err := r.queries.GetCost(ctx, costID)
	if err != nil {
		return nil, err
	}

	installmentModels, err := r.queries.GetInstallments(ctx, costID)
	if err != nil {
		return nil, err
	}

	installments := make([]domain.Installment, len(installmentModels))
	for i, installment := range installmentModels {
		installments[i] = domain.Installment{
			PaymentDate: installment.PaymentDate.Time,
			Amount:      installment.Amount,
		}
	}

	input := domain.RestoreCostProps{
		CostID:       costModel.CostID,
		ProjectID:    costModel.ProjectID,
		CostType:     domain.CostType(costModel.CostType),
		Description:  costModel.Description,
		Comment:      costModel.Comment.String,
		Amount:       costModel.Amount,
		Currency:     domain.Currency(costModel.Currency),
		Installments: installments,
		CreatedAt:    costModel.CreatedAt.Time,
		UpdatedAt:    costModel.UpdatedAt.Time,
	}

	cost := domain.RestoreCost(input)
	err = cost.Validate()
	if err != nil {
		return nil, err
	}

	return domain.RestoreCost(input), nil
}

func (r *CostRepositoryPostgres) UpdateCost(ctx context.Context, cost *domain.Cost) error {
	err := r.queries.UpdateCost(ctx, db.UpdateCostParams{
		CostID:      cost.CostID,
		ProjectID:   cost.ProjectID,
		Description: cost.Description,
		Comment:     pgtype.Text{String: cost.Comment, Valid: true},
		Amount:      cost.Amount,
		Currency:    string(cost.Currency),
		UpdatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	return err
}
