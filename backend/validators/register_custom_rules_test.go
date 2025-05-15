// nolint
package validators_test

import (
	"MatchManiaAPI/validators"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func dummyValidator(fl validator.FieldLevel) bool {
	return true
}

func TestMustRegisterCustomRules_Success(t *testing.T) {
	validate := validator.New()

	assert.NotPanics(t, func() {
		validators.MustRegisterCustomRules(validate)
	})
}

func TestMustRegisterCustomRule_Success(t *testing.T) {
	v := validator.New()

	assert.NotPanics(t, func() {
		validators.MustRegisterCustomRule(v, "dummy", dummyValidator)
	})

	err := v.Var("test", "dummy")
	assert.NoError(t, err)
}

func TestMustRegisterCustomRule_Panic(t *testing.T) {
	v := validator.New()

	require.PanicsWithValue(t,
		"Failed to register custom validation rule: function Key cannot be empty",
		func() {
			validators.MustRegisterCustomRule(v, "", dummyValidator)
		},
	)
}
