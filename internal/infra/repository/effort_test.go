package repository_test

import (
	"context"
	"testing"
	"time"

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

type EffortRepositoryTestSuite struct {
	suite.Suite
	dbpool *pgxpool.Pool
	m      *migrate.Migrate
	repo   domain.EstimationRepository
}

func (s *EffortRepositoryTestSuite) SetupSuite() {
	s.dbpool, s.m = testutils.DBSetup()
	s.repo = repository.NewEstimationRepositoryPostgres(s.dbpool)
}

func (s *EffortRepositoryTestSuite) TearDownSuite() {
	defer s.dbpool.Close()
	err := s.m.Down()
	s.Nil(err)
}

func (s *EffortRepositoryTestSuite) SetupSubTest() {
	err := testutils.TruncateTables(s.dbpool)
	if err != nil {
		s.T().Fatal(err)
	}
}

func TestIntegrationEffortRepository(t *testing.T) {
	suite.Run(t, new(EffortRepositoryTestSuite))
}

func (s *EffortRepositoryTestSuite) TestEffort() {
	s.Run("should create a new effort", func() {
		ctx := context.Background()

		manager := s.arrangeManager(ctx)
		baseline := s.arrangeBaseline(ctx, manager.UserID, manager.UserID)
		competence := s.arrangeCompetence(ctx)

		txm := db.NewTransactionManager(s.dbpool)
		txm.Register("EstimationRepository", func(q *db.Queries) any {
			return repository.NewEstimationRepositoryTxmPostgres(q)
		})
		err := txm.Do(ctx, func(ctx context.Context, tx db.TransactionInterface) error {
			repo, err := db.GetAs[domain.EstimationRepository](tx, "EstimationRepository")
			if err != nil {
				return err
			}

			effort := testutils.NewEffortFakeBuilder().
				WithBaselineID(baseline.BaselineID).
				WithCompetenceID(competence.CompetenceID).
				Build()

			err = repo.CreateEffort(ctx, effort)
			if err != nil {
				return err
			}

			createdEffort, err := repo.GetEffort(ctx, effort.EffortID)
			if err != nil {
				return err
			}
			s.Equal(effort.EffortID, createdEffort.EffortID)
			s.Equal(effort.BaselineID, createdEffort.BaselineID)
			s.Equal(effort.CompetenceID, createdEffort.CompetenceID)
			s.Equal(effort.Comment, createdEffort.Comment)
			s.Equal(effort.Hours, createdEffort.Hours)
			s.Equal(effort.EffortAllocations, createdEffort.EffortAllocations)
			s.True(createdEffort.CreatedAt.After(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)))
			s.True(createdEffort.CreatedAt.Before(time.Now().UTC()))
			s.True(createdEffort.UpdatedAt.IsZero())
			return err
		})
		s.Nil(err)
	})

	s.Run("should update an effort", func() {
		ctx := context.Background()

		manager := s.arrangeManager(ctx)
		baseline := s.arrangeBaseline(ctx, manager.UserID, manager.UserID)
		competence := s.arrangeCompetence(ctx)
		effort := s.arrangeEffort(ctx, baseline.BaselineID, competence.CompetenceID)

		txm := db.NewTransactionManager(s.dbpool)
		txm.Register("EstimationRepository", func(q *db.Queries) any {
			return repository.NewEstimationRepositoryTxmPostgres(q)
		})
		err := txm.Do(ctx, func(ctx context.Context, tx db.TransactionInterface) error {
			repo, err := db.GetAs[domain.EstimationRepository](tx, "EstimationRepository")
			if err != nil {
				return err
			}

			comment := "new comment"
			effort.ChangeComment(&comment)
			err = effort.Validate()
			if err != nil {
				s.FailNowf("invalid effort: %s", err.Error())
			}

			err = repo.UpdateEffort(ctx, effort)
			if err != nil {
				return err
			}

			updatedEffort, err := repo.GetEffort(ctx, effort.EffortID)
			if err != nil {
				s.FailNowf("invalid effort: %s", err.Error())
			}
			s.Equal(effort.EffortID, updatedEffort.EffortID)
			s.Equal(effort.BaselineID, updatedEffort.BaselineID)
			s.Equal(effort.CompetenceID, updatedEffort.CompetenceID)
			s.Equal(effort.Comment, updatedEffort.Comment)
			s.Equal(effort.Hours, updatedEffort.Hours)
			s.Equal(effort.EffortAllocations, updatedEffort.EffortAllocations)
			s.True(updatedEffort.CreatedAt.After(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)))
			s.True(updatedEffort.CreatedAt.Before(time.Now().UTC()))
			s.True(updatedEffort.UpdatedAt.After(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)))
			s.True(updatedEffort.UpdatedAt.Before(time.Now().UTC()))
			return err
		})
		s.Nil(err)
	})

	s.Run("should delete an effort", func() {
		ctx := context.Background()

		manager := s.arrangeManager(ctx)
		baseline := s.arrangeBaseline(ctx, manager.UserID, manager.UserID)
		competence := s.arrangeCompetence(ctx)
		effort := s.arrangeEffort(ctx, baseline.BaselineID, competence.CompetenceID)

		txm := db.NewTransactionManager(s.dbpool)
		txm.Register("EstimationRepository", func(q *db.Queries) any {
			return repository.NewEstimationRepositoryTxmPostgres(q)
		})
		err := txm.Do(ctx, func(ctx context.Context, tx db.TransactionInterface) error {
			repo, err := db.GetAs[domain.EstimationRepository](tx, "EstimationRepository")
			if err != nil {
				return err
			}

			err = repo.DeleteEffort(ctx, effort.EffortID)
			if err != nil {
				return err
			}

			_, err = repo.GetEffort(ctx, effort.EffortID)
			if err == nil {
				s.FailNowf("effort not deleted: %s", effort.EffortID)
			}
			return nil
		})
		s.Nil(err)

		row := s.dbpool.QueryRow(ctx, "SELECT COUNT(*) FROM effort_allocations")

		var count int64
		err = row.Scan(&count)
		if err != nil {
			s.FailNow("scan failed: %s", err.Error())
		}
		s.Equal(int64(0), count)
	})

}

