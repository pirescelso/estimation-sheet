package http

import (
	"encoding/json"
	"net/http"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/service"
	"github.com/celsopires1999/estimation/internal/usecase"
)

type portfoliosHandler struct {
	createPortfolioUseCase *usecase.CreatePortfolioUseCase
	deletePortfolioUseCase *usecase.DeletePortfolioUseCase
	service                *service.EstimationService
}

func newPortfoliosHandler(
	createPortfolioUseCase *usecase.CreatePortfolioUseCase,
	deletePortfolioUseCase *usecase.DeletePortfolioUseCase,
	service *service.EstimationService,
) *portfoliosHandler {
	return &portfoliosHandler{
		createPortfolioUseCase,
		deletePortfolioUseCase,
		service,
	}
}

func (h *portfoliosHandler) createPortfolio(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreatePortfolioInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.createPortfolioUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, output)
}

func (h *portfoliosHandler) deletePortfolio(w http.ResponseWriter, r *http.Request) {
	input := usecase.DeletePortfolioInputDTO{
		PortfolioID: r.PathValue("portfolioID")}

	output, err := h.deletePortfolioUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusNoContent, output)
}

func (h *portfoliosHandler) getPortfolioById(w http.ResponseWriter, r *http.Request) {
	input := service.GetPortfolioInputDTO{
		PortfolioID: r.PathValue("portfolioID"),
	}

	output, err := h.service.GetPortfolio(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, output)
}

func (h *portfoliosHandler) listPortfolios(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if len(query) == 0 {
		h.listAllPortfolios(w, r)
		return
	}

	h.listPortfoliosByPlanID(w, r)
}

func (h *portfoliosHandler) listPortfoliosByPlanID(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	planID := query.Get("planID")

	input := service.ListPortfoliosInputDTO{
		PlanID: planID,
	}

	output, err := h.service.ListPortfoliosByPlanID(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, output)
}

func (h *portfoliosHandler) listAllPortfolios(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.ListAllPortfolios(r.Context())
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, output)
}
