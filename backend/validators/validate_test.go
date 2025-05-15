// nolint
package validators_test

import (
	"MatchManiaAPI/validators"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

// Setup your validator.Validator instance if needed
func init() {
	validators.Validator = validator.New()
}

// --- Test helper structs ---

type ValidStruct struct {
	Name string `validate:"required"`
}

type InvalidStruct struct {
	Name string `validate:"required"`
}

func TestValidate_Success(t *testing.T) {
	obj := ValidStruct{
		Name: "Aurimas",
	}

	err := validators.Validate(obj)

	assert.NoError(t, err)
}

func TestValidate_Error(t *testing.T) {
	obj := InvalidStruct{
		Name: "",
	}

	err := validators.Validate(obj)

	assert.Error(t, err)
}
