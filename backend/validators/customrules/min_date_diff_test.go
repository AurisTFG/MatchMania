// nolint
package customrules

import (
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type minDateDiffTestStruct struct {
	StartDate time.Time
	EndDate   time.Time `validate:"minDateDiff=30"`
}

func TestMinDateDiffValidator(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("minDateDiff", MinDateDiffValidator)
	assert.NoError(t, err)

	now := time.Now()

	tests := []struct {
		name      string
		input     minDateDiffTestStruct
		wantError bool
	}{
		{
			name: "Valid_MinimumDifference",
			input: minDateDiffTestStruct{
				StartDate: now,
				EndDate:   now.AddDate(0, 0, 30), // exactly 30 days
			},
			wantError: false,
		},
		{
			name: "Valid_MoreThanMinimum",
			input: minDateDiffTestStruct{
				StartDate: now,
				EndDate:   now.AddDate(0, 0, 40), // more than 30 days
			},
			wantError: false,
		},
		{
			name: "Invalid_LessThanMinimum",
			input: minDateDiffTestStruct{
				StartDate: now,
				EndDate:   now.AddDate(0, 0, 29), // 29 days
			},
			wantError: true,
		},
		{
			name: "Invalid_ExactSameDay",
			input: minDateDiffTestStruct{
				StartDate: now,
				EndDate:   now, // same day
			},
			wantError: true,
		},
		{
			name: "Invalid_EndBeforeStart",
			input: minDateDiffTestStruct{
				StartDate: now,
				EndDate:   now.AddDate(0, 0, -1), // end before start
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

func TestMinDateDiffValidator_InvalidType(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("minDateDiff", MinDateDiffValidator)
	assert.NoError(t, err)

	type invalidTypeStruct struct {
		StartDate string
		EndDate   string `validate:"minDateDiff=30"`
	}

	input := invalidTypeStruct{
		StartDate: "not-a-date",
		EndDate:   "also-not-a-date",
	}

	err = validate.Struct(input)
	assert.Error(t, err)
}

func TestMinDateDiffValidator_InvalidParam(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("minDateDiff", MinDateDiffValidator)
	assert.NoError(t, err)

	now := time.Now()

	type invalidParamStruct struct {
		StartDate time.Time
		EndDate   time.Time `validate:"minDateDiff=invalid"`
	}

	input := invalidParamStruct{
		StartDate: now,
		EndDate:   now.AddDate(0, 0, 29),
	}

	err = validate.Struct(input)
	assert.Error(t, err)
}
