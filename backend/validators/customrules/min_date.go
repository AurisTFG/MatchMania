package customrules

import (
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

// Usage: validate:"minDate=-365" (min date is 365 days ago)
func MinDateValidator(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}

	offsetDays, err := strconv.Atoi(fl.Param())
	if err != nil {
		offsetDays = 0
	}

	minDate := time.Now().AddDate(0, 0, offsetDays)
	return date.After(minDate)
}
