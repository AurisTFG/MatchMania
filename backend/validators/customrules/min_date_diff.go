package customrules

import (
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

// Usage: validate:"minDateDiff=30" (minimum date difference is 30 days)
func MinDateDiffValidator(fl validator.FieldLevel) bool {
	endDate, ok1 := fl.Field().Interface().(time.Time)
	startField := fl.Parent().FieldByName("StartDate")
	startDate, ok2 := startField.Interface().(time.Time)
	if !ok1 || !ok2 {
		return false
	}

	minDays := 30
	if param := fl.Param(); param != "" {
		if days, err := strconv.Atoi(param); err == nil {
			minDays = days
		}
	}

	diff := endDate.Sub(startDate)

	return diff >= time.Duration(minDays)*24*time.Hour
}
