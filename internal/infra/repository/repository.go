package repository

import (
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type estimationRepositoryPostgres struct {
	queries *db.Queries
}

func NewEstimationRepositoryPostgres(dbpool *pgxpool.Pool) *estimationRepositoryPostgres {
	return &estimationRepositoryPostgres{
		queries: db.New(dbpool),
	}
}

func NewEstimationRepositoryTxmPostgres(q *db.Queries) *estimationRepositoryPostgres {
	return &estimationRepositoryPostgres{
		queries: q,
	}
}
