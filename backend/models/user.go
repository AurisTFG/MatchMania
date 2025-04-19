package models

import (
	"MatchManiaAPI/models/enums"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel

	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	// Country   string // TODO: Add country to user
	Role     enums.Role `gorm:"type:role;default:'user'"`
	Password string

	Sessions []Session `gorm:"constraint:OnDelete:CASCADE;"`
	Seasons  []Season  `gorm:"constraint:OnDelete:CASCADE;"`
	Teams    []Team    `gorm:"constraint:OnDelete:CASCADE;"`
	Results  []Result  `gorm:"constraint:OnDelete:CASCADE;"`
}

func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}

// func (dto *SignUpDto) Validate() error {
// 	var validate = validator.New()

// 	return userValidationErrorHandler(validate.Struct(dto))
// }

// func (dto *LoginDto) Validate() error {
// 	var validate = validator.New()

// 	return userValidationErrorHandler(validate.Struct(dto))
// }

// func (dto *UpdateUserDto) Validate() error {
// 	var validate = validator.New()

// 	return userValidationErrorHandler(validate.Struct(dto))
// }

func userValidationErrorHandler(err error) error {
	if err == nil {
		return nil
	}

	var errorMessage string
	var validationErrors validator.ValidationErrors

	if !errors.As(err, &validationErrors) {
		return errors.New("validation error")
	}

	for _, err := range validationErrors {
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
				errorMessage = "Password must be at least 8 characters long."
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
