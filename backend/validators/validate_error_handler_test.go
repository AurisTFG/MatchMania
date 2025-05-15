// nolint
package validators_test

import (
	"MatchManiaAPI/validators"
	"errors"
	"reflect"
	"testing"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type fakeFieldError struct {
	tag         string
	structField string
	value       interface{}
	param       string
}

func (f fakeFieldError) Tag() string             { return f.tag }
func (f fakeFieldError) StructField() string     { return f.structField }
func (f fakeFieldError) Value() interface{}      { return f.value }
func (f fakeFieldError) Param() string           { return f.param }
func (f fakeFieldError) Namespace() string       { return "" }
func (f fakeFieldError) StructNamespace() string { return "" }
func (f fakeFieldError) Field() string           { return "" }
func (f fakeFieldError) StructFieldOK() bool     { return true }
func (f fakeFieldError) Kind() reflect.Kind      { return reflect.String }
func (f fakeFieldError) Type() reflect.Type      { return nil }
func (f fakeFieldError) ValueOK() bool           { return true }
func (f fakeFieldError) ParamOK() bool           { return true }
func (f fakeFieldError) ActualTag() string       { return f.tag }
func (f fakeFieldError) Translation() string     { return "" }
func (f fakeFieldError) Error() string           { return f.tag + " error" }

func (f fakeFieldError) Translate(ut ut.Translator) string {
	return f.tag + " translation"
}

func createValidationError(tag, field string, value interface{}, param string) validator.ValidationErrors {
	return validator.ValidationErrors{fakeFieldError{
		tag:         tag,
		structField: field,
		value:       value,
		param:       param,
	}}
}

func TestValidateErrorHandler_NilError(t *testing.T) {
	err := validators.ValidateErrorHandler(nil)
	assert.NoError(t, err)
}

func TestValidateErrorHandler_UnknownErrorType(t *testing.T) {
	err := validators.ValidateErrorHandler(errors.New("some random error"))
	assert.Error(t, err)
	assert.Equal(t, "unknown validation error", err.Error())
}

func TestValidateErrorHandler_RequiredTag(t *testing.T) {
	err := validators.ValidateErrorHandler(createValidationError("required", "Name", "", ""))
	assert.Equal(t, "Name is required.", err.Error())
}

func TestValidateErrorHandler_NefieldTag(t *testing.T) {
	err := validators.ValidateErrorHandler(createValidationError("nefield", "Password", "", "ConfirmPassword"))
	assert.Equal(t, "Password must not be equal to ConfirmPassword.", err.Error())
}

func TestValidateErrorHandler_MinTag(t *testing.T) {
	err := validators.ValidateErrorHandler(createValidationError("min", "Username", "abc", "5"))
	assert.Contains(t, err.Error(), "Username must be at least 5 characters long")
	assert.Contains(t, err.Error(), "Received: 3")
}

func TestValidateErrorHandler_MaxTag(t *testing.T) {
	err := validators.ValidateErrorHandler(createValidationError("max", "Username", "abcdefg", "5"))
	assert.Contains(t, err.Error(), "Username can be up to 5 characters long")
	assert.Contains(t, err.Error(), "Received: 7")
}

func TestValidateErrorHandler_EmailTag(t *testing.T) {
	err := validators.ValidateErrorHandler(createValidationError("email", "Email", "invalid", ""))
	assert.Equal(t, "Email must be a valid email address.", err.Error())
}

func TestValidateErrorHandler_UrlTag(t *testing.T) {
	err := validators.ValidateErrorHandler(createValidationError("url", "Website", "not-a-url", ""))
	assert.Equal(t, "Website must be a valid URL. Received: not-a-url", err.Error())
}

func TestValidateErrorHandler_UuidTag(t *testing.T) {
	err := validators.ValidateErrorHandler(createValidationError("uuid", "ID", "not-a-uuid", ""))
	assert.Equal(t, "ID must be a valid UUID. Received: not-a-uuid", err.Error())
}

func TestValidateErrorHandler_ScoreTag(t *testing.T) {
	err := validators.ValidateErrorHandler(createValidationError("score", "Score", "150", ""))
	assert.Equal(t, "Score must be between 0 and 100. Received: 150", err.Error())
}

func TestValidateErrorHandler_MinDateTag_Past(t *testing.T) {
	err := validators.ValidateErrorHandler(createValidationError("minDate", "StartDate", "-7", "-7"))
	assert.Equal(t, "StartDate cannot be more than 7 days in the past.", err.Error())
}

func TestValidateErrorHandler_MinDateTag_Future(t *testing.T) {
	err := validators.ValidateErrorHandler(createValidationError("minDate", "StartDate", "7", "7"))
	assert.Equal(t, "StartDate must be at least 7 days in the future.", err.Error())
}

func TestValidateErrorHandler_MaxDateTag_Past(t *testing.T) {
	err := validators.ValidateErrorHandler(createValidationError("maxDate", "EndDate", "-5", "-5"))
	assert.Equal(t, "EndDate cannot be more than 5 days in the past.", err.Error())
}

func TestValidateErrorHandler_MaxDateTag_Future(t *testing.T) {
	err := validators.ValidateErrorHandler(createValidationError("maxDate", "EndDate", "10", "10"))
	assert.Equal(t, "EndDate must be at least 10 days in the future.", err.Error())
}

func TestValidateErrorHandler_DateDiffTag(t *testing.T) {
	err := validators.ValidateErrorHandler(createValidationError("dateDiff", "EndDate", "2", "2"))
	assert.Equal(t, "End Date must be after Start Date and be 2 days apart.", err.Error())
}

func TestValidateErrorHandler_DefaultCase(t *testing.T) {
	err := validators.ValidateErrorHandler(createValidationError("someUnknownTag", "Field", "", ""))
	assert.Equal(t, "Field failed on someUnknownTag validation.", err.Error())
}
