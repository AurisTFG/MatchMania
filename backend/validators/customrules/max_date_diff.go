package customrules

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

// Usage: validate:"maxDateDiff=90" (maximum date difference is 90 days)
func MaxDateDiffValidator(fl validator.FieldLevel) bool {
	endDate, ok1 := fl.Field().Interface().(time.Time)
	startField := fl.Parent().FieldByName("StartDate")
	startDate, ok2 := startField.Interface().(time.Time)
	if !ok1 || !ok2 {
		return false
	}

	maxDays := 90
	if param := fl.Param(); param != "" {
		if days, err := strconv.Atoi(param); err == nil {
			maxDays = days
		}
	}

	diff := endDate.Sub(startDate)
	fmt.Println("MaxDateDiffValidator", startDate, endDate, diff, maxDays)
	return diff <= time.Duration(maxDays)*24*time.Hour
}
