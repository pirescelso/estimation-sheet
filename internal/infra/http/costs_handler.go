package http

import (
	"encoding/json"
	"net/http"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/usecase"
)

type costsHandler struct {
	createCostUseCase *usecase.CreateCostUseCase
	updateCostUseCase *usecase.UpdateCostUseCase
	deleteCostUseCase *usecase.DeleteCostUseCase
}

func newCostsHandler(create *usecase.CreateCostUseCase, update *usecase.UpdateCostUseCase, delete *usecase.DeleteCostUseCase) *costsHandler {
	return &costsHandler{create, update, delete}
}

func (h *costsHandler) createCost(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateCostInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	input.BaselineID = r.PathValue("baselineID")

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.createCostUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, output)
}

func (h *costsHandler) updateCost(w http.ResponseWriter, r *http.Request) {
	var input usecase.UpdateCostInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	input.CostID = r.PathValue("costID")
	input.BaselineID = r.PathValue("baselineID")

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.updateCostUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, output)
}

func (h *costsHandler) deleteCost(w http.ResponseWriter, r *http.Request) {
	input := usecase.DeleteCostInputDTO{
		CostID:     r.PathValue("costID"),
		BaselineID: r.PathValue("baselineID"),
	}

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.deleteCostUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusNoContent, output)
}
