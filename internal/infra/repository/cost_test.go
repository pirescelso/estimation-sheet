package repository_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/celsopires1999/estimation/internal/infra/repository"
	"github.com/celsopires1999/estimation/internal/service"
	"github.com/celsopires1999/estimation/internal/testutils"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type CostRepositoryTestSuite struct {
	suite.Suite
	dbpool   *pgxpool.Pool
	m        *migrate.Migrate
	baseline *domain.Baseline
}

func (s *CostRepositoryTestSuite) SetupSuite() {
	s.dbpool, s.m = testutils.DBSetup()
}

func (s *CostRepositoryTestSuite) TearDownSuite() {
	defer s.dbpool.Close()
	err := s.m.Down()
	s.Nil(err)
}

func (s *CostRepositoryTestSuite) SetupSubTest() {
	ctx := context.Background()
	err := testutils.TruncateTables(s.dbpool)
	if err != nil {
		s.T().Fatal(err)
	}

	baselineRepo := repository.NewEstimationRepositoryPostgres(s.dbpool)
	service := service.NewEstimationService(s.dbpool)
	userParams := testutils.NewUserFakeBuilder().WithManager().Build()
	createdUser, err := service.CreateUser(ctx, userParams)
	if err != nil {
		s.T().Fatal(err)
	}

	s.baseline = testutils.NewBaselineFakeBuilder().
		WithManagerID(createdUser.UserID).
		WithEstimatorID(createdUser.UserID).
		Build()

	err = baselineRepo.CreateBaseline(ctx, s.baseline)
	if err != nil {
		s.T().Fatal(err)
	}
}

func TestIntegrationCostRepository(t *testing.T) {
	suite.Run(t, new(CostRepositoryTestSuite))
}

