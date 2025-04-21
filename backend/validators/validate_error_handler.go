package validators

import (
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func ValidateErrorHandler(err error) error {
	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return errors.New("unknown validation error")
	}

	errorMessage := "Unhandled validation error"
	for _, err := range validationErrors {
		tag := err.Tag()
		field := err.StructField()
		value := err.Value()
		param := err.Param()

		switch tag {
		case "required":
			errorMessage = field + " is required."
		case "min":
			errorMessage = field + " must be at least " + param + " characters long. Received: " + strconv.Itoa(
				len(value.(string)),
			)
		case "max":
			errorMessage = field + " can be up to " + param + " characters long. Received: " + strconv.Itoa(
				len(value.(string)),
			)
		case "email":
			errorMessage = field + " must be a valid email address."
		case "score":
			errorMessage = "Score must be between 0 and 100. Received: " + value.(string)
		case "minDate":
			if param[0] == '-' {
				errorMessage = field + " cannot be more than " + param[1:] + " days in the past."
			} else {
				errorMessage = field + " must be at least " + param + " days in the future."
			}
		case "maxDate":
			if param[0] == '-' {
				errorMessage = field + " cannot be more than " + param[1:] + " days in the past."
			} else {
				errorMessage = field + " must be at least " + param + " days in the future."
			}
		case "dateDiff":
			errorMessage = "End Date must be after Start Date and be " + param + " days apart."
		default:
			errorMessage = field + " failed on " + tag + " validation."
		}
	}

	return errors.New(errorMessage)
}
