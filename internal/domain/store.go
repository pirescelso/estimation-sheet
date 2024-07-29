package domain

import "context"

type Store interface {
	ValidateUser(ctx context.Context, userID string) (isManager bool, err error)
	ValidatePlan(ctx context.Context, planID string) error
}
