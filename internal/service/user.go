package service

import (
	"context"

	"github.com/celsopires1999/estimation/internal/mapper"
)

func (s *EstimationService) ListUsers(ctx context.Context, input ListUsersInputDTO) (*ListUsersOutputDTO, error) {
	users, err := s.queries.FindAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	usersOutput := make([]mapper.UserOutput, len(users))
	for i, user := range users {
		usersOutput[i] = mapper.UserOutputFromDb(user)
	}

	return &ListUsersOutputDTO{usersOutput}, nil
}

type ListUsersInputDTO struct{}

type ListUsersOutputDTO struct {
	Users []mapper.UserOutput `json:"users"`
}
