// nolint
package utils_test

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"MatchManiaAPI/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

type errorCloser struct{}

func (e *errorCloser) Read(p []byte) (n int, err error) {
	copy(p, "ok")
	return 2, io.EOF
}

func (e *errorCloser) Close() error {
	return errors.New("forced close error")
}

func TestHttpRequest_Success(t *testing.T) {
	// Arrange
	expectedResponse := map[string]string{"message": "hello"}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := server.Client()
	req, err := http.NewRequest(http.MethodPost, server.URL, nil)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	body := map[string]string{"key": "value"}

	// Act
	response, err := utils.HttpRequest[map[string]string](client, req, body)

	// Assert
	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, expectedResponse, *response)
}

func TestHttpRequest_EncodingError(t *testing.T) {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, "http://example.com", nil)

	body := make(chan int) // channels cannot be marshaled into JSON

	response, err := utils.HttpRequest[map[string]string](client, req, body)

	assert.Nil(t, response)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "encoding request body")
}

func TestHttpRequest_SendingError(t *testing.T) {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, "http://invalid-url", nil)

	response, err := utils.HttpRequest[map[string]string](client, req, nil)

	assert.Nil(t, response)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "sending request")
}

func TestHttpRequest_BadStatusCode(t *testing.T) {
	// Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Bad Request`))
	}))
	defer server.Close()

	client := server.Client()
	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)

	// Act
	response, err := utils.HttpRequest[map[string]string](client, req, nil)

	// Assert
	assert.Nil(t, response)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "bad response")
	assert.Contains(t, err.Error(), "Bad Request")
}

func TestHttpRequest_DecodeError(t *testing.T) {
	// Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`not a json`))
	}))
	defer server.Close()

	client := server.Client()
	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)

	// Act
	response, err := utils.HttpRequest[map[string]string](client, req, nil)

	// Assert
	assert.Nil(t, response)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "decoding response")
}

func TestHttpRequest_CloseBodyError(t *testing.T) {
	client := &http.Client{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	resp.Body = &errorCloser{}

	badClient := &http.Client{
		Transport: roundTripperFunc(func(r *http.Request) (*http.Response, error) {
			return resp, nil
		}),
	}

	response, err := utils.HttpRequest[map[string]string](badClient, req, nil)

	assert.Nil(t, response)
	assert.Error(t, err)
}

func TestGetBasicAuthHeader_Success(t *testing.T) {
	email := "user@example.com"
	password := "password123"

	header := utils.GetBasicAuthHeader(email, password)

	expectedCredentials := base64.StdEncoding.EncodeToString([]byte(email + ":" + password))
	expectedHeader := "Basic " + expectedCredentials

	assert.Equal(t, expectedHeader, header)
}

func TestGetBasicAuthHeader_Panic(t *testing.T) {
	assert.PanicsWithValue(t, "Email or password is empty", func() {
		utils.GetBasicAuthHeader("", "password")
	})

	assert.PanicsWithValue(t, "Email or password is empty", func() {
		utils.GetBasicAuthHeader("email", "")
	})
}

func TestGetUbisoftAuthHeader_Success(t *testing.T) {
	ticket := "ticket123"

	header := utils.GetUbisoftAuthHeader(ticket)

	assert.Equal(t, "ubi_v1 t=ticket123", header)
}

func TestGetUbisoftAuthHeader_Panic(t *testing.T) {
	assert.PanicsWithValue(t, "Ticket is empty", func() {
		utils.GetUbisoftAuthHeader("")
	})
}

func TestGetNadeoAuthHeader_Success(t *testing.T) {
	token := "token123"

	header := utils.GetNadeoAuthHeader(token)

	assert.Equal(t, "nadeo_v1 t=token123", header)
}

func TestGetNadeoAuthHeader_Panic(t *testing.T) {
	assert.PanicsWithValue(t, "Token is empty", func() {
		utils.GetNadeoAuthHeader("")
	})
}
