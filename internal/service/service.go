package service

import (
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EstimationService struct {
	queries *db.Queries
}

func NewEstimationService(dbpool *pgxpool.Pool) *EstimationService {
	return &EstimationService{
		queries: db.New(dbpool),
	}
}
