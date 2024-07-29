package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/celsopires1999/estimation/internal/infra/repository"
	"github.com/celsopires1999/estimation/internal/service"
	"github.com/celsopires1999/estimation/internal/testutils"
	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type BaselineRepositoryTestSuite struct {
	suite.Suite
	dbpool      *pgxpool.Pool
	m           *migrate.Migrate
	managerID   string
	estimatorID string
}

func (s *BaselineRepositoryTestSuite) SetupSuite() {
	s.dbpool, s.m = testutils.DBSetup()
}

func (s *BaselineRepositoryTestSuite) TearDownSuite() {
	defer s.dbpool.Close()
	err := s.m.Down()
	s.Nil(err)
}

func (s *BaselineRepositoryTestSuite) SetupSubTest() {
	ctx := context.Background()
	err := testutils.TruncateTables(s.dbpool)
	if err != nil {
		s.T().Fatal(err)
	}

	service := service.NewEstimationService(s.dbpool)
	managerParams := testutils.NewUserFakeBuilder().WithManager().Build()
	manager, err := service.CreateUser(ctx, managerParams)
	if err != nil {
		s.T().Fatal(err)
	}

	estimatorParams := testutils.NewUserFakeBuilder().WithEstimator().Build()
	estimator, err := service.CreateUser(ctx, estimatorParams)
	if err != nil {
		s.T().Fatal(err)
	}

	s.managerID = manager.UserID
	s.estimatorID = estimator.UserID
}

func TestIntegrationBaselineRepository(t *testing.T) {
	suite.Run(t, new(BaselineRepositoryTestSuite))
}

func (s *BaselineRepositoryTestSuite) TestIntegrationCreateBaseline() {
	s.Run("should create baseline with transaction manager", func() {
		ctx := context.Background()
		txm := db.NewTransactionManager(s.dbpool)
		txm.Register("EstimationRepository", func(q *db.Queries) any {
			return repository.NewEstimationRepositoryTxmPostgres(q)
		})
		err := txm.Do(ctx, func(ctx context.Context, tx db.TransactionInterface) error {
			baselineRepo, err := db.GetAs[domain.EstimationRepository](tx, "EstimationRepository")
			if err != nil {
				return err
			}
			baseline := testutils.NewBaselineFakeBuilder().
				WithManagerID(s.managerID).
				WithEstimatorID(s.estimatorID).
				Build()

			err = baselineRepo.CreateBaseline(ctx, baseline)
			if err != nil {
				return err
			}

			createdBaseline, err := baselineRepo.GetBaseline(ctx, baseline.BaselineID)
			if err != nil {
				return err
			}
			s.Equal(baseline.BaselineID, createdBaseline.BaselineID)
			s.Equal(baseline.Description, createdBaseline.Description)
			s.Equal(baseline.StartDate, createdBaseline.StartDate)
			s.Equal(baseline.ManagerID, createdBaseline.ManagerID)
			s.Equal(baseline.EstimatorID, createdBaseline.EstimatorID)

			return err
		})
		s.Nil(err)
	})

	s.Run("should create baseline without transaction manager", func() {
		baselineRepo := repository.NewEstimationRepositoryPostgres(s.dbpool)
		baseline := testutils.NewBaselineFakeBuilder().
			WithManagerID(s.managerID).
			WithEstimatorID(s.estimatorID).
			Build()
		err := baselineRepo.CreateBaseline(context.Background(), baseline)
		s.Nil(err)

		createdBaseline, err := baselineRepo.GetBaseline(context.Background(), baseline.BaselineID)
		s.Nil(err)

		s.Equal(baseline.BaselineID, createdBaseline.BaselineID)
		s.Equal(baseline.Description, createdBaseline.Description)
		s.Equal(baseline.StartDate, createdBaseline.StartDate)
		s.Equal(baseline.ManagerID, createdBaseline.ManagerID)
		s.Equal(baseline.EstimatorID, createdBaseline.EstimatorID)
	})

	s.Run("should return error when baseline is not found", func() {
		baselineRepo := repository.NewEstimationRepositoryPostgres(s.dbpool)
		_, err := baselineRepo.GetBaseline(context.Background(), uuid.New().String())
		s.NotNil(err)
		var errNotFound *common.NotFoundError
		s.True(errors.As(err, &errNotFound))
		s.EqualError(err, errNotFound.Error())
	})
}
