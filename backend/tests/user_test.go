package tests

import (
	"MatchManiaAPI/models"
	"testing"
)

func TestPasswordHashing(t *testing.T) {
	// Create a new user
	user := models.User{
		Email:    "test@example.com",
		Password: "password",
	}

	if err := user.HashPassword(); err != nil {
		t.Errorf("Error hashing password")
	}

	if user.Password == "password" {
		t.Errorf("Password was not hashed")
	}

	if !user.ComparePassword("password") {
		t.Errorf("Password was not compared correctly")
	}
}
