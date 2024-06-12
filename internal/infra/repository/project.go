package repository

import (
	"context"
	"time"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type ProjectRepositoryPostgres struct {
	queries *db.Queries
}

func NewProjectRepositoryPostgres(q *db.Queries) *ProjectRepositoryPostgres {
	return &ProjectRepositoryPostgres{
		queries: q,
	}
}

func (r *ProjectRepositoryPostgres) CreateProject(ctx context.Context, project *domain.Project) error {
	err := r.queries.CreateProject(ctx, db.CreateProjectParams{
		ProjectID:   project.ProjectID,
		Description: "TBD",
		StartDate:   pgtype.Date{Time: project.StarDate, Valid: true},
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})

	return err
}
