// nolint
package config_test

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/docs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupSwagger_InDevMode(t *testing.T) {
	gin.SetMode(gin.TestMode)

	env := &config.Env{
		IsDev:     true,
		ServerURL: "localhost:8080",
	}

	server := gin.New()

	config.SetupSwagger(server, env)

	assert.Equal(t, "localhost:8080", docs.SwaggerInfo.Host)

	req := httptest.NewRequest(http.MethodGet, "/swagger/index.html", nil)
	resp := httptest.NewRecorder()

	server.ServeHTTP(resp, req)

	assert.NotEqual(t, http.StatusNotFound, resp.Code)
}

func TestSetupSwagger_NotInDevMode(t *testing.T) {
	gin.SetMode(gin.TestMode)

	env := &config.Env{
		IsDev:     false,
		ServerURL: "production.com",
	}

	server := gin.New()

	config.SetupSwagger(server, env)
	assert.NotEqual(t, "production.com", docs.SwaggerInfo.Host)

	req := httptest.NewRequest(http.MethodGet, "/swagger/index.html", nil)
	resp := httptest.NewRecorder()

	server.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}
