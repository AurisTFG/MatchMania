package validators

import (
	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	Validator = validator.New()

	Validator.RegisterValidation("notadmin", func(fl validator.FieldLevel) bool {
		return fl.Field().String() != "admin"
	})
}

func Validate(s any) error {
	err := Validator.Struct(s)
	if err == nil {
		return nil
	}

	// var messages []string
	// for _, e := range err.(validator.ValidationErrors) {
	// 	messages = append(messages, e.Field()+" failed on "+e.Tag())
	// }
	return err
}
