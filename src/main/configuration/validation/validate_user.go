package validation

import (
	"HunCoding/src/main/configuration/rest_err"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"
	ut "github.com/go-playground/universal-translator"
	entranslation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ := unt.GetTranslator("en")
		err := entranslation.RegisterDefaultTranslations(val, transl)
		if err != nil {
			return
		}
	}
}

func ValidateUserError(validationErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validationErr, &jsonValidationError) {
		var errorCauses []rest_err.Causes

		for _, e := range validationErr.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
