package http

import (
	"encoding/json"
	"net/http"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/usecase"
)

type baselineHandler struct {
	createBaselineUseCase *usecase.CreateBaselineUseCase
	updateBaselineUseCase *usecase.UpdateBaselineUseCase
	deleteBaselineUseCase *usecase.DeleteBaselineUseCase
}

func newBaselinesHandler(create *usecase.CreateBaselineUseCase, update *usecase.UpdateBaselineUseCase, delete *usecase.DeleteBaselineUseCase) *baselineHandler {
	return &baselineHandler{create, update, delete}
}

func (h *baselineHandler) createBaseline(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateBaselineInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.createBaselineUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, output)
}

func (h *baselineHandler) updateBaseline(w http.ResponseWriter, r *http.Request) {
	var input usecase.UpdateBaselineInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	input.BaselineID = r.PathValue("baselineID")

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.updateBaselineUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, output)
}

func (h *baselineHandler) deleteBaseline(w http.ResponseWriter, r *http.Request) {
	input := usecase.DeleteBaselineInputDTO{
		BaselineID: r.PathValue("baselineID"),
	}

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.deleteBaselineUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusNoContent, output)
}
