// nolint
package httpresponses_test

import (
	"MatchManiaAPI/models/dtos/responses/errors"
	"MatchManiaAPI/utils/httpresponses"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	return c, w
}

func TestOK(t *testing.T) {
	c, w := createTestContext()

	obj := map[string]string{"key": "value"}

	httpresponses.OK(c, obj)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, obj, response)
}

func TestCreated(t *testing.T) {
	c, w := createTestContext()

	httpresponses.Created(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "null", w.Body.String())
}

func TestNoContent(t *testing.T) {
	c, w := createTestContext()

	httpresponses.NoContent(c)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.String()) 
}

func TestBadRequest(t *testing.T) {
	c, w := createTestContext()

	msg := "bad input"

	httpresponses.BadRequest(c, msg)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var errResp errors.ErrorDto
	err := json.Unmarshal(w.Body.Bytes(), &errResp)
	assert.NoError(t, err)
	assert.Equal(t, msg, errResp.Message)
}

func TestUnauthorized(t *testing.T) {
	c, w := createTestContext()

	httpresponses.Unauthorized(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var errResp errors.ErrorDto
	err := json.Unmarshal(w.Body.Bytes(), &errResp)
	assert.NoError(t, err)
	assert.Equal(t, "Unauthorized access", errResp.Message)
}

func TestForbidden(t *testing.T) {
	c, w := createTestContext()

	httpresponses.Forbidden(c)

	assert.Equal(t, http.StatusForbidden, w.Code)

	var errResp errors.ErrorDto
	err := json.Unmarshal(w.Body.Bytes(), &errResp)
	assert.NoError(t, err)
	assert.Equal(t, "You do not have permission to access this resource", errResp.Message)
}

func TestNotFound(t *testing.T) {
	c, w := createTestContext()

	msg := "resource not found"

	httpresponses.NotFound(c, msg)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var errResp errors.ErrorDto
	err := json.Unmarshal(w.Body.Bytes(), &errResp)
	assert.NoError(t, err)
	assert.Equal(t, msg, errResp.Message)
}

func TestUnprocessableEntity(t *testing.T) {
	c, w := createTestContext()

	msg := "unprocessable entity"

	httpresponses.UnprocessableEntity(c, msg)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

	var errResp errors.ErrorDto
	err := json.Unmarshal(w.Body.Bytes(), &errResp)
	assert.NoError(t, err)
	assert.Equal(t, msg, errResp.Message)
}

func TestInternalServerError(t *testing.T) {
	c, w := createTestContext()

	msg := "server error"

	httpresponses.InternalServerError(c, msg)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var errResp errors.ErrorDto
	err := json.Unmarshal(w.Body.Bytes(), &errResp)
	assert.NoError(t, err)
	assert.Equal(t, msg, errResp.Message)
}
