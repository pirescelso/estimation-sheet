package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *estimationRepositoryPostgres) CreateCompetence(ctx context.Context, competence *domain.Competence) error {
	err := r.queries.InsertCompetence(ctx, db.InsertCompetenceParams{
		CompetenceID: competence.CompetenceID,
		Code:         competence.Code,
		Name:         competence.Name,
		CreatedAt:    pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return common.NewConflictError(fmt.Errorf("competence code %s already exists", competence.Code))
			}
			return common.NewConflictError(err)
		}
	}

	return err
}

func (r *estimationRepositoryPostgres) GetCompetence(ctx context.Context, competenceID string) (*domain.Competence, error) {
	competenceModel, err := r.queries.FindCompetenceById(ctx, competenceID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, common.NewNotFoundError(fmt.Errorf("competence with id %s not found", competenceID))
		}
		return nil, err
	}

	props := domain.RestoreCompetenceProps{
		CompetenceID: competenceModel.CompetenceID,
		Code:         competenceModel.Code,
		Name:         competenceModel.Name,
		CreatedAt:    competenceModel.CreatedAt.Time,
		UpdatedAt:    competenceModel.UpdatedAt.Time,
	}

	competence := domain.RestoreCompetence(props)
	err = competence.Validate()
	if err != nil {
		return nil, err
	}
	return competence, nil
}

func (r *estimationRepositoryPostgres) UpdateCompetence(ctx context.Context, competence *domain.Competence) error {
	err := r.queries.UpdateCompetence(ctx, db.UpdateCompetenceParams{
		CompetenceID: competence.CompetenceID,
		Code:         competence.Code,
		Name:         competence.Name,
		UpdatedAt:    pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return common.NewNotFoundError(fmt.Errorf("competence with id %s not found", competence.CompetenceID))
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return common.NewConflictError(fmt.Errorf("competence code %s already exists", competence.Code))
			}
			return common.NewConflictError(err)
		}
	}

	return err
}

func (r *estimationRepositoryPostgres) DeleteCompetence(ctx context.Context, competenceID string) error {
	_, err := r.queries.DeleteCompetence(ctx, competenceID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return common.NewNotFoundError(fmt.Errorf("competence with id %s not found", competenceID))
		}
		return err
	}
	return nil
}
