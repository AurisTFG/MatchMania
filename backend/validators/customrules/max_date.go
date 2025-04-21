package customrules

import (
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

// Usage: validate:"maxDate=365" (max date is 365 days from now)
func MaxDateValidator(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}

	offsetDays, err := strconv.Atoi(fl.Param())
	if err != nil {
		offsetDays = 0
	}

	maxDate := time.Now().AddDate(0, 0, offsetDays)
	return date.Before(maxDate)
}
