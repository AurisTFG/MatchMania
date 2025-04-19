package models

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Team struct {
	BaseModel

	Name string `gorm:"not null"`
	Elo  uint   `gorm:"not null"`

	HomeResults []Result `gorm:"foreignKey:TeamId;constraint:OnDelete:CASCADE"`
	AwayResults []Result `gorm:"foreignKey:OpponentTeamId;constraint:OnDelete:CASCADE"`

	SeasonId uuid.UUID `gorm:"not null"`
	UserId   uuid.UUID `gorm:"not null"`
}

// func (dto *CreateTeamDto) Validate() error {
// 	var validate = validator.New()

// 	return teamValidationErrorHandler(validate.Struct(dto))
// }

// func (dto *UpdateTeamDto) Validate() error {
// 	var validate = validator.New()

// 	return teamValidationErrorHandler(validate.Struct(dto))
// }

func teamValidationErrorHandler(err error) error {
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
		case "Name":
			switch tag {
			case "required":
				errorMessage = "Name is required."
			case "min":
				errorMessage = "Name must be at least 3 characters long."
			case "max":
				errorMessage = "Name can be up to 100 characters long."
			}
		}

		if errorMessage == "" {
			errorMessage = fmt.Sprintf("Validation failed on field %s.", field)
		}
	}

	return errors.New(errorMessage)
}
