package validators

import (
	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate = nil

func init() {
	validator := validator.New()

	MustRegisterCustomRules(validator)

	Validator = validator
}
