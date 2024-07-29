package common

import (
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
	Validate.RegisterValidation("twodecimals", TwoDecimalsValidator)
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

func TwoDecimalsValidator(fl validator.FieldLevel) bool {
	v := fl.Field()
	if v.Kind() != reflect.Float64 {
		return false
	}

	return IsTwoDecimals(v.Float())
}
