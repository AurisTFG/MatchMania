// nolint
package config_test

import (
	"MatchManiaAPI/config"
	"os"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func resetViper() {
	viper.Reset()
}

func SetupValidEnvVars() {
	os.Setenv("DATABASE_URL", "postgres://localhost:5432/test")
	os.Setenv("SERVER_URL", "http://localhost:8080")
	os.Setenv("CLIENT_URL", "http://localhost:3000")
	os.Setenv("JWT_ACCESS_TOKEN_SECRET", "secret")
	os.Setenv("JWT_REFRESH_TOKEN_SECRET", "secret")
	os.Setenv("JWT_ISSUER", "issuer")
	os.Setenv("JWT_AUDIENCE", "audience")
	os.Setenv("JWT_ACCESS_TOKEN_DURATION", "15m")
	os.Setenv("JWT_REFRESH_TOKEN_DURATION", "7h")
	os.Setenv("TRACKMANIA_OAUTH_CLIENT_ID", "clientid")
	os.Setenv("TRACKMANIA_OAUTH_CLIENT_SECRET", "clientsecret")
	os.Setenv("TRACKMANIA_OAUTH_SCOPES", "scope1,scope2")
	os.Setenv("TRACKMANIA_OAUTH_REDIRECT_URL", "http://localhost/callback")
	os.Setenv("TRACKMANIA_API_EMAIL", "email@example.com")
	os.Setenv("TRACKMANIA_API_PASSWORD", "password")
	os.Setenv("TRACKMANIA_API_CLUB_ID", "12345")

	// Development only
	os.Setenv("USER_EMAIL", "user@example.com")
	os.Setenv("USER_PASSWORD", "userpass")
	os.Setenv("ADMIN_EMAIL", "admin@example.com")
	os.Setenv("ADMIN_PASSWORD", "adminpass")
}

func TestLoadEnv_InvalidEnvironment(t *testing.T) {
	resetViper()
	env, err := config.LoadEnv("staging")
	assert.Nil(t, env)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid environment name")
}

func TestLoadEnv_UnmarshalError(t *testing.T) {
	resetViper()
	SetupValidEnvVars()

	os.Setenv("TRACKMANIA_API_CLUB_ID", "invalid_integer")

	env, err := config.LoadEnv("prod")
	assert.Nil(t, env)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to decode config into struct")
}

func TestEnv_Validate_CompleteConfig_Dev(t *testing.T) {
	resetViper()
	SetupValidEnvVars()

	env, err := config.LoadEnv("dev")
	require.NoError(t, err)
	require.NotNil(t, env)

	assert.True(t, env.IsDev)
	assert.False(t, env.IsProd)
	assert.Equal(t, "postgres://localhost:5432/test", env.DatabaseURL)
	assert.Equal(t, 12345, env.TrackmaniaApiClubId)

	assert.Equal(t, 15*time.Minute, env.JWTAccessTokenDuration)
	assert.Equal(t, 7*time.Hour, env.JWTRefreshTokenDuration)
}

func TestEnv_Validate_ProdModeWithoutDevCreds(t *testing.T) {
	resetViper()
	SetupValidEnvVars()

	os.Unsetenv("USER_EMAIL")
	os.Unsetenv("USER_PASSWORD")
	os.Unsetenv("ADMIN_EMAIL")
	os.Unsetenv("ADMIN_PASSWORD")

	env, err := config.LoadEnv("prod")
	require.NoError(t, err)
	require.NotNil(t, env)

	assert.False(t, env.IsDev)
	assert.True(t, env.IsProd)
}

func TestEnv_Validate_MissingDatabaseUrl(t *testing.T) {
	resetViper()
	SetupValidEnvVars()
	os.Unsetenv("DATABASE_URL")

	_, err := config.LoadEnv("prod")

	require.Error(t, err)
}

func TestEnv_Validate_MissingServerUrl(t *testing.T) {
	resetViper()
	SetupValidEnvVars()
	os.Unsetenv("SERVER_URL")

	_, err := config.LoadEnv("prod")

	require.Error(t, err)
}

func TestEnv_Validate_MissingClientUrl(t *testing.T) {
	resetViper()
	SetupValidEnvVars()
	os.Unsetenv("CLIENT_URL")

	_, err := config.LoadEnv("prod")

	require.Error(t, err)
}

func TestEnv_Validate_MissingJwtAccessTokenSecret(t *testing.T) {
	resetViper()
	SetupValidEnvVars()
	os.Unsetenv("JWT_ACCESS_TOKEN_SECRET")

	_, err := config.LoadEnv("prod")

	require.Error(t, err)
}

func TestEnv_Validate_TrackmaniaOAuthClientID(t *testing.T) {
	resetViper()
	SetupValidEnvVars()
	os.Unsetenv("TRACKMANIA_OAUTH_CLIENT_ID")

	_, err := config.LoadEnv("prod")

	require.Error(t, err)
}

func TestEnv_Validate_TrackmaniaApiEmail(t *testing.T) {
	resetViper()
	SetupValidEnvVars()
	os.Unsetenv("TRACKMANIA_API_EMAIL")

	_, err := config.LoadEnv("prod")

	require.Error(t, err)
}

func TestEnv_Validate_UserEmail(t *testing.T) {
	resetViper()
	SetupValidEnvVars()
	os.Unsetenv("USER_EMAIL")

	_, err := config.LoadEnv("dev")

	require.Error(t, err)
}
