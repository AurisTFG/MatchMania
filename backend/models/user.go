package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UUID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"unique"`
	Email     string         `gorm:"unique"`
	//Country   string // TODO: Add country to user
	Role     Role `gorm:"type:role;default:'user'"`
	Password string

	Sessions []Session `gorm:"foreignKey:UserUUID"`
	Seasons  []Season  `gorm:"foreignKey:UserUUID"`
	Teams    []Team    `gorm:"foreignKey:UserUUID"`
	Results  []Result  `gorm:"foreignKey:UserUUID"`
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

type UserDto struct {
	UUID     uuid.UUID `json:"id" example:"526432ea-822b-4b5b-b1b3-34e8ab453e03"`
	Username string    `json:"username" example:"john_doe_123"`
	Email    string    `json:"email" example:"email@example.com"`
	Role     Role      `json:"role" example:"admin"`
}

type SignUpDto struct {
	Username string `json:"username" validate:"required,min=3,max=100" example:"john_doe_123"`
	Email    string `json:"email" validate:"required,email" example:"email@example.com"`
	Password string `json:"password" validate:"required,min=8,max=255" example:"VeryStrongPassword"`
}

type LoginDto struct {
	Email    string `json:"email" validate:"required,email" example:"email@example.com"`
	Password string `json:"password" validate:"required,min=8,max=255" example:"VeryStrongPassword"`
}

type UpdateUserDto struct {
	Username string `json:"username" validate:"omitempty,min=3,max=100" example:"john_doe_123"`
	Email    string `json:"email" validate:"omitempty,email" example:"email@example.com"`
	Password string `json:"password" validate:"omitempty,min=8,max=255" example:"VeryStrongPassword"`
}

func (dto *SignUpDto) ToUser() User {
	return User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func (dto *UpdateUserDto) ToUser() User {
	return User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func (u *User) ToDto() UserDto {
	return UserDto{
		UUID:     u.UUID,
		Username: u.Username,
		Role:     u.Role,
		Email:    u.Email,
	}
}

func ToUserDtos(users []User) []UserDto {
	userDtos := make([]UserDto, len(users))

	for i, user := range users {
		userDtos[i] = user.ToDto()
	}

	return userDtos
}

func (dto *SignUpDto) Validate() error {
	var validate = validator.New()

	return userValidationErrorHandler(validate.Struct(dto))
}

func (dto *LoginDto) Validate() error {
	var validate = validator.New()

	return userValidationErrorHandler(validate.Struct(dto))
}

func (dto *UpdateUserDto) Validate() error {
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
