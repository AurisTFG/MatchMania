package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Result struct {
	BaseModel

	MatchStartDate time.Time `gorm:"not null"`
	MatchEndDate   time.Time `gorm:"not null"`
	Score          string    `gorm:"not null"`
	OpponentScore  string    `gorm:"not null"`

	Team         Team
	OpponentTeam Team

	SeasonId       uuid.UUID `gorm:"not null"`
	TeamId         uuid.UUID `gorm:"not null"`
	OpponentTeamId uuid.UUID `gorm:"not null"`
	UserId         uuid.UUID `gorm:"not null"`
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

// func (dto *CreateResultDto) Validate() error {
// 	var validate = validator.New()

// 	if err := validate.RegisterValidation("startDate", startDateValidatorResult); err != nil {
// 		return err
// 	}
// 	if err := validate.RegisterValidation("endDate", endDateValidatorResult); err != nil {
// 		return err
// 	}
// 	if err := validate.RegisterValidation("dateDiff", dateDiffValidatorResult); err != nil {
// 		return err
// 	}
// 	if err := validate.RegisterValidation("score", scoreValidatorResult); err != nil {
// 		return err
// 	}

// 	return resultValidationErrorHandler(validate.Struct(dto))
// }

// func (dto *UpdateResultDto) Validate() error {
// 	var validate = validator.New()

// 	if err := validate.RegisterValidation("startDate", startDateValidatorResult); err != nil {
// 		return err
// 	}
// 	if err := validate.RegisterValidation("endDate", endDateValidatorResult); err != nil {
// 		return err
// 	}
// 	if err := validate.RegisterValidation("dateDiff", dateDiffValidatorResult); err != nil {
// 		return err
// 	}
// 	if err := validate.RegisterValidation("score", scoreValidatorResult); err != nil {
// 		return err
// 	}

// 	return resultValidationErrorHandler(validate.Struct(dto))
// }

func resultValidationErrorHandler(err error) error {
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
