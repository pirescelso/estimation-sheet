package usecase_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/celsopires1999/estimation/internal/infra/repository"

	"github.com/celsopires1999/estimation/internal/testutils"
	"github.com/celsopires1999/estimation/internal/usecase"
	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type CreatePortfolioUseCaseTestSuite struct {
	suite.Suite
	dbpool     *pgxpool.Pool
	m          *migrate.Migrate
	txm        db.TransactionManagerInterface
	repository domain.EstimationRepository
}

func (s *CreatePortfolioUseCaseTestSuite) SetupSuite() {
	s.dbpool, s.m = testutils.DBSetup()
	s.txm = db.NewTransactionManager(s.dbpool)
	s.txm.Register("EstimationRepository", func(q *db.Queries) any {
		return repository.NewEstimationRepositoryTxmPostgres(q)
	})

	s.repository = repository.NewEstimationRepositoryPostgres(s.dbpool)
}

func (s *CreatePortfolioUseCaseTestSuite) TearDownSuite() {
	defer s.dbpool.Close()
	err := s.m.Down()
	s.Nil(err)
}

func (s *CreatePortfolioUseCaseTestSuite) SetupSubTest() {
	err := testutils.TruncateTables(s.dbpool)
	if err != nil {
		s.T().Fatal(err)
	}
}

func TestIntegrationCreatePortfolioUseCase(t *testing.T) {
	suite.Run(t, new(CreatePortfolioUseCaseTestSuite))
}

func (s *CreatePortfolioUseCaseTestSuite) TestIntegrationCreatePortfolio() {
	s.Run("should create portfolio with shift of 8 months", func() {
		ctx := context.Background()
		baseline := s.createDependenciesBaseline(ctx)
		s.createDependencies8Months(ctx, baseline)
		plan := s.createDependenciesPlan(ctx)
		input := usecase.CreatePortfolioInputDTO{
			BaselineID:  baseline.BaselineID,
			PlanID:      plan.PlanID,
			ShiftMonths: 8,
		}

		uc := usecase.NewCreatePortfolioUseCase(s.txm)
		output, err := uc.Execute(ctx, input)

		s.Nil(err)
		s.NotNil(output)

		if err != nil {
			s.T().Fatal(err)
		}
		_, err = uuid.Parse(output.PortfolioID)
		s.Nil(err)

		portfolio, err := s.repository.GetPortfolio(ctx, output.PortfolioID)
		s.Nil(err)
		s.Equal(baseline.StartDate.AddDate(0, 8, 0), portfolio.StartDate)

		budgets, err := s.repository.GetBudgetManyByPortfolioID(ctx, output.PortfolioID)
		s.Nil(err)
		s.Equal(2, len(budgets))

		for _, budget := range budgets {
			for _, allocation := range budget.BudgetAllocations {
				if allocation.AllocationDate.Month() == 9 && allocation.AllocationDate.Year() == 2022 {
					s.Equal(610.15, allocation.Amount)
				}
				if allocation.AllocationDate.Month() == 4 && allocation.AllocationDate.Year() == 2023 {
					s.Equal(279.61, allocation.Amount)
				}
				if allocation.AllocationDate.Month() == 10 && allocation.AllocationDate.Year() == 2023 {
					s.Equal(821.68, allocation.Amount)
				}
				if allocation.AllocationDate.Month() == 8 && allocation.AllocationDate.Year() == 2024 {
					s.Equal(676.84, allocation.Amount)
				}
			}
		}
	})

	s.Run("should create portfolio in BRL", func() {
		ctx := context.Background()
		baseline := s.createDependenciesBaseline(ctx)
		costs := s.createDependenciesBRL(ctx, baseline)
		plan := s.createDependenciesPlan(ctx)
		input := usecase.CreatePortfolioInputDTO{
			BaselineID:  baseline.BaselineID,
			PlanID:      plan.PlanID,
			ShiftMonths: 0,
		}

		uc := usecase.NewCreatePortfolioUseCase(s.txm)
		output, err := uc.Execute(ctx, input)

		s.Nil(err)
		s.NotNil(output)
		_, err = uuid.Parse(output.PortfolioID)
		s.Nil(err)

		portfolio, err := s.repository.GetPortfolio(ctx, output.PortfolioID)
		s.Nil(err)
		s.Equal(baseline.StartDate.AddDate(0, 0, 0), portfolio.StartDate)

		budgets, err := s.repository.GetBudgetManyByPortfolioID(ctx, output.PortfolioID)
		s.Nil(err)
		s.Equal(1, len(budgets))
		s.Equal(970.53, budgets[0].Amount)
		s.Equal(970.53, budgets[0].BudgetAllocations[0].Amount)
		s.Equal(costs[0].CostAllocations[0].AllocationDate, budgets[0].BudgetAllocations[0].AllocationDate)
	})
	s.Run("should create portfolio with running cost", func() {
		ctx := context.Background()
		baseline := s.createDependenciesBaseline(ctx)
		costs := s.createDependenciesRC(ctx, baseline)
		plan := s.createDependenciesPlan(ctx)
		input := usecase.CreatePortfolioInputDTO{
			BaselineID:  baseline.BaselineID,
			PlanID:      plan.PlanID,
			ShiftMonths: 0,
		}

		uc := usecase.NewCreatePortfolioUseCase(s.txm)
		output, err := uc.Execute(ctx, input)
		s.Nil(err)
		budgets, err := s.repository.GetBudgetManyByPortfolioID(ctx, output.PortfolioID)
		s.Nil(err)
		s.Equal(107_640.00, budgets[0].Amount)
		s.Equal(107_640.00, budgets[0].BudgetAllocations[0].Amount)
		s.Equal(costs[0].CostAllocations[0].AllocationDate, budgets[0].BudgetAllocations[0].AllocationDate)
	})

	s.Run("should not create portfolio with existing baseline and plan", func() {
		ctx := context.Background()
		baseline := s.createDependenciesBaseline(ctx)
		s.createDependencies8Months(ctx, baseline)
		plan := s.createDependenciesPlan(ctx)
		input := usecase.CreatePortfolioInputDTO{
			BaselineID:  baseline.BaselineID,
			PlanID:      plan.PlanID,
			ShiftMonths: 8,
		}

		uc := usecase.NewCreatePortfolioUseCase(s.txm)
		output, err := uc.Execute(ctx, input)

		s.Nil(err)
		s.NotNil(output)

		if err != nil {
			s.T().Fatal(err)
		}

		newBaseline := domain.NewBaseline(
			baseline.Code,
			baseline.Review+1,
			baseline.Title,
			baseline.Description,
			baseline.StartDate,
			baseline.Duration,
			baseline.ManagerID,
			baseline.EstimatorID,
		)

		err = s.repository.CreateBaseline(ctx, newBaseline)
		if err != nil {
			s.T().Fatal(err)
		}

		s.createDependencies8Months(ctx, newBaseline)

		input = usecase.CreatePortfolioInputDTO{
			BaselineID:  newBaseline.BaselineID,
			PlanID:      plan.PlanID,
			ShiftMonths: 8,
		}

		_, err = uc.Execute(ctx, input)

		expectedError := common.NewConflictError(fmt.Errorf("portfolio for plan id %s and baseline code %s already exists", plan.PlanID, newBaseline.Code))

		s.Equal(expectedError, err)
	})

}