func (s *CostRepositoryTestSuite) TestIntegrationCreateCost() {
	type testCase struct {
		label           string
		costType        domain.CostType
		description     string
		comment         string
		amount          float64
		currency        domain.Currency
		costAllocations []domain.CostAllocationProps
	}
	testCases := []testCase{
		{
			label:       "should create cost one time cost",
			costType:    domain.OneTimeCost,
			description: "Mão de obra do PMO",
			comment:     "estimativa do Ferraz",
			amount:      100.0,
			currency:    domain.EUR,
			costAllocations: []domain.CostAllocationProps{
				{Year: 2020, Month: time.January, Amount: 60.},
				{Year: 2020, Month: time.August, Amount: 40.},
			},
		},
		{
			label:       "should create cost running cost",
			costType:    domain.RunningCost,
			description: "aluguel de espaço de armazenamento",
			comment:     "preço mensal",
			amount:      1000.30,
			currency:    domain.USD,
			costAllocations: []domain.CostAllocationProps{
				{Year: 2020, Month: time.January, Amount: 600.30},
				{Year: 2020, Month: time.August, Amount: 400.},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("with %s", tc.label), func() {
			ctx := context.Background()
			txm := db.NewTransactionManager(s.dbpool)
			txm.Register("EstimationRepository", func(q *db.Queries) any {
				return repository.NewEstimationRepositoryTxmPostgres(q)
			})
			err := txm.Do(ctx, func(ctx context.Context, tx db.TransactionInterface) error {
				costRepo, err := db.GetAs[domain.EstimationRepository](tx, "EstimationRepository")
				if err != nil {
					return err
				}

				props := domain.NewCostProps{
					BaselineID:      s.baseline.BaselineID,
					CostType:        tc.costType,
					Description:     tc.description,
					Comment:         tc.comment,
					Amount:          tc.amount,
					Currency:        tc.currency,
					CostAllocations: tc.costAllocations,
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
				s.Equal(cost.BaselineID, model.BaselineID)
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
}

func (s *CostRepositoryTestSuite) TestIntegrationCreateCostError() {
	type testCase struct {
		label           string
		costType        domain.CostType
		description     string
		comment         string
		amount          float64
		currency        domain.Currency
		costAllocations []domain.CostAllocationProps
	}
	testCases := []testCase{
		{
			label:       "should not allow duplicated one time cost",
			costType:    domain.OneTimeCost,
			description: "Mão de obra do PMO",
			comment:     "estimativa do Ferraz",
			amount:      100.0,
			currency:    domain.EUR,
			costAllocations: []domain.CostAllocationProps{
				{Year: 2020, Month: time.January, Amount: 60.},
				{Year: 2020, Month: time.August, Amount: 40.},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("with %s", tc.label), func() {
			ctx := context.Background()
			txm := db.NewTransactionManager(s.dbpool)
			txm.Register("EstimationRepository", func(q *db.Queries) any {
				return repository.NewEstimationRepositoryTxmPostgres(q)
			})
			err := txm.Do(ctx, func(ctx context.Context, tx db.TransactionInterface) error {
				costRepo, err := db.GetAs[domain.EstimationRepository](tx, "EstimationRepository")
				if err != nil {
					return err
				}

				props := domain.NewCostProps{
					BaselineID:      s.baseline.BaselineID,
					CostType:        tc.costType,
					Description:     tc.description,
					Comment:         tc.comment,
					Amount:          tc.amount,
					Currency:        tc.currency,
					CostAllocations: tc.costAllocations,
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

				expectedErr := common.NewConflictError(fmt.Errorf("cost type: '%s' description: '%s' already exists in the baseline id: '%s'",
					string(cost.CostType), cost.Description, cost.BaselineID))

				err = costRepo.CreateCost(ctx, cost)
				s.EqualError(err, expectedErr.Error())
				return err
			})
			s.NotNil(err)
		})
	}
}

func (s *CostRepositoryTestSuite) TestIntegrationCreateManyCostError() {
	s.Run("should not allow duplicated on creating many costs", func() {
		ctx := context.Background()
		txm := db.NewTransactionManager(s.dbpool)
		txm.Register("EstimationRepository", func(q *db.Queries) any {
			return repository.NewEstimationRepositoryTxmPostgres(q)
		})

		cost := testutils.NewCostFakeBuilder().Build()

		costs := []*domain.Cost{cost, cost}

		err := txm.Do(ctx, func(ctx context.Context, tx db.TransactionInterface) error {
			costRepo, err := db.GetAs[domain.EstimationRepository](tx, "EstimationRepository")
			if err != nil {
				return err
			}
			expectedErr := common.NewConflictError(fmt.Errorf("duplicated cost on creating many costs: ERROR: duplicate key value violates unique constraint \"costs_pkey\" (SQLSTATE 23505)"))

			err = costRepo.CreateCostMany(ctx, costs)
			s.EqualError(err, expectedErr.Error())
			return err
		})
		s.NotNil(err)
	})
}

func (s *CostRepositoryTestSuite) TestIntegrationUpdateCostError() {
	s.Run("should not allow duplicated on updating cost", func() {
		ctx := context.Background()
		txm := db.NewTransactionManager(s.dbpool)
		txm.Register("EstimationRepository", func(q *db.Queries) any {
			return repository.NewEstimationRepositoryTxmPostgres(q)
		})

		costs := []*domain.Cost{
			testutils.NewCostFakeBuilder().WithBaselineID(s.baseline.BaselineID).WithCostType("investment").WithDescription("test cost 1").Build(),
			testutils.NewCostFakeBuilder().WithBaselineID(s.baseline.BaselineID).WithCostType("investment").WithDescription("test cost 2").Build()}

		err := txm.Do(ctx, func(ctx context.Context, tx db.TransactionInterface) error {
			costRepo, err := db.GetAs[domain.EstimationRepository](tx, "EstimationRepository")
			if err != nil {
				return err
			}

			err = costRepo.CreateCostMany(ctx, costs)
			if err != nil {
				s.T().Fatal(err)
			}
			description := "test cost 1"
			costs[1].ChangeDescription(&description)

			err = costRepo.UpdateCost(ctx, costs[1])

			expectedErr := common.NewConflictError(fmt.Errorf("cost type: '%s' description: '%s' already exists in the baseline id: '%s'",
				string(costs[1].CostType), costs[1].Description, costs[1].BaselineID))

			s.EqualError(err, expectedErr.Error())
			return err
		})
		s.NotNil(err)
	})
}
