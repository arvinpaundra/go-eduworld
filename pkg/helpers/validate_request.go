package helpers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/arvinpaundra/go-eduworld/pkg/constants"
	"github.com/go-playground/validator/v10"
)

type ValidationError map[string]string

func ValidateRequest(value interface{}) ValidationError {
	// register tag json to validate instead validate field from struct
	constants.Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]

		// skip tag if want to be ignored
		if name == "-" {
			return ""
		}

		return name
	})

	err := constants.Validate.Struct(value)

	if err != nil {
		return validationErrorMessage(err)
	}

	return nil
}

func validationErrorMessage(validationError error) ValidationError {
	// store error fields and error messages
	errFields := make(map[string]string)

	// make error message for each invalid field
	for _, err := range validationError.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			errFields[err.Field()] = "this field is required"
			break
		case "email":
			errFields[err.Field()] = "invalid email format"
			break
		case "min":
			errFields[err.Field()] = fmt.Sprintf("min length %s characters", err.Param())
			break
		case "max":
			errFields[err.Field()] = fmt.Sprintf("max length %s characters", err.Param())
			break
		case "required_if":
			errFields[err.Field()] = "this field is required"
			break
		default:
			errFields[err.Field()] = err.Error()
			break
		}
	}

	return errFields
}
