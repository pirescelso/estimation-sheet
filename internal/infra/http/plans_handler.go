package http

import (
	"encoding/json"
	"net/http"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/service"
	"github.com/celsopires1999/estimation/internal/usecase"
)

type plansHandler struct {
	createPlanUseCase *usecase.CreatePlanUseCase
	getPlanUseCase    *usecase.GetPlanUseCase
	updatePlanUseCase *usecase.UpdatePlanUseCase
	deletePlanUseCase *usecase.DeletePlanUseCase
	service           *service.EstimationService
}

func newPlansHandler(
	createPlanUseCase *usecase.CreatePlanUseCase,
	getPlanUseCase *usecase.GetPlanUseCase,
	updatePlanUseCase *usecase.UpdatePlanUseCase,
	deletePlanUseCase *usecase.DeletePlanUseCase,
	service *service.EstimationService,
) *plansHandler {
	return &plansHandler{createPlanUseCase, getPlanUseCase, updatePlanUseCase, deletePlanUseCase, service}
}

func (h *plansHandler) createPlan(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreatePlanInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.createPlanUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, output)
}

func (h *plansHandler) updatePlan(w http.ResponseWriter, r *http.Request) {
	var input usecase.UpdatePlanInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	input.PlanID = r.PathValue("planID")

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.updatePlanUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, output)
}

func (h *plansHandler) getPlan(w http.ResponseWriter, r *http.Request) {

	input := usecase.GetPlanInputDTO{
		PlanID: r.PathValue("planID"),
	}

	output, err := h.getPlanUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, output)
}

func (h *plansHandler) deletePlan(w http.ResponseWriter, r *http.Request) {
	input := usecase.DeletePlanInputDTO{
		PlanID: r.PathValue("planID"),
	}

	output, err := h.deletePlanUseCase.Execute(r.Context(), input.PlanID)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusNoContent, output)
}

func (h *plansHandler) listPlans(w http.ResponseWriter, r *http.Request) {
	input := service.ListPlansInputDTO{}
	output, err := h.service.ListPlans(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, output)
}
