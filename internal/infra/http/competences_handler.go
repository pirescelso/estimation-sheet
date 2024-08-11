package http

import (
	"encoding/json"
	"net/http"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/service"
	"github.com/celsopires1999/estimation/internal/usecase"
)

type competencesHandler struct {
	createCompetenceUseCase *usecase.CreateCompetenceUseCase
	updateCompetenceUseCase *usecase.UpdateCompetenceUseCase
	deleteCompetenceUseCase *usecase.DeleteCompetenceUseCase
	getCompetenceUseCase    *usecase.GetCompetenceUseCase
	service                 *service.EstimationService
}

func newCompetencesHandler(
	createCompetenceUseCase *usecase.CreateCompetenceUseCase,
	updateCompetenceUseCase *usecase.UpdateCompetenceUseCase,
	deleteCompetenceUseCase *usecase.DeleteCompetenceUseCase,
	getCompetenceUseCase *usecase.GetCompetenceUseCase,
	service *service.EstimationService,
) *competencesHandler {
	return &competencesHandler{createCompetenceUseCase, updateCompetenceUseCase, deleteCompetenceUseCase, getCompetenceUseCase, service}
}

func (h *competencesHandler) createCompetence(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateCompetenceInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.createCompetenceUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, output)
}

func (h *competencesHandler) updateCompetence(w http.ResponseWriter, r *http.Request) {
	var input usecase.UpdateCompetenceInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	input.CompetenceID = r.PathValue("competenceID")

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.updateCompetenceUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, output)
}

func (h *competencesHandler) deleteCompetence(w http.ResponseWriter, r *http.Request) {
	input := usecase.DeleteCompetenceInputDTO{
		CompetenceID: r.PathValue("competenceID"),
	}

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.deleteCompetenceUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusNoContent, output)
}

func (h *competencesHandler) getCompetence(w http.ResponseWriter, r *http.Request) {
	input := usecase.GetCompetenceInputDTO{
		CompetenceID: r.PathValue("competenceID"),
	}
	output, err := h.getCompetenceUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, output)
}

func (h *competencesHandler) listCompetences(w http.ResponseWriter, r *http.Request) {
	input := service.ListCompetencesInputDTO{}
	output, err := h.service.ListCompetences(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, output)
}
