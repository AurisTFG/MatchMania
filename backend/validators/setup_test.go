// nolint
package validators_test

import (
	"MatchManiaAPI/validators"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorInitialization(t *testing.T) {
	assert.NotNil(t, validators.Validator, "Validator should be initialized and not nil")
}
