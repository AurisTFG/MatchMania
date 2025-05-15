// nolint
package customrules

import (
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type maxDateTestStruct struct {
	Date time.Time `validate:"maxDate=7"`
}

func TestMaxDateValidator(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("maxDate", MaxDateValidator)
	assert.NoError(t, err)

	now := time.Now()

	tests := []struct {
		name      string
		input     maxDateTestStruct
		wantError bool
	}{
		{
			name: "Valid_BeforeMaxDate",
			input: maxDateTestStruct{
				Date: now.AddDate(0, 0, 6),
			},
			wantError: false,
		},
		{
			name: "Invalid_AfterMaxDate",
			input: maxDateTestStruct{
				Date: now.AddDate(0, 0, 8),
			},
			wantError: true,
		},
		{
			name: "Invalid_ExactlyMaxDate",
			input: maxDateTestStruct{
				Date: now.AddDate(0, 0, 7),
			},
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

func TestMaxDateValidator_InvalidType(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("maxDate", MaxDateValidator)
	assert.NoError(t, err)

	type invalidStruct struct {
		Date string `validate:"maxDate=7"`
	}

	input := invalidStruct{Date: "not-a-date"}

	err = validate.Struct(input)
	assert.Error(t, err)
}

func TestMaxDateValidator_InvalidParam(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("maxDate", MaxDateValidator)
	assert.NoError(t, err)

	type invalidStruct struct {
		Date time.Time `validate:"maxDate=INVALID_PARAMETER"`
	}

	input := invalidStruct{
		Date: time.Now().Add(-24 * time.Hour),
	}

	err = validate.Struct(input)
	assert.NoError(t, err)
}
