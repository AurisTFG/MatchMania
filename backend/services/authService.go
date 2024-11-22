package services

import (
	"MatchManiaAPI/initializers"
	"MatchManiaAPI/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(SignUpDto *models.SignUpDto) (*models.User, error) {
	newUser := SignUpDto.ToUser()

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser.Password = string(hash)

	result := initializers.DB.Create(&newUser)

	return &newUser, result.Error
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	result := initializers.DB.Where("email = ?", email).First(&user)

	return &user, result.Error
}

func GetUserByID(userID uint) (*models.User, error) {
	var user models.User

	result := initializers.DB.Where("id = ?", userID).First(&user)

	return &user, result.Error
}
