package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Result struct {
	gorm.Model
	MatchStartDate time.Time `gorm:"not null"`
	MatchEndDate   time.Time `gorm:"not null"`
	Score          string    `gorm:"not null"`
	OpponentScore  string    `gorm:"not null"`

	Team         Team `gorm:"foreignKey:TeamID"`
	OpponentTeam Team `gorm:"foreignKey:OpponentTeamID"`

	UserUUID       uuid.UUID `gorm:"type:uuid;not null"`
	TeamID         uint      `gorm:"not null"`
	OpponentTeamID uint      `gorm:"not null"`
	SeasonID       uint      `gorm:"not null"`
}

type ResultDto struct {
	ID             uint      `json:"id" example:"7"`
	MatchStartDate time.Time `json:"matchStartDate" example:"2025-06-01T00:00:00Z"`
	MatchEndDate   time.Time `json:"matchEndDate" example:"2025-06-01T00:40:00Z"`
	Score          string    `json:"score" example:"16"`
	OpponentScore  string    `json:"opponentScore" example:"14"`
	TeamID         uint      `json:"team" example:"6"`
	OpponentTeamID uint      `json:"opponentTeam" example:"7"`
}

type CreateResultDto struct {
	MatchStartDate time.Time `json:"matchStartDate" validate:"required,startDate" example:"2025-06-01T00:00:00Z"`
	MatchEndDate   time.Time `json:"matchEndDate" validate:"required,endDate,dateDiff,gtfield=MatchStartDate" example:"2025-06-01T00:40:00Z"`
	Score          string    `json:"score" validate:"score" example:"16"`
	OpponentScore  string    `json:"opponentScore" validate:"score" example:"14"`
	OpponentTeamID uint      `json:"opponentTeamId" validate:"required" example:"7"`
}

type UpdateResultDto struct {
	MatchStartDate time.Time `json:"matchStartDate" validate:"required,startDate" example:"2025-06-01T00:00:00Z"`
	MatchEndDate   time.Time `json:"matchEndDate" validate:"required,endDate,dateDiff,gtfield=MatchStartDate" example:"2025-06-01T00:40:00Z"`
	Score          string    `json:"score" validate:"score" example:"16"`
	OpponentScore  string    `json:"opponentScore" validate:"score" example:"14"`
}

func (r *Result) ToDto() ResultDto {
	return ResultDto{
		ID:             r.ID,
		MatchStartDate: r.MatchStartDate,
		MatchEndDate:   r.MatchEndDate,
		Score:          r.Score,
		OpponentScore:  r.OpponentScore,
		TeamID:         r.TeamID,
		OpponentTeamID: r.OpponentTeamID,
	}
}

func ToResultDtos(results []Result) []ResultDto {
	resultDtos := make([]ResultDto, len(results))

	for i, result := range results {
		resultDtos[i] = result.ToDto()
	}

	return resultDtos
}

func (dto *CreateResultDto) ToResult() Result {
	return Result{
		MatchStartDate: dto.MatchStartDate,
		MatchEndDate:   dto.MatchEndDate,
		Score:          dto.Score,
		OpponentScore:  dto.OpponentScore,
		OpponentTeamID: dto.OpponentTeamID,
	}
}

func (dto *UpdateResultDto) ToResult() Result {
	return Result{
		MatchStartDate: dto.MatchStartDate,
		MatchEndDate:   dto.MatchEndDate,
		Score:          dto.Score,
		OpponentScore:  dto.OpponentScore,
	}
}

func startDateValidatorResult(fl validator.FieldLevel) bool {
	date := fl.Field().Interface().(time.Time)

	minDate := time.Now().AddDate(0, 0, -2)
	maxDate := time.Now().AddDate(0, 0, 2)

	return date.After(minDate) && date.Before(maxDate)
}

func endDateValidatorResult(fl validator.FieldLevel) bool {
	date := fl.Field().Interface().(time.Time)

	minDate := time.Now().AddDate(0, 0, -2)
	maxDate := time.Now().AddDate(0, 0, 2)

	return date.After(minDate) && date.Before(maxDate)
}

func dateDiffValidatorResult(fl validator.FieldLevel) bool {
	endDate := fl.Field().Interface().(time.Time)
	startDate := fl.Parent().FieldByName("MatchStartDate").Interface().(time.Time)

	diff := endDate.Sub(startDate)
	minDiff := time.Hour * 3 // 3 hours

	return diff >= minDiff
}

func scoreValidatorResult(fl validator.FieldLevel) bool {
	score := fl.Field().Interface().(string)

	scoreInt, err := strconv.Atoi(score)
	if err != nil {
		return false
	}

	return scoreInt >= 0 && scoreInt <= 100
}

func (dto *CreateResultDto) Validate() error {
	var validate = validator.New()

	if err := validate.RegisterValidation("startDate", startDateValidatorResult); err != nil {
		return err
	}
	if err := validate.RegisterValidation("endDate", endDateValidatorResult); err != nil {
		return err
	}
	if err := validate.RegisterValidation("dateDiff", dateDiffValidatorResult); err != nil {
		return err
	}
	if err := validate.RegisterValidation("score", scoreValidatorResult); err != nil {
		return err
	}

	return resultValidationErrorHandler(validate.Struct(dto))
}

func (dto *UpdateResultDto) Validate() error {
	var validate = validator.New()

	if err := validate.RegisterValidation("startDate", startDateValidatorResult); err != nil {
		return err
	}
	if err := validate.RegisterValidation("endDate", endDateValidatorResult); err != nil {
		return err
	}
	if err := validate.RegisterValidation("dateDiff", dateDiffValidatorResult); err != nil {
		return err
	}
	if err := validate.RegisterValidation("score", scoreValidatorResult); err != nil {
		return err
	}

	return resultValidationErrorHandler(validate.Struct(dto))
}

func resultValidationErrorHandler(err error) error {
	if err == nil {
		return nil
	}

	var errorMessage string

	for _, err := range err.(validator.ValidationErrors) {
		field := err.StructField()
		tag := err.Tag()

		switch field {
		case "MatchStartDate":
			switch tag {
			case "required":
				errorMessage = "Match Start Date is required."
			case "startDate":
				errorMessage = "Match Start Date must be within 2 days from now."
			}

		case "MatchEndDate":
			switch tag {
			case "required":
				errorMessage = "Match End Date is required."
			case "endDate":
				errorMessage = "Match End Date must be within 2 days from now."
			case "dateDiff":
				errorMessage = "Match End Date must be at least 3 hours later than the Match Start Date."
			case "gtfield":
				errorMessage = "Match End Date must be later than the Match Start Date."
			}
		case "Score":
			switch tag {
			case "score":
				errorMessage = "Score must be between 0 and 100."
			}
		case "OpponentScore":
			switch tag {
			case "score":
				errorMessage = "Opponent Score must be between 0 and 100."
			}
		case "OpponentTeamID":
			switch tag {
			case "required":
				errorMessage = "Opponent Team ID is required."
			}
		}

		if errorMessage == "" {
			errorMessage = fmt.Sprintf("Validation failed on field %s.", field)
		}
	}

	return errors.New(errorMessage)
}
