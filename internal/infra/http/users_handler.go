package http

import (
	"encoding/json"
	"net/http"

	"github.com/celsopires1999/estimation/internal/common"
	"github.com/celsopires1999/estimation/internal/service"
	"github.com/celsopires1999/estimation/internal/usecase"
)

type usersHandler struct {
	createUserUseCase *usecase.CreateUserUseCase
	updateUserUseCase *usecase.UpdateUserUseCase
	getUserUseCase    *usecase.GetUserUseCase
	deleteUserUseCase *usecase.DeleteUserUseCase
	service           *service.EstimationService
}

func newUsersHandler(
	createUserUseCase *usecase.CreateUserUseCase,
	updateUserUseCase *usecase.UpdateUserUseCase,
	getUserUseCase *usecase.GetUserUseCase,
	deleteUserUseCase *usecase.DeleteUserUseCase,
	service *service.EstimationService,
) *usersHandler {
	return &usersHandler{createUserUseCase, updateUserUseCase, getUserUseCase, deleteUserUseCase, service}
}

func (h *usersHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateUserInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.createUserUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, output)
}

func (h *usersHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	var input usecase.UpdateUserInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	input.UserID = userID

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.updateUserUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, output)
}

func (h *usersHandler) getUser(w http.ResponseWriter, r *http.Request) {
	input := usecase.GetUserInputDTO{
		UserID: r.PathValue("userID"),
	}

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.getUserUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, output)
}

func (h *usersHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	input := usecase.DeleteUserInputDTO{
		UserID: r.PathValue("userID"),
	}

	if errors := common.ValidatePayload(input); errors != nil {
		writeValidationError(w, errors)
		return
	}

	output, err := h.deleteUserUseCase.Execute(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusNoContent, output)
}

func (h *usersHandler) listUsers(w http.ResponseWriter, r *http.Request) {
	input := service.ListUsersInputDTO{}
	qs := r.URL.Query()

	input.Name = readString(qs, "name", "")
	input.Filters.SortSafelist = []string{"name", "created_at", "-name", "-created_at"}

	if err := setFilter(qs, &input.Filters); err != nil {
		writeBadRequest(w, err.Error())
	}

	output, err := h.service.ListUsers(r.Context(), input)
	if err != nil {
		writeDomainError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, output)
}
