package validators

import (
	"MatchManiaAPI/validators/customrules"

	"github.com/go-playground/validator/v10"
)

func MustRegisterCustomRule(validate *validator.Validate, tag string, fn validator.Func) {
	if err := validate.RegisterValidation(tag, fn); err != nil {
		panic("Failed to register custom validation rule: " + err.Error())
	}
}

func MustRegisterCustomRules(validate *validator.Validate) {
	MustRegisterCustomRule(validate, "minDate", customrules.MinDateValidator)
	MustRegisterCustomRule(validate, "maxDate", customrules.MaxDateValidator)

	MustRegisterCustomRule(validate, "minDateDiff", customrules.MinDateDiffValidator)
	MustRegisterCustomRule(validate, "maxDateDiff", customrules.MaxDateDiffValidator)

	MustRegisterCustomRule(validate, "score", customrules.ScoreValidator)
}
