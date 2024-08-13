package usecase_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/celsopires1999/estimation/internal/infra/repository"
	"github.com/celsopires1999/estimation/internal/testutils"
	"github.com/celsopires1999/estimation/internal/usecase"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type CreateCostUsecaseTestSuite struct {
	suite.Suite
	dbpool   *pgxpool.Pool
	m        *migrate.Migrate
	baseline *domain.Baseline
}

func (s *CreateCostUsecaseTestSuite) SetupSuite() {
	s.dbpool, s.m = testutils.DBSetup()
}

func (s *CreateCostUsecaseTestSuite) TearDownSuite() {
	defer s.dbpool.Close()
	err := s.m.Down()
	s.Nil(err)
}

func (s *CreateCostUsecaseTestSuite) SetupSubTest() {
	ctx := context.Background()
	err := testutils.TruncateTables(s.dbpool)
	if err != nil {
		s.T().Fatal(err)
	}

	repository := repository.NewEstimationRepositoryPostgres(s.dbpool)

	user := testutils.NewUserFakeBuilder().WithManager().Build()
	err = repository.CreateUser(ctx, user)
	if err != nil {
		s.T().Fatal(err)
	}

	s.baseline = testutils.NewBaselineFakeBuilder().
		WithStartDate(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)).
		WithManagerID(user.UserID).
		WithEstimatorID(user.UserID).
		Build()

	err = repository.CreateBaseline(ctx, s.baseline)
	if err != nil {
		s.T().Fatal(err)
	}
}

func TestIntegrationCreateCostUseCase(t *testing.T) {
	suite.Run(t, new(CreateCostUsecaseTestSuite))
}

func (s *CreateCostUsecaseTestSuite) TestIntegrationCreateCost() {
	type testCase struct {
		label           string
		costType        domain.CostType
		description     string
		comment         string
		amount          float64
		currency        domain.Currency
		tax             float64
		applyInflation  bool
		costAllocations []usecase.CostAllocationInput
	}

	testCases := []testCase{
		{
			label:          "should create cost with one time cost",
			costType:       domain.OneTimeCost,
			description:    "Mão de obra do PMO",
			comment:        "estimativa do Ferraz",
			amount:         100.0,
			currency:       domain.EUR,
			tax:            0.0,
			applyInflation: false,
			costAllocations: []usecase.CostAllocationInput{
				{Year: 2020, Month: 1, Amount: 60.},
				{Year: 2020, Month: 8, Amount: 40.},
			},
		},
		{
			label:          "should create cost with running cost",
			costType:       domain.RunningCost,
			description:    "Mão de obra do PMO",
			comment:        "estimativa do Ferraz",
			amount:         100.0,
			currency:       domain.EUR,
			tax:            10.0,
			applyInflation: false,
			costAllocations: []usecase.CostAllocationInput{
				{Year: 2020, Month: 1, Amount: 60.},
				{Year: 2020, Month: 8, Amount: 40.},
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

			input := usecase.CreateCostInputDTO{
				BaselineID:      s.baseline.BaselineID,
				CostType:        tc.costType.String(),
				Description:     tc.description,
				Comment:         tc.comment,
				Amount:          tc.amount,
				Currency:        tc.currency.String(),
				Tax:             tc.tax,
				ApplyInflation:  tc.applyInflation,
				CostAllocations: tc.costAllocations,
			}

			uc := usecase.NewCreateCostUseCase(txm)
			cost, err := uc.Execute(ctx, input)
			s.Nil(err)

			s.Equal(s.baseline.BaselineID, cost.BaselineID)
			s.Equal(tc.costType.String(), cost.CostType)
			s.Equal(tc.description, cost.Description)
			s.Equal(tc.comment, cost.Comment)
			s.Equal(tc.amount, cost.Amount)
			s.Equal(tc.currency.String(), cost.Currency)
			s.Equal(tc.tax, cost.Tax)
			s.Equal(tc.applyInflation, cost.ApplyInflation)
			s.Equal(len(tc.costAllocations), len(cost.CostAllocations))
			for i := range tc.costAllocations {
				s.Equal(tc.costAllocations[i].Year, cost.CostAllocations[i].Year)
				s.Equal(tc.costAllocations[i].Month, cost.CostAllocations[i].Month)
				s.Equal(tc.costAllocations[i].Amount, cost.CostAllocations[i].Amount)
			}
		})
	}

	s.Run("should fail on invalid cost allocation date", func() {
		ctx := context.Background()
		txm := db.NewTransactionManager(s.dbpool)
		txm.Register("EstimationRepository", func(q *db.Queries) any {
			return repository.NewEstimationRepositoryTxmPostgres(q)
		})

		faker := testutils.NewCostFakeBuilder()
		input := usecase.CreateCostInputDTO{
			BaselineID:     s.baseline.BaselineID,
			CostType:       faker.CostType.String(),
			Description:    faker.Description,
			Comment:        faker.Comment,
			Amount:         1_550.00,
			Currency:       faker.Currency.String(),
			Tax:            0.0,
			ApplyInflation: faker.ApplyInflation,
			CostAllocations: []usecase.CostAllocationInput{
				{Year: 2020, Month: 1, Amount: 1_000.00},
				{Year: 2019, Month: 12, Amount: 550.00},
			},
		}
		uc := usecase.NewCreateCostUseCase(txm)
		_, err := uc.Execute(ctx, input)
		s.Equal(usecase.ErrCostAllocationDateIsInvalid, err)
	})
}
