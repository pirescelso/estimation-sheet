package usecase_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/celsopires1999/estimation/internal/infra/repository"
	"github.com/celsopires1999/estimation/internal/usecase"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/suite"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	MIGRATION_PATH = "file://../../sql/migrations"
)

type CostUsecaseTestSuite struct {
	suite.Suite
	conn    *pgx.Conn
	m       *migrate.Migrate
	project *domain.Project
}

func (s *CostUsecaseTestSuite) SetupSuite() {
	dbURL := "postgres://postgres:postgres@db:5432/postgres?sslmode=disable"
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}

	m, err := migrate.New(MIGRATION_PATH, dbURL)
	s.Nil(err)
	s.NotNil(m)
	err = m.Up()
	s.Nil(err)
	s.conn = conn
	s.m = m

	projectRepo := repository.NewProjectRepositoryPostgres(db.New(conn))

	date, _ := time.Parse("02-01-2006", "01-01-2020")
	s.project = domain.NewProject(date)
	err = projectRepo.CreateProject(context.Background(), s.project)
	s.Nil(err)
}

func (s *CostUsecaseTestSuite) TearDownSuite() {
	defer s.conn.Close(context.Background())
	err := s.m.Down()
	s.Nil(err)
}

func TestCostUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CostUsecaseTestSuite))
}

func (s *CostUsecaseTestSuite) TestCreateCost() {
	s.Run("should create cost", func() {
		type testCase struct {
			label        string
			projectID    string
			costType     domain.CostType
			description  string
			comment      string
			amount       float64
			currency     domain.Currency
			installments []usecase.Installment
		}

		testCases := []testCase{
			{
				label:       "one time cost",
				projectID:   s.project.ProjectID,
				costType:    domain.OneTimeCost,
				description: "Mão de obra do PMO",
				comment:     "estimativa do Ferraz",
				amount:      100.0,
				currency:    domain.EUR,
				installments: []usecase.Installment{
					{Year: 2020, Month: 1, Amount: 60.},
					{Year: 2020, Month: 8, Amount: 40.},
				},
			},
			{
				label:       "running cost",
				projectID:   s.project.ProjectID,
				costType:    domain.RunningCost,
				description: "Mão de obra do PMO",
				comment:     "estimativa do Ferraz",
				amount:      100.0,
				currency:    domain.EUR,
				installments: []usecase.Installment{
					{Year: 2020, Month: 1, Amount: 60.},
					{Year: 2020, Month: 8, Amount: 40.},
				},
			},
		}

		for _, tc := range testCases {
			s.Run(fmt.Sprintf("with %s", tc.label), func() {
				ctx := context.Background()
				txm := db.NewTransactionManager(ctx, s.conn)
				txm.Register("CostRepo", func(q *db.Queries) any {
					return repository.NewCostRepositoryPostgres(q)
				})
				input := usecase.CreateCostInputDTO{
					ProjectID:    tc.projectID,
					CostType:     string(tc.costType),
					Description:  tc.description,
					Comment:      tc.comment,
					Amount:       tc.amount,
					Currency:     string(tc.currency),
					Installments: tc.installments,
				}

				uc := usecase.NewCreateCostUseCase(txm)
				cost, err := uc.Execute(ctx, input)
				s.Nil(err)

				s.Equal(tc.projectID, cost.ProjectID)
				s.Equal(string(tc.costType), cost.CostType)
				s.Equal(tc.description, cost.Description)
				s.Equal(tc.comment, cost.Comment)
				s.Equal(tc.amount, cost.Amount)
				s.Equal(string(tc.currency), cost.Currency)
				s.Equal(len(tc.installments), len(cost.Installments))
				for i := range tc.installments {
					s.Equal(tc.installments[i].Year, cost.Installments[i].Year)
					s.Equal(tc.installments[i].Month, cost.Installments[i].Month)
					s.Equal(tc.installments[i].Amount, cost.Installments[i].Amount)
				}
			})
		}
	})
}
