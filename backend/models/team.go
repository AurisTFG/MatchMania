package models

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Name string `gorm:"not null"`
	Elo  uint   `gorm:"not null"`

	HomeResults []Result `gorm:"foreignKey:TeamID"`
	AwayResults []Result `gorm:"foreignKey:OpponentTeamID"`

	SeasonID uint      `gorm:"not null"`
	UserUUID uuid.UUID `gorm:"type:uuid;not null"`
}

type TeamDto struct {
	ID   uint   `example:"6"        json:"id"`
	Name string `example:"BIG Clan" json:"name"`
	Elo  uint   `example:"1200"     json:"elo"`

	SeasonID uint      `example:"5"                                    json:"seasonId"`
	UserUUID uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000" json:"userUuid"`
}

type CreateTeamDto struct {
	Name string `example:"BIG Clan" json:"name" validate:"required,min=3,max=100"`
}

type UpdateTeamDto struct {
	Name string `example:"BIG Clan" json:"name" validate:"required,min=3,max=100"`
}

func (t *Team) ToDto() TeamDto {
	return TeamDto{
		ID:       t.ID,
		Name:     t.Name,
		Elo:      t.Elo,
		SeasonID: t.SeasonID,
		UserUUID: t.UserUUID,
	}
}

func ToTeamDtos(teams []Team) []TeamDto {
	teamDtos := make([]TeamDto, len(teams))

	for i, team := range teams {
		teamDtos[i] = team.ToDto()
	}

	return teamDtos
}

func (dto *CreateTeamDto) ToTeam() Team {
	return Team{
		Name: dto.Name,
	}
}

func (dto *UpdateTeamDto) ToTeam() Team {
	return Team{
		Name: dto.Name,
	}
}

func (dto *CreateTeamDto) Validate() error {
	var validate = validator.New()

	return teamValidationErrorHandler(validate.Struct(dto))
}

func (dto *UpdateTeamDto) Validate() error {
	var validate = validator.New()

	return teamValidationErrorHandler(validate.Struct(dto))
}

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
