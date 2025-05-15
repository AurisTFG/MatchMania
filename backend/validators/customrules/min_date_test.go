// nolint
package customrules

import (
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type minDateTestStruct struct {
	Date time.Time `validate:"minDate=-7"`
}

func TestMinDateValidator(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("minDate", MinDateValidator)
	assert.NoError(t, err)

	now := time.Now()

	tests := []struct {
		name      string
		input     minDateTestStruct
		wantError bool
	}{
		{
			name:      "Valid_After_MinDate",
			input:     minDateTestStruct{Date: now.AddDate(0, 0, -6)},
			wantError: false,
		},
		{
			name:      "Invalid_Before_MinDate",
			input:     minDateTestStruct{Date: now.AddDate(0, 0, -8)},
			wantError: true,
		},
		{
			name:      "Invalid_Exactly_MinDate",
			input:     minDateTestStruct{Date: now.AddDate(0, 0, -7)},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Struct(tt.input)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestMinDateValidator_InvalidType(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("minDate", MinDateValidator)
	assert.NoError(t, err)

	type invalidStruct struct {
		Date string `validate:"minDate=-7"`
	}

	input := invalidStruct{Date: "not-a-date"}

	err = validate.Struct(input)
	assert.Error(t, err)
}

func TestMinDateValidator_InvalidParam(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("minDate", MinDateValidator)
	assert.NoError(t, err)

	type invalidStruct struct {
		Date time.Time `validate:"minDate=INVALID_PARAMETER"`
	}

	input := invalidStruct{
		Date: time.Now().Add(24 * time.Hour),
	}

	err = validate.Struct(input)
	assert.NoError(t, err)
}
