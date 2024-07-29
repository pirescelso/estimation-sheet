package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/celsopires1999/estimation/internal/common"
)

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func writeValidationError(w http.ResponseWriter, errors []common.PayloadValidationError) {
	m := struct {
		StatusCode int    `json:"status_code"`
		Error      string `json:"error"`
		Message    any    `json:"message"`
	}{
		StatusCode: http.StatusUnprocessableEntity,
		Error:      "Unprocessable Entity",
		Message:    map[string][]common.PayloadValidationError{"invalid_payload": errors},
	}

	writeJSON(w, http.StatusUnprocessableEntity, m)
}

func writeDomainError(w http.ResponseWriter, err error) {
	var errNotFound *common.NotFoundError
	if errors.As(err, &errNotFound) {
		writeNotFound(w, errNotFound.Error())
		return
	}

	var errConflict *common.ConflictError
	if errors.As(err, &errConflict) {
		writeConflict(w, errConflict.Error())
		return
	}
	writeError(w, http.StatusInternalServerError, err)
}

func writeNotFound(w http.ResponseWriter, msg string) {
	m := struct {
		StatusCode int    `json:"status_code"`
		Error      string `json:"error"`
		Message    string `json:"message"`
	}{
		StatusCode: http.StatusNotFound,
		Error:      "Not Found",
		Message:    msg,
	}
	writeJSON(w, http.StatusNotFound, m)
}

func writeConflict(w http.ResponseWriter, msg string) {
	m := struct {
		StatusCode int    `json:"status_code"`
		Error      string `json:"error"`
		Message    string `json:"message"`
	}{
		StatusCode: http.StatusConflict,
		Error:      "Conflict",
		Message:    msg,
	}
	writeJSON(w, http.StatusConflict, m)
}

func writeError(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, map[string]string{"error": err.Error()})
}

func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(v)
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}
