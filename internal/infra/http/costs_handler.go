package http

import (
	"encoding/json"
	"net/http"

	"github.com/celsopires1999/estimation/internal/usecase"
)

type CostsHandler struct {
	costsUsecase *usecase.CreateCostUseCase
}

func NewCostsHandler(costsUsecase *usecase.CreateCostUseCase) *CostsHandler {
	return &CostsHandler{
		costsUsecase: costsUsecase,
	}
}

func (h *CostsHandler) CreateCost(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateCostInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if errors := ValidatePayload(input); errors != nil {
		WriteValidationError(w, http.StatusBadRequest, errors)
		return
	}

	output, err := h.costsUsecase.Execute(r.Context(), input)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	WriteJSON(w, http.StatusCreated, output)
}
