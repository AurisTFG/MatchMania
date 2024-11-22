package models

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Role     Role   `gorm:"type:role;default:'user'"`
	Password string
}

type UserDto struct {
	ID       uint   `json:"id" example:"1"`
	Username string `json:"username" example:"AurisTFG"`
	Email    string `json:"email" example:"email@gmail.com"`
	Role     Role   `json:"role" example:"admin"`
}

type SignUpDto struct {
	Username string `json:"username" validate:"required,min=3,max=100" example:"AurisTFG"`
	Email    string `json:"email" validate:"required,email" example:"email@gmail.com"`
	Password string `json:"password" validate:"required,min=3,max=100" example:"VeryStrongPassword"`
}

type LoginDto struct {
	Email    string `json:"email" validate:"required,email" example:"email@gmail.com"`
	Password string `json:"password" validate:"required,min=3,max=100" example:"VeryStrongPassword"`
}

func (dto *SignUpDto) ToUser() User {
	return User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func (u *User) ToDto() UserDto {
	return UserDto{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Role:     u.Role,
	}
}

func (dto *SignUpDto) Validate() error {
	var validate = validator.New()

	return userValidationErrorHandler(validate.Struct(dto))
}

func (dto *LoginDto) Validate() error {
	var validate = validator.New()

	return userValidationErrorHandler(validate.Struct(dto))
}

func userValidationErrorHandler(err error) error {
	if err == nil {
		return nil
	}

	var errorMessage string

	for _, err := range err.(validator.ValidationErrors) {
		field := err.StructField()
		tag := err.Tag()

		switch field {
		case "Username":
			switch tag {
			case "required":
				errorMessage = "Username is required."
			case "min":
				errorMessage = "Username must be at least 3 characters long."
			case "max":
				errorMessage = "Username can be up to 100 characters long."
			}
		case "Email":
			switch tag {
			case "required":
				errorMessage = "Email is required."
			case "email":
				errorMessage = "Email must be a valid email address."
			}
		case "Password":
			switch tag {
			case "required":
				errorMessage = "Password is required."
			case "min":
				errorMessage = "Password must be at least 3 characters long."
			case "max":
				errorMessage = "Password can be up to 100 characters long."
			}
		}

		if errorMessage == "" {
			errorMessage = fmt.Sprintf("Validation failed on field %s.", field)
		}
	}

	return errors.New(errorMessage)
}
