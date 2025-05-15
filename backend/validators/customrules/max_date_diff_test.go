// nolint
package customrules

import (
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type maxDateDiffTestStruct struct {
	StartDate time.Time
	EndDate   time.Time `validate:"maxDateDiff=90"`
}

func TestMaxDateDiffValidator(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("maxDateDiff", MaxDateDiffValidator)
	assert.NoError(t, err)

	now := time.Now()

	tests := []struct {
		name      string
		input     maxDateDiffTestStruct
		wantError bool
	}{
		{
			name: "Valid_ExactlyMaxDays",
			input: maxDateDiffTestStruct{
				StartDate: now,
				EndDate:   now.AddDate(0, 0, 90),
			},
			wantError: false,
		},
		{
			name: "Valid_LessThanMaxDays",
			input: maxDateDiffTestStruct{
				StartDate: now,
				EndDate:   now.AddDate(0, 0, 60),
			},
			wantError: false,
		},
		{
			name: "Invalid_MoreThanMaxDays",
			input: maxDateDiffTestStruct{
				StartDate: now,
				EndDate:   now.AddDate(0, 0, 91),
			},
			wantError: true,
		},
		{
			name: "Invalid_EndBeforeStart",
			input: maxDateDiffTestStruct{
				StartDate: now,
				EndDate:   now.AddDate(0, 0, -1),
			},
			wantError: false,
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

func TestMaxDateDiffValidator_InvalidType(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("maxDateDiff", MaxDateDiffValidator)
	assert.NoError(t, err)

	type invalidTypeStruct struct {
		StartDate string
		EndDate   string `validate:"maxDateDiff=90"`
	}

	input := invalidTypeStruct{
		StartDate: "not-a-date",
		EndDate:   "still-not-a-date",
	}

	err = validate.Struct(input)
	assert.Error(t, err)
}

func TestMaxDateDiffValidator_InvalidParam(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("maxDateDiff", MaxDateDiffValidator)
	assert.NoError(t, err)

	now := time.Now()

	type invalidParamStruct struct {
		StartDate time.Time
		EndDate   time.Time `validate:"maxDateDiff=invalid"`
	}

	input := invalidParamStruct{
		StartDate: now,
		EndDate:   now.AddDate(0, 0, 1),
	}

	err = validate.Struct(input)
	assert.NoError(t, err)
}
