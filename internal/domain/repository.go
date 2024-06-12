package domain

import "context"

type CostRepository interface {
	CreateCost(ctx context.Context, cost *Cost) error
	GetCost(ctx context.Context, costID string) (*Cost, error)
	UpdateCost(ctx context.Context, cost *Cost) error
}
