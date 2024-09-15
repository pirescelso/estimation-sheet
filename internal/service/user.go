package service

import (
	"context"

	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/celsopires1999/estimation/internal/mapper"
)

func (s *EstimationService) ListUsers(ctx context.Context, input ListUsersInputDTO) (*ListUsersOutputDTO, error) {
	users, metadata, err := s.queries.SearchUsers(ctx, input.Name, input.Filters)
	if err != nil {
		return nil, err
	}

	usersOutput := make([]mapper.UserOutput, len(users))
	for i, user := range users {
		usersOutput[i] = mapper.UserOutputFromDb(user)
	}

	return &ListUsersOutputDTO{metadata, usersOutput}, nil
}

type ListUsersInputDTO struct {
	Name string
	db.Filters
}

type ListUsersOutputDTO struct {
	Metadata db.Metadata         `json:"metadata"`
	Users    []mapper.UserOutput `json:"users"`
}
