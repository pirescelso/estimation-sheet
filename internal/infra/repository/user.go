package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/jackc/pgx/v5"
)

func (r *estimationRepositoryPostgres) ValidateUser(ctx context.Context, userID string) (isManager bool, err error) {
	user, err := r.queries.FindUserById(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, common.NewNotFoundError(fmt.Errorf("user with id %s not found", userID))
		}
		return false, err
	}
	return user.UserType == "manager", err
}
