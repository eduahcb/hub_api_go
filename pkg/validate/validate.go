package validate

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidateErrors struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func camelToSnake(s string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(snake)
}

func createMessage(field, tag, param string) string {
	switch tag {
	case "email":
		return fmt.Sprintf("The field '%s' must be a valid email", field)
	case "required":
		return fmt.Sprintf("The field '%s' is required", field)
	case "min":
		return fmt.Sprintf("The field '%s' must be at least %s characters long", field, param)
	case "max":
		return fmt.Sprintf("The field '%s' must not exceed %s characters", field, param)
	case "eqfield":
		return fmt.Sprintf("The field '%s' must be equal to the field '%s'", field, param)
	default:
		return fmt.Sprintf("Validation error on field '%s'", field)
	}
}

func ValidateRequestBody(body interface{}) ([]ValidateErrors, error) {
	var errors []ValidateErrors

	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := camelToSnake(err.Field())

			errors = append(errors, ValidateErrors{
				Field:   field,
				Message: createMessage(field, err.Tag(), err.Param()),
			})
		}

		return errors, err
	}

	return errors, err
}
