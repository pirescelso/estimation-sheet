package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New(validator.WithRequiredStructEnabled())
	transl   ut.Translator
)

func init() {
	en := en.New()
	unt := ut.New(en, en)
	transl, _ = unt.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(Validate, transl)
}

type PayloadValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func ValidatePayload(input any) []PayloadValidationError {
	err := Validate.Struct(input)
	if err == nil {
		return nil
	}

	reflected := reflect.ValueOf(input)

	fieldErrors := err.(validator.ValidationErrors)
	validationErrors := make([]PayloadValidationError, len(fieldErrors))
	for i, e := range fieldErrors {
		v := PayloadValidationError{}
		field, _ := reflected.Type().FieldByName(e.StructField())

		v.Error = e.Translate(transl)
		v.Field = field.Tag.Get("json")
		if v.Field == "" {
			v.Field = strings.ToLower(e.StructField())
		}

		validationErrors[i] = v
	}

	return validationErrors
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteValidationError(w http.ResponseWriter, status int, errors []PayloadValidationError) {
	WriteJSON(w, status, map[string][]PayloadValidationError{"invalid_payload": errors})
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
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
