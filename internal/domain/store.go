package domain

import "context"

type Store interface {
	ValidatePlan(ctx context.Context, planID string) error
}
