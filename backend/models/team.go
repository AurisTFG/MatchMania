package models

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Name string `gorm:"not null"`
	Elo  uint   `gorm:"not null"`

	SeasonID uint `gorm:"not null"`

	// Players     []Player `gorm:"foreignKey:TeamID"`
	HomeResults []Result `gorm:"foreignKey:TeamID"`
	AwayResults []Result `gorm:"foreignKey:OpponentTeamID"`
}

type TeamDto struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Elo  uint   `json:"elo"`
}

type CreateTeamDto struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}

type UpdateTeamDto struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}

func (t *Team) ToDto() TeamDto {
	return TeamDto{
		ID:   t.ID,
		Name: t.Name,
		Elo:  t.Elo,
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

	for _, err := range err.(validator.ValidationErrors) {
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
