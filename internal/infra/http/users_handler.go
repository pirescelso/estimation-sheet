package http

import (
	"encoding/json"
	"net/http"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/service"
)

type usersHandler struct {
	service *service.EstimationService
}

func newUsersHandler(service *service.EstimationService) *usersHandler {
	return &usersHandler{service}
}

func (h *usersHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var input service.CreateUserInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.service.CreateUser(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, output)
}

func (h *usersHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	var input service.UpdateUserInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	input.UserID = userID

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.service.UpdateUser(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, output)
}

func (h *usersHandler) getUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")

	output, err := h.service.GetUser(r.Context(), userID)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, output)
}

func (h *usersHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")

	if err := h.service.DeleteUser(r.Context(), userID); err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusNoContent, nil)
}

func (h *usersHandler) listUsers(w http.ResponseWriter, r *http.Request) {
	input := service.ListUsersInputDTO{}
	output, err := h.service.ListUsers(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, output)
}
