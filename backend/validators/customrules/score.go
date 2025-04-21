package customrules

import (
	"strconv"

	"github.com/go-playground/validator/v10"
)

// Usage: validate:"score" (score must be between 0 and 100)
func ScoreValidator(fl validator.FieldLevel) bool {
	score := fl.Field().Interface().(string)

	scoreInt, err := strconv.Atoi(score)
	if err != nil {
		return false
	}

	return scoreInt >= 0 && scoreInt <= 100
}
