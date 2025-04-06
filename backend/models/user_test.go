package models_test

import (
	"MatchManiaAPI/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	correctPassword := "StrongPassword123"
	user := &models.User{
		Password: correctPassword,
	}

	err := user.HashPassword()
	require.NoError(t, err, "HashPassword should not return an error")
	assert.NotEqual(t, correctPassword, user.Password, "Password should be hashed")

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(correctPassword))
	assert.NoError(t, err, "Hashed password should match the original password")
}

func TestComparePassword(t *testing.T) {
	correctPassword := "StrongPassword123"
	wrongPassword := "WrongPassword123"
	user := &models.User{
		Password: correctPassword,
	}

	err := user.HashPassword()
	require.NoError(t, err, "HashPassword should not return an error")

	isMatch := user.ComparePassword(correctPassword)
	assert.True(t, isMatch, "ComparePassword should return true for the correct password")

	isMatch = user.ComparePassword(wrongPassword)
	assert.False(t, isMatch, "ComparePassword should return false for an incorrect password")
}
