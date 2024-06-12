package repository_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/celsopires1999/estimation/internal/infra/repository"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/suite"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	MIGRATION_PATH = "file://../../../sql/migrations"
)

type CostRepositoryTestSuite struct {
	suite.Suite
	conn    *pgx.Conn
	m       *migrate.Migrate
	project *domain.Project
}

func (s *CostRepositoryTestSuite) SetupSuite() {
	dbURL := "postgres://postgres:postgres@db:5432/postgres?sslmode=disable"
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}

	m, err := migrate.New(MIGRATION_PATH, dbURL)
	s.Nil(err)
	s.NotNil(m)
	// sql := `
	// 	DO $$ DECLARE
	// 		r RECORD;
	// 	BEGIN
	// 		FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = current_schema()) LOOP
	// 			EXECUTE 'TRUNCATE TABLE ' || quote_ident(r.tablename) || ' CASCADE;';
	// 		END LOOP;
	// 	END;
	// 	$$;`

	// _, err = conn.Exec(context.Background(), sql)
	// s.Nil(err)

	err = m.Up()
	s.Nil(err)

	s.conn = conn
	s.m = m

	projectRepo := repository.NewProjectRepositoryPostgres(db.New(conn))
	s.project = domain.NewProject(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC))
	err = projectRepo.CreateProject(context.Background(), s.project)
	s.Nil(err)
}

func (s *CostRepositoryTestSuite) TearDownSuite() {
	defer s.conn.Close(context.Background())
	err := s.m.Down()
	s.Nil(err)
}

func TestCostRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CostRepositoryTestSuite))
}

func (s *CostRepositoryTestSuite) TestCreateCost() {
	s.Run("should create cost", func() {
		type testCase struct {
			label        string
			projectID    string
			costType     domain.CostType
			description  string
			comment      string
			amount       float64
			currency     domain.Currency
			installments []domain.NewInstallmentProps
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
				installments: []domain.NewInstallmentProps{
					{Year: 2020, Month: time.January, Amount: 60.},
					{Year: 2020, Month: time.August, Amount: 40.},
				},
			},
			{
				label:       "running cost",
				projectID:   s.project.ProjectID,
				costType:    domain.RunningCost,
				description: "aluguel de espaço de armazenamento",
				comment:     "preço mensal",
				amount:      1000.30,
				currency:    domain.USD,
				installments: []domain.NewInstallmentProps{
					{Year: 2020, Month: time.January, Amount: 600.30},
					{Year: 2020, Month: time.August, Amount: 400.},
				},
			},
		}

		for _, tc := range testCases {
			s.Run(fmt.Sprintf("with %s", tc.label), func() {
				ctx := context.Background()
				txm := db.NewTransactionManager(ctx, s.conn)
				txm.Register("costRepo", func(q *db.Queries) any {
					return repository.NewCostRepositoryPostgres(q)
				})
				err := txm.Do(ctx, func() error {
					getRepo, err := txm.GetRepository(ctx, "costRepo")
					if err != nil {
						return err
					}
					costRepo := getRepo.(*repository.CostRepositoryPostgres)

					props := domain.NewCostProps{
						ProjectID:    tc.projectID,
						CostType:     tc.costType,
						Description:  tc.description,
						Comment:      tc.comment,
						Amount:       tc.amount,
						Currency:     tc.currency,
						Installments: tc.installments,
					}
					cost := domain.NewCost(props)
					err = cost.Validate()
					if err != nil {
						return err
					}

					err = costRepo.CreateCost(ctx, cost)
					if err != nil {
						return err
					}

					model, err := costRepo.GetCost(ctx, cost.CostID)
					if err != nil {
						return err
					}
					s.Equal(cost.CostID, model.CostID)
					s.Equal(cost.ProjectID, model.ProjectID)
					s.Equal(cost.CostType, model.CostType)
					s.Equal(cost.Description, model.Description)
					s.Equal(cost.Comment, model.Comment)
					s.Equal(cost.Amount, model.Amount)
					s.Equal(cost.Currency, model.Currency)
					return err
				})
				s.Nil(err)

			})
		}
	})
}
