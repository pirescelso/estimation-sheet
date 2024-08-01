package http

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/celsopires1999/estimation/internal/infra/repository"
	"github.com/celsopires1999/estimation/internal/service"
	"github.com/celsopires1999/estimation/internal/usecase"
)

func Handler(ctx context.Context, dbpool *pgxpool.Pool) *http.ServeMux {
	txm := db.NewTransactionManager(dbpool)
	txm.Register("EstimationRepository", func(q *db.Queries) any {
		return repository.NewEstimationRepositoryTxmPostgres(q)
	})
	repository := repository.NewEstimationRepositoryPostgres(dbpool)

	// Services (CRUDs & Queries)
	service := service.NewEstimationService(dbpool)

	// UseCases
	createPlanUseCase := usecase.NewCreatePlanUseCase(repository)
	getPlanUseCase := usecase.NewGetPlanUseCase(repository)
	updatePlanUseCase := usecase.NewUpdatePlanUseCase(repository)
	deletePlanUseCase := usecase.NewDeletePlanUseCase(repository)

	createBaselineUseCase := usecase.NewCreateBaselineUseCase(repository)
	updateBaselineUseCase := usecase.NewUpdateBaselineUseCase(repository)
	deleteBaselineUseCase := usecase.NewDeleteBaselineUseCase(repository)

	createCostUsecase := usecase.NewCreateCostUseCase(txm)
	updateCostUseCase := usecase.NewUpdateCostUseCase(txm)
	deleteCostUseCase := usecase.NewDeleteCostUseCase(txm)
	getCostsByBaselineIDUseCase := usecase.NewGetCostsByBaselineIDUseCase(repository)

	createPortfolioUseCase := usecase.NewCreatePortfolioUseCase(txm)
	deletePortfolioUseCase := usecase.NewDeletePortfolioUseCase(txm)

	// Handlers
	usersHandler := newUsersHandler(service)
	plansHandler := newPlansHandler(createPlanUseCase, getPlanUseCase, updatePlanUseCase, deletePlanUseCase, service)
	baselinesHandler := newBaselinesHandler(createBaselineUseCase, updateBaselineUseCase, deleteBaselineUseCase, getCostsByBaselineIDUseCase, service)
	costsHandler := newCostsHandler(createCostUsecase, updateCostUseCase, deleteCostUseCase)
	portfoliosHandler := newPortfoliosHandler(createPortfolioUseCase, deletePortfolioUseCase, service)

	// Routes
	r := http.NewServeMux()
	r.HandleFunc("POST /users", usersHandler.createUser)
	r.HandleFunc("PATCH /users/{userID}", usersHandler.updateUser)
	r.HandleFunc("DELETE /users/{userID}", usersHandler.deleteUser)
	r.HandleFunc("GET /users/{userID}", usersHandler.getUser)
	r.HandleFunc("GET /users", usersHandler.listUsers)

	r.HandleFunc("POST /plans", plansHandler.createPlan)
	r.HandleFunc("PATCH /plans/{planID}", plansHandler.updatePlan)
	r.HandleFunc("DELETE /plans/{planID}", plansHandler.deletePlan)
	r.HandleFunc("GET /plans/{planID}", plansHandler.getPlan)
	r.HandleFunc("GET /plans", plansHandler.listPlans)

	r.HandleFunc("POST /baselines", baselinesHandler.createBaseline)
	r.HandleFunc("PATCH /baselines/{baselineID}", baselinesHandler.updateBaseline)
	r.HandleFunc("DELETE /baselines/{baselineID}", baselinesHandler.deleteBaseline)
	r.HandleFunc("GET /baselines/{baselineID}", baselinesHandler.getBaseline)
	r.HandleFunc("GET /baselines", baselinesHandler.listBaselines)
	r.HandleFunc("GET /baselines/{baselineID}/costs", baselinesHandler.getCostsByBaselineID)

	r.HandleFunc("POST /baselines/{baselineID}/costs", costsHandler.createCost)
	r.HandleFunc("PATCH /baselines/{baselineID}/costs/{costID}", costsHandler.updateCost)
	r.HandleFunc("DELETE /baselines/{baselineID}/costs/{costID}", costsHandler.deleteCost)

	r.HandleFunc("POST /portfolios", portfoliosHandler.createPortfolio)
	r.HandleFunc("DELETE /portfolios/{portfolioID}", portfoliosHandler.deletePortfolio)
	r.HandleFunc("GET /portfolios/{portfolioID}", portfoliosHandler.getPortfolioById)
	r.HandleFunc("GET /portfolios", portfoliosHandler.listPortfolios)

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", r))
	return v1
}