func (s *CreatePortfolioUseCaseTestSuite) createDependenciesBaseline(ctx context.Context) *domain.Baseline {
	user := testutils.NewUserFakeBuilder().WithManager().Build()
	err := s.repository.CreateUser(ctx, user)
	if err != nil {
		s.T().Fatal(err)
	}
	baseline := testutils.NewBaselineFakeBuilder().
		WithManagerID(user.UserID).
		WithEstimatorID(user.UserID).
		WithStartDate(time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)).
		Build()

	err = s.repository.CreateBaseline(ctx, baseline)
	if err != nil {
		s.T().Fatal(err)
	}

	return baseline
}

func (s *CreatePortfolioUseCaseTestSuite) createDependencies8Months(ctx context.Context, baseline *domain.Baseline) []*domain.Cost {
	costs := make([]*domain.Cost, 0)

	costs = append(costs, testutils.NewCostFakeBuilder().
		WithBaselineID(baseline.BaselineID).
		WithAmount(880.30).
		WithCurrency("BRL").
		WithApplyInflation(true).
		WithCostAllocationProps([]domain.CostAllocationProps{
			{Year: 2022, Month: time.January, Amount: 610.15},
			{Year: 2022, Month: time.August, Amount: 270.15},
		}).
		WithTax(0.00).
		Build())

	costs = append(costs, testutils.NewCostFakeBuilder().
		WithBaselineID(baseline.BaselineID).
		WithAmount(220.20).
		WithCurrency("EUR").
		WithApplyInflation(false).
		WithCostAllocationProps([]domain.CostAllocationProps{
			{Year: 2023, Month: time.February, Amount: 120.15},
			{Year: 2023, Month: time.December, Amount: 100.05},
		}).
		WithTax(23.00).
		Build())

	err := s.repository.CreateCostMany(ctx, costs)
	if err != nil {
		s.T().Fatal(err)
	}

	return costs
}

func (s *CreatePortfolioUseCaseTestSuite) createDependenciesBRL(ctx context.Context, baseline *domain.Baseline) []*domain.Cost {
	costs := make([]*domain.Cost, 0)

	costs = append(costs, testutils.NewCostFakeBuilder().
		WithBaselineID(baseline.BaselineID).
		WithAmount(880.30).
		WithCurrency("BRL").
		WithTax(10.25).
		WithApplyInflation(true).
		WithCostAllocationProps([]domain.CostAllocationProps{
			{Year: 2022, Month: time.August, Amount: 880.30},
		}).
		Build())

	err := s.repository.CreateCostMany(ctx, costs)
	if err != nil {
		s.T().Fatal(err)
	}

	return costs
}

func (s *CreatePortfolioUseCaseTestSuite) createDependenciesPlan(ctx context.Context) *domain.Plan {
	plan := testutils.NewPlanFakeBuilder().Build()
	err := s.repository.CreatePlan(ctx, plan)
	if err != nil {
		s.T().Fatal(err)
	}

	return plan
}

func (s *CreatePortfolioUseCaseTestSuite) createDependenciesRC(ctx context.Context, baseline *domain.Baseline) []*domain.Cost {
	costs := make([]*domain.Cost, 0)

	costs = append(costs, testutils.NewCostFakeBuilder().
		WithBaselineID(baseline.BaselineID).
		WithCostType(domain.RunningCost).
		WithAmount(100_000.00).
		WithCurrency("BRL").
		WithTax(0.00).
		WithApplyInflation(true).
		WithCostAllocationProps([]domain.CostAllocationProps{
			{Year: 2024, Month: time.August, Amount: 100_000.00},
		}).
		Build())

	err := s.repository.CreateCostMany(ctx, costs)
	if err != nil {
		s.T().Fatal(err)
	}

	return costs
}
