package validators

import (
	"MatchManiaAPI/validators/customrules"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomRules(validate *validator.Validate) {
	if err := validate.RegisterValidation("dateDiff", customrules.DateDiffValidator); err != nil {
		panic("Failed to register 'dateDiff' validation: " + err.Error())
	}
	if err := validate.RegisterValidation("minDate", customrules.MinDateValidator); err != nil {
		panic("Failed to register 'minDate' validation: " + err.Error())
	}
	if err := validate.RegisterValidation("maxDate", customrules.MaxDateValidator); err != nil {
		panic("Failed to register 'maxDate' validation: " + err.Error())
	}
	if err := validate.RegisterValidation("score", customrules.ScoreValidator); err != nil {
		panic("Failed to register 'score' validation: " + err.Error())
	}
}
