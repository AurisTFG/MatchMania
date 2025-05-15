// nolint
package customrules

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type scoreTestStruct struct {
	Score string `validate:"score"`
}

func TestScoreValidator(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("score", ScoreValidator)
	assert.NoError(t, err)

	tests := []struct {
		name    string
		input   scoreTestStruct
		wantErr bool
	}{
		{name: "Valid_Zero", input: scoreTestStruct{Score: "0"}, wantErr: false},
		{name: "Valid_Hundred", input: scoreTestStruct{Score: "100"}, wantErr: false},
		{name: "Valid_MidRange", input: scoreTestStruct{Score: "50"}, wantErr: false},
		{name: "Invalid_Negative", input: scoreTestStruct{Score: "-1"}, wantErr: true},
		{name: "Invalid_AboveHundred", input: scoreTestStruct{Score: "101"}, wantErr: true},
		{name: "Invalid_NonNumeric", input: scoreTestStruct{Score: "abc"}, wantErr: true},
		{name: "Invalid_EmptyString", input: scoreTestStruct{Score: ""}, wantErr: true},
		{name: "Invalid_Whitespace", input: scoreTestStruct{Score: " "}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Struct(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
