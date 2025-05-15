// nolint
package utils_test

import (
	"MatchManiaAPI/utils"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func createTestJWT(claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Using "secret" to sign, even if it's not verified (ParseUnverified ignores signature)
	tokenStr, _ := token.SignedString([]byte("secret"))
	return tokenStr
}

func TestGetJwtClaim_Success(t *testing.T) {
	tokenStr := createTestJWT(jwt.MapClaims{
		"username": "aurimas",
		"age":      float64(23),
	})

	claim, err := utils.GetJwtClaim[string](tokenStr, "username")
	assert.NoError(t, err)
	assert.Equal(t, "aurimas", claim)

	age, err := utils.GetJwtClaim[float64](tokenStr, "age")
	assert.NoError(t, err)
	assert.Equal(t, float64(23), age)
}

func TestGetJwtClaim_InvalidJWT(t *testing.T) {
	invalidToken := "this.is.not.valid.jwt"

	_, err := utils.GetJwtClaim[string](invalidToken, "username")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "parsing JWT")
}

func TestGetJwtClaim_ClaimNotFound(t *testing.T) {
	tokenStr := createTestJWT(jwt.MapClaims{
		"username": "aurimas",
	})

	_, err := utils.GetJwtClaim[string](tokenStr, "nonexistent")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "claim 'nonexistent' not found")
}

func TestGetJwtClaim_WrongType(t *testing.T) {
	tokenStr := createTestJWT(jwt.MapClaims{
		"username": "aurimas",
	})

	_, err := utils.GetJwtClaim[int](tokenStr, "username")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "claim 'username' is not of expected type")
}

func TestGetJwtClaim_InvalidClaimsType(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{})
	tokenStr, _ := token.SignedString([]byte("secret"))

	_, err := utils.GetJwtClaim[string](tokenStr, "username")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "claim 'username' not found")
}
