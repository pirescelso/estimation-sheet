package http

import (
	"encoding/json"
	"net/http"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/usecase"
)

type effortsHandler struct {
	createEffortUseCase *usecase.CreateEffortUseCase
	updateEffortUseCase *usecase.UpdateEffortUseCase
	deleteEffortUseCase *usecase.DeleteEffortUseCase
}

func newEffortsHandler(create *usecase.CreateEffortUseCase, update *usecase.UpdateEffortUseCase, delete *usecase.DeleteEffortUseCase) *effortsHandler {
	return &effortsHandler{create, update, delete}
}

func (h *effortsHandler) createEffort(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateEffortInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	input.BaselineID = r.PathValue("baselineID")

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.createEffortUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, output)
}

func (h *effortsHandler) updateEffort(w http.ResponseWriter, r *http.Request) {
	var input usecase.UpdateEffortInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	input.EffortID = r.PathValue("effortID")
	input.BaselineID = r.PathValue("baselineID")

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.updateEffortUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, output)
}

func (h *effortsHandler) deleteEffort(w http.ResponseWriter, r *http.Request) {
	input := usecase.DeleteEffortInputDTO{
		EffortID:   r.PathValue("effortID"),
		BaselineID: r.PathValue("baselineID"),
	}

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.deleteEffortUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusNoContent, output)
}
