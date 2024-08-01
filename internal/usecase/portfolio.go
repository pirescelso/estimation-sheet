package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/domain"
	"github.com/celsopires1999/estimation/internal/infra/db"
)

// CreatePortfolioUseCase is responsible for creating a new portfolio in the system
type CreatePortfolioUseCase struct {
	txm db.TransactionManagerInterface
}

type CreatePortfolioInputDTO struct {
	BaselineID  string `json:"baseline_id" validate:"required,uuid4"`
	PlanID      string `json:"plan_id" validate:"required,uuid4"`
	ShiftMonths int    `json:"shift_months" validate:"gte=0,lte=36"`
}

type CreatePortfolioOutputDTO struct {
	PortfolioID string `json:"portfolio_id"`
}

func NewCreatePortfolioUseCase(
	txm db.TransactionManagerInterface,
) *CreatePortfolioUseCase {
	return &CreatePortfolioUseCase{txm}
}

func (uc *CreatePortfolioUseCase) Execute(ctx context.Context, input CreatePortfolioInputDTO) (*CreatePortfolioOutputDTO, error) {
	var output CreatePortfolioOutputDTO

	err := uc.txm.Do(ctx, func(ctx context.Context, tx db.TransactionInterface) error {
		repository, err := db.GetAs[domain.EstimationRepository](tx, "EstimationRepository")
		if err != nil {
			return err
		}

		if err := uc.validate(ctx, repository, input); err != nil {
			return err
		}

		baseline, err := repository.GetBaseline(ctx, input.BaselineID)
		if err != nil {
			var notFoundErr *common.NotFoundError
			if errors.As(err, &notFoundErr) {
				return common.NewConflictError(fmt.Errorf("baseline id %s does not exist", input.BaselineID))
			}
			return err
		}

		plan, err := repository.GetPlan(ctx, input.PlanID)
		if err != nil {
			var notFoundErr *common.NotFoundError
			if errors.As(err, &notFoundErr) {
				return common.NewConflictError(fmt.Errorf("plan id %s does not exist", input.PlanID))
			}
			return err
		}

		costs, err := repository.GetCostManyByBaselineID(ctx, input.BaselineID)
		if err != nil {
			return err
		}

		inflation := plan.GetInflation()
		exchange := plan.GetExchange()

		portfolioService := domain.NewPortfolioService(input.PlanID, baseline, costs, inflation, exchange, input.ShiftMonths)
		portfolio, budgets, err := portfolioService.GeneratePortfolio()
		if err != nil {
			return err
		}

		err = repository.CreatePortfolio(ctx, portfolio)
		if err != nil {
			return err
		}

		err = repository.CreateBudgetMany(ctx, budgets)
		if err != nil {
			return err
		}

		output.PortfolioID = portfolio.PortfolioID

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (uc *CreatePortfolioUseCase) validate(ctx context.Context, repository domain.EstimationRepository, input CreatePortfolioInputDTO) error {
	if err := repository.ValidatePlan(ctx, input.PlanID); err != nil {
		var notFoundErr *common.NotFoundError
		if errors.As(err, &notFoundErr) {
			return common.NewConflictError(fmt.Errorf("plan id %s does not exist", input.PlanID))
		}
		return err
	}

	return nil
}

// DeletePortfolioUseCase is responsible for deleting a portfolio in the system
type DeletePortfolioUseCase struct {
	txm db.TransactionManagerInterface
}

type DeletePortfolioInputDTO struct {
	PortfolioID string `json:"baseline_id" validate:"required"`
}

type DeletePortfolioOutputDTO struct{}

func NewDeletePortfolioUseCase(
	txm db.TransactionManagerInterface,
) *DeletePortfolioUseCase {
	return &DeletePortfolioUseCase{txm}
}

func (uc *DeletePortfolioUseCase) Execute(ctx context.Context, input DeletePortfolioInputDTO) (*DeletePortfolioOutputDTO, error) {
	err := uc.txm.Do(ctx, func(ctx context.Context, tx db.TransactionInterface) error {
		repository, err := db.GetAs[domain.EstimationRepository](tx, "EstimationRepository")
		if err != nil {
			return err
		}

		_, err = repository.GetPortfolio(ctx, input.PortfolioID)
		if err != nil {
			return err
		}

		err = repository.DeleteBudgetsByPortfolioID(ctx, input.PortfolioID)
		if err != nil {
			return err
		}

		err = repository.DeletePortfolio(ctx, input.PortfolioID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &DeletePortfolioOutputDTO{}, nil
}
