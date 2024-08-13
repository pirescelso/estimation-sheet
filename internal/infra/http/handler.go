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

	// Services for Queries
	service := service.NewEstimationService(dbpool)

	// UseCases
	createUserUseCase := usecase.NewCreateUserUseCase(repository)
	getUserUseCase := usecase.NewGetUserUseCase(repository)
	updateUserUseCase := usecase.NewUpdateUserUseCase(repository)
	deleteUserUseCase := usecase.NewDeleteUserUseCase(repository)

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

	createCompetenceUseCase := usecase.NewCreateCompetenceUseCase(repository)
	updateCompetenceUseCase := usecase.NewUpdateCompetenceUseCase(repository)
	deleteCompetenceUseCase := usecase.NewDeleteCompetenceUseCase(repository)
	getCompetenceUseCase := usecase.NewGetCompetenceUseCase(repository)

	createEffortUseCase := usecase.NewCreateEffortUseCase(txm)
	updateEffortUseCase := usecase.NewUpdateEfforttUseCase(txm)
	deleteEffortUseCase := usecase.NewDeleteEffortUseCase(txm)
	getEffortsByBaselineIDUseCase := usecase.NewGetEffortsByBaselineIDUseCase(repository)

	createPortfolioUseCase := usecase.NewCreatePortfolioUseCase(txm)
	deletePortfolioUseCase := usecase.NewDeletePortfolioUseCase(txm)

	// Handlers
	usersHandler := newUsersHandler(createUserUseCase, updateUserUseCase, getUserUseCase, deleteUserUseCase, service)
	plansHandler := newPlansHandler(createPlanUseCase, getPlanUseCase, updatePlanUseCase, deletePlanUseCase, service)
	baselinesHandler := newBaselinesHandler(createBaselineUseCase, updateBaselineUseCase, deleteBaselineUseCase, getCostsByBaselineIDUseCase, getEffortsByBaselineIDUseCase, service)
	costsHandler := newCostsHandler(createCostUsecase, updateCostUseCase, deleteCostUseCase)
	competencesHandler := newCompetencesHandler(createCompetenceUseCase, updateCompetenceUseCase, deleteCompetenceUseCase, getCompetenceUseCase, service)
	effortsHandler := newEffortsHandler(createEffortUseCase, updateEffortUseCase, deleteEffortUseCase)
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

	r.HandleFunc("POST /competences", competencesHandler.createCompetence)
	r.HandleFunc("PATCH /competences/{competenceID}", competencesHandler.updateCompetence)
	r.HandleFunc("DELETE /competences/{competenceID}", competencesHandler.deleteCompetence)
	r.HandleFunc("GET /competences/{competenceID}", competencesHandler.getCompetence)
	r.HandleFunc("GET /competences", competencesHandler.listCompetences)

	r.HandleFunc("POST /baselines", baselinesHandler.createBaseline)
	r.HandleFunc("PATCH /baselines/{baselineID}", baselinesHandler.updateBaseline)
	r.HandleFunc("DELETE /baselines/{baselineID}", baselinesHandler.deleteBaseline)
	r.HandleFunc("GET /baselines/{baselineID}", baselinesHandler.getBaseline)
	r.HandleFunc("GET /baselines", baselinesHandler.listBaselines)
	r.HandleFunc("GET /baselines/{baselineID}/costs", baselinesHandler.getCostsByBaselineID)
	r.HandleFunc("GET /baselines/{baselineID}/efforts", baselinesHandler.getEffortsByBaselineID)

	r.HandleFunc("POST /baselines/{baselineID}/costs", costsHandler.createCost)
	r.HandleFunc("PATCH /baselines/{baselineID}/costs/{costID}", costsHandler.updateCost)
	r.HandleFunc("DELETE /baselines/{baselineID}/costs/{costID}", costsHandler.deleteCost)

	r.HandleFunc("POST /baselines/{baselineID}/efforts", effortsHandler.createEffort)
	r.HandleFunc("PATCH /baselines/{baselineID}/efforts/{effortID}", effortsHandler.updateEffort)
	r.HandleFunc("DELETE /baselines/{baselineID}/efforts/{effortID}", effortsHandler.deleteEffort)

	r.HandleFunc("POST /portfolios", portfoliosHandler.createPortfolio)
	r.HandleFunc("DELETE /portfolios/{portfolioID}", portfoliosHandler.deletePortfolio)
	r.HandleFunc("GET /portfolios/{portfolioID}", portfoliosHandler.getPortfolioById)
	r.HandleFunc("GET /portfolios", portfoliosHandler.listPortfolios)

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", r))
	return v1
}
