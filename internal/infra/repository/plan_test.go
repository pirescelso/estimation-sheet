package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/celsopires1999/estimation/internal/infra/repository"
	"github.com/celsopires1999/estimation/internal/testutils"
	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type PlanRepositoryTestSuite struct {
	suite.Suite
	dbpool *pgxpool.Pool
	m      *migrate.Migrate
}

func (s *PlanRepositoryTestSuite) SetupSuite() {
	s.dbpool, s.m = testutils.DBSetup()
}

func (s *PlanRepositoryTestSuite) TearDownSuite() {
	defer s.dbpool.Close()
	err := s.m.Down()
	s.Nil(err)
}

func (s *PlanRepositoryTestSuite) SetupSubTest() {
	err := testutils.TruncateTables(s.dbpool)
	if err != nil {
		s.T().Fatal(err)
	}
}

func TestIntegrationPlanlineRepository(t *testing.T) {
	suite.Run(t, new(PlanRepositoryTestSuite))
}

func (s *PlanRepositoryTestSuite) TestIntegrationCreatePlan() {
	s.Run("should create plan with transaction manager", func() {
		ctx := context.Background()
		txm := db.NewTransactionManager(s.dbpool)
		txm.Register("EstimationRepository", func(q *db.Queries) any {
			return repository.NewEstimationRepositoryTxmPostgres(q)
		})
		err := txm.Do(ctx, func(ctx context.Context, tx db.TransactionInterface) error {
			repo, err := db.GetAs[domain.EstimationRepository](tx, "EstimationRepository")
			if err != nil {
				return err
			}
			plan := testutils.NewPlanFakeBuilder().Build()

			err = repo.CreatePlan(ctx, plan)
			if err != nil {
				return err
			}

			createdPlan, err := repo.GetPlan(ctx, plan.PlanID)
			if err != nil {
				return err
			}
			s.Equal(plan.PlanID, createdPlan.PlanID)
			s.Equal(plan.Code, createdPlan.Code)
			s.Equal(plan.Name, createdPlan.Name)
			s.Equal(plan.Assumptions, createdPlan.Assumptions)

			return err
		})
		s.Nil(err)
	})

	s.Run("should return error when plan is not found", func() {
		repo := repository.NewEstimationRepositoryPostgres(s.dbpool)
		_, err := repo.GetPlan(context.Background(), uuid.New().String())
		s.NotNil(err)
		var errNotFound *common.NotFoundError
		s.True(errors.As(err, &errNotFound))
		s.EqualError(err, errNotFound.Error())
	})
}
