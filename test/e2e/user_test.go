package e2e_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/repository"
	"github.com/celsopires1999/estimation/internal/testutils"
	"github.com/golang-migrate/migrate/v4"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
)

type UserE2ETestSuite struct {
	suite.Suite
	dbpool *pgxpool.Pool
	m      *migrate.Migrate
	repo   domain.EstimationRepository
}

func (s *UserE2ETestSuite) SetupSuite() {
	s.dbpool, s.m = testutils.DBSetup()
	s.repo = repository.NewEstimationRepositoryPostgres(s.dbpool)
}

func (s *UserE2ETestSuite) TearDownSuite() {
	defer s.dbpool.Close()
	err := s.m.Down()
	s.Nil(err)
}

func (s *UserE2ETestSuite) SetupSubTest() {
	err := testutils.TruncateTables(s.dbpool)
	if err != nil {
		s.T().Fatal(err)
	}
}

func TestE2EUser(t *testing.T) {
	suite.Run(t, new(UserE2ETestSuite))
}

func (s *UserE2ETestSuite) TestE2EUser() {
	s.Run("POST /api/v1/users", func() {
		ctx := context.Background()
		c := http.Client{}

		input := userInput{
			Email:    "john.doe@userland.com",
			UserName: "john1234",
			Name:     "John Doe",
			UserType: "manager",
		}

		message, err := json.Marshal(input)
		s.Nil(err)

		payload := bytes.NewReader(message)

		request, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:9000/api/v1/users", payload)
		if err != nil {
			s.T().Fatal(err)
		}
		response, err := c.Do(request)
		s.Nil(err)
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		s.Nil(err)

		s.Equal(http.StatusCreated, response.StatusCode)

		var output userOutput
		err = json.Unmarshal(body, &output)
		s.Nil(err)

		_, err = uuid.Parse(output.UserID)
		s.Nil(err, "UserID should be a valid uuidv4")
		s.Equal(input.Email, output.Email)
		s.Equal(input.UserName, output.UserName)
		s.Equal(input.Name, output.Name)
		s.Equal(input.UserType, output.UserType)
		createdAt, err := time.Parse(time.RFC3339, output.CreatedAt.Format(time.RFC3339))
		s.Nil(err)
		s.True(createdAt.After(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)))
		s.True(createdAt.Before(time.Now()))
		s.True(output.UpdatedAt.IsZero())
	})

	s.Run("PATCH /api/v1/users/:id", func() {
		ctx := context.Background()
		manager := s.arrangeManager(ctx)

		c := http.Client{}
		input := userInput{
			Email:    "john.doe@userland.com",
			UserName: "john1234",
			Name:     "John Doe",
			UserType: "manager",
		}

		message, err := json.Marshal(input)
		s.Nil(err)

		payload := bytes.NewReader(message)

		request, err := http.NewRequestWithContext(ctx, "PATCH", "http://localhost:9000/api/v1/users/"+manager.UserID, payload)
		if err != nil {
			s.T().Fatal(err)
		}
		response, err := c.Do(request)
		s.Nil(err)
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		s.Nil(err)

		s.Equal(http.StatusOK, response.StatusCode)

		var output userOutput
		err = json.Unmarshal(body, &output)
		s.Nil(err)

		_, err = uuid.Parse(output.UserID)
		s.Nil(err, "UserID should be a valid uuidv4")
		s.Equal(input.Email, output.Email)
		s.Equal(input.UserName, output.UserName)
		s.Equal(input.Name, output.Name)
		s.Equal(input.UserType, output.UserType)
		createdAt, err := time.Parse(time.RFC3339, output.CreatedAt.Format(time.RFC3339))
		s.Nil(err)
		s.True(createdAt.After(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)))
		s.True(createdAt.Before(time.Now()))
		updatedAt, err := time.Parse(time.RFC3339, output.UpdatedAt.Format(time.RFC3339))
		s.Nil(err)
		s.True(updatedAt.After(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)))
		s.True(updatedAt.Before(time.Now()))
	})

	s.Run("GET /api/v1/users/:id", func() {
		ctx := context.Background()
		manager := s.arrangeManager(ctx)

		c := http.Client{}

		request, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:9000/api/v1/users/"+manager.UserID, nil)
		if err != nil {
			s.T().Fatal(err)
		}
		response, err := c.Do(request)
		s.Nil(err)
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		s.Nil(err)

		s.Equal(http.StatusOK, response.StatusCode)

		var output userOutput
		err = json.Unmarshal(body, &output)
		s.Nil(err)
		s.Equal(manager.UserID, output.UserID)
		s.Equal(manager.Email, output.Email)
		s.Equal(manager.UserName, output.UserName)
		s.Equal(manager.Name, output.Name)
		s.Equal(manager.UserType.String(), output.UserType)
		createdAt, err := time.Parse(time.RFC3339, output.CreatedAt.Format(time.RFC3339))
		s.Nil(err)
		s.True(createdAt.After(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)))
		s.True(createdAt.Before(time.Now()))
		s.True(output.UpdatedAt.IsZero())
	})

	s.Run("DELETE /api/v1/users/:id", func() {
		ctx := context.Background()
		manager := s.arrangeManager(ctx)

		c := http.Client{}

		request, err := http.NewRequestWithContext(ctx, "DELETE", "http://localhost:9000/api/v1/users/"+manager.UserID, nil)

		if err != nil {
			s.T().Fatal(err)
		}
		response, err := c.Do(request)
		s.Nil(err)
		s.Equal(http.StatusNoContent, response.StatusCode)

	})
}

func (s *UserE2ETestSuite) arrangeManager(ctx context.Context) *domain.User {
	user := testutils.NewUserFakeBuilder().WithManager().Build()
	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		s.T().Fatal(err)
	}
	return user
}