func (s *EffortRepositoryTestSuite) arrangeManager(ctx context.Context) *service.CreateUserOutputDTO {
	service := service.NewEstimationService(s.dbpool)
	userParams := testutils.NewUserFakeBuilder().WithManager().Build()
	createdUser, err := service.CreateUser(ctx, userParams)
	if err != nil {
		s.T().Fatal(err)
	}
	return createdUser
}

func (s *EffortRepositoryTestSuite) arrangeBaseline(ctx context.Context, managerID string, estimatorID string) *domain.Baseline {
	baseline := testutils.NewBaselineFakeBuilder().
		WithManagerID(managerID).
		WithEstimatorID(estimatorID).
		Build()

	err := s.repo.CreateBaseline(ctx, baseline)
	if err != nil {
		s.T().Fatal(err)
	}

	return baseline
}

func (s *EffortRepositoryTestSuite) arrangeCompetence(ctx context.Context) *domain.Competence {
	competence := testutils.NewCompetenceFakeBuilder().Build()
	err := s.repo.CreateCompetence(ctx, competence)
	if err != nil {
		s.T().Fatal(err)
	}

	return competence
}

func (s *EffortRepositoryTestSuite) arrangeEffort(ctx context.Context, baselineID string, competenceID string) *domain.Effort {
	effort := testutils.NewEffortFakeBuilder().
		WithBaselineID(baselineID).
		WithCompetenceID(competenceID).
		Build()
	err := s.repo.CreateEffort(ctx, effort)
	if err != nil {
		s.T().Fatal(err)
	}

	return effort
}
