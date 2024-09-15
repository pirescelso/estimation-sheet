package db_test

import (
	"context"
	"testing"
	"time"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/celsopires1999/estimation/internal/testutils"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type UserDBTestSuite struct {
	suite.Suite
	dbpool  *pgxpool.Pool
	m       *migrate.Migrate
	queries *db.Queries
}

func (s *UserDBTestSuite) SetupSuite() {
	s.dbpool, s.m = testutils.DBSetup()
	s.queries = db.New(s.dbpool)
}

func (s *UserDBTestSuite) TearDownSuite() {
	defer s.dbpool.Close()
	err := s.m.Down()
	s.Nil(err)
}

func (s *UserDBTestSuite) SetupSubTest() {
	err := testutils.TruncateTables(s.dbpool)
	if err != nil {
		s.T().Fatal(err)
	}
}

func TestIntegrationUserDB(t *testing.T) {
	suite.Run(t, new(UserDBTestSuite))
}

func (s *UserDBTestSuite) TestSearchUsers() {
	type userTestCase struct {
		test    string
		name    string
		filters db.Filters
	}

	tests := []userTestCase{
		{
			test: "should search users without name returning all 3 users",
			name: "",
			filters: db.Filters{
				Page:         1,
				PageSize:     9,
				Sort:         "name",
				SortSafelist: []string{"name"},
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.test, func() {
			ctx := context.Background()
			var users = make([]domain.User, 3)
			for i := range len(users) {
				users[i] = *testutils.NewUserFakeBuilder().Build()
			}

			s.arrangeUsers(ctx, users)

			foundUsers, _, err := s.queries.SearchUsers(ctx, tt.name, tt.filters)

			s.Nil(err)
			s.Equal(len(users), len(foundUsers))
		})
	}
}

func (s *UserDBTestSuite) TestSearchUserByName() {
	type userTestCase struct {
		test    string
		name    string
		filters db.Filters
	}

	tests := []userTestCase{
		{
			test: "should search users with name = John Doe",
			name: "John Doe",
			filters: db.Filters{
				Page:         1,
				PageSize:     1,
				Sort:         "name",
				SortSafelist: []string{"name"},
			},
		},
		{
			test: "should search users with name = john%",
			name: "john%",
			filters: db.Filters{
				Page:         1,
				PageSize:     1,
				Sort:         "name",
				SortSafelist: []string{"name"},
			},
		},
		{
			test: "should search users with name = %john%",
			name: "%john%",
			filters: db.Filters{
				Page:         1,
				PageSize:     1,
				Sort:         "name",
				SortSafelist: []string{"name"},
			},
		},
		{
			test: "should search users with name = %john doe%",
			name: "%john%",
			filters: db.Filters{
				Page:         1,
				PageSize:     1,
				Sort:         "name",
				SortSafelist: []string{"name"},
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.test, func() {
			ctx := context.Background()
			var users = make([]domain.User, 2)
			users[0] = *testutils.NewUserFakeBuilder().WithName("John Doe").Build()
			users[1] = *testutils.NewUserFakeBuilder().WithName("Mary Doe").Build()
			s.arrangeUsers(ctx, users)

			foundUsers, _, err := s.queries.SearchUsers(ctx, tt.name, tt.filters)

			s.Nil(err)
			s.Equal(1, len(foundUsers))
			s.Equal(users[0].UserID, foundUsers[0].UserID)
			s.Equal(users[0].Email, foundUsers[0].Email)
			s.Equal(users[0].UserName, foundUsers[0].UserName)
			s.Equal(users[0].Name, foundUsers[0].Name)
		})
	}
}

func (s *UserDBTestSuite) arrangeUsers(ctx context.Context, users []domain.User) {
	for _, user := range users {
		err := s.queries.InsertUser(ctx, db.InsertUserParams{
			UserID:    user.UserID,
			Email:     user.Email,
			UserName:  user.UserName,
			Name:      user.Name,
			UserType:  user.UserType.String(),
			CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		})

		if err != nil {
			s.T().Fatal(err)
		}
	}
}
