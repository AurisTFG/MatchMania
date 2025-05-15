// nolint
package config_test

import (
	"MatchManiaAPI/config"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupServer_DevMode(t *testing.T) {
	env := &config.Env{
		IsDev:     true,
		IsProd:    false,
		ClientURL: "http://localhost:3000",
	}

	server, err := config.SetupServer(env)

	assert.NoError(t, err)
	assert.NotNil(t, server)
	assert.Equal(t, gin.DebugMode, gin.Mode())
}

func TestSetupServer_ProdMode(t *testing.T) {
	env := &config.Env{
		IsDev:     false,
		IsProd:    true,
		ClientURL: "https://production.com",
	}

	server, err := config.SetupServer(env)
	
	assert.NoError(t, err)
	assert.NotNil(t, server)
	assert.Equal(t, gin.ReleaseMode, gin.Mode())
}
