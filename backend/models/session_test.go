// nolint
package models_test

import (
	"MatchManiaAPI/models"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHashToken(t *testing.T) {
	session := &models.Session{
		LastRefreshToken: "testToken",
	}

	err := session.HashToken()
	require.NoError(t, err, "HashToken should not return an error")

	decoded, err := base64.StdEncoding.DecodeString(session.LastRefreshToken)
	require.NoError(t, err, "LastRefreshToken should be a valid base64 string")

	assert.Len(t, decoded, 16+32, "Decoded token should contain 16 bytes of salt and 32 bytes of hash")
}

func TestCompareToken(t *testing.T) {
	session := &models.Session{
		LastRefreshToken: "testToken",
	}

	err := session.HashToken()
	require.NoError(t, err, "HashToken should not return an error")

	valid := session.CompareToken("testToken")
	assert.True(t, valid, "CompareToken should return true for the correct token")

	invalid := session.CompareToken("wrongToken")
	assert.False(t, invalid, "CompareToken should return false for an incorrect token")
}
