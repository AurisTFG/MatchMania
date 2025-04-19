package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Season struct {
	BaseModel

	Name      string    `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`

	UserId uuid.UUID `gorm:"not null"`
	User   User

	Teams []Team `gorm:"constraint:OnDelete:CASCADE;"`
}

func (s *Season) BeforeDelete(tx *gorm.DB) error {
	results := []Result{}
	tx.Where("season_id = ?", s.Id).Find(&results)

	for _, result := range results {
		if err := tx.Delete(&result).Error; err != nil {
			return err
		}
	}

	teams := []Team{}
	tx.Where("season_id = ?", s.Id).Find(&teams)

	for _, team := range teams {
		if err := tx.Delete(&team).Error; err != nil {
			return err
		}
	}

	return nil
}

func startDateValidator(fl validator.FieldLevel) bool {
	date := fl.Field().Interface().(time.Time)
	today := time.Now()
	yearAgo := today.AddDate(-1, 0, 0)
	tenYearsFromNow := today.AddDate(10, 0, 0)

	return date.After(yearAgo) && date.Before(tenYearsFromNow)
}

func endDateValidator(fl validator.FieldLevel) bool {
	date := fl.Field().Interface().(time.Time)
	today := time.Now()
	tenYearsFromNow := today.AddDate(10, 0, 0)

	return date.After(today) && date.Before(tenYearsFromNow)
}

func dateDiffValidator(fl validator.FieldLevel) bool {
	endDate := fl.Field().Interface().(time.Time)
	startDate := fl.Parent().FieldByName("StartDate").Interface().(time.Time)

	diff := endDate.Sub(startDate)
	minMonths := time.Hour * 24 * 30 * 1 // 1 month

	return diff >= minMonths
}

// func (dto *CreateSeasonDto) Validate() error {
// 	var validate = validator.New()

// 	if err := validate.RegisterValidation("startDate", startDateValidator); err != nil {
// 		return err
// 	}
// 	if err := validate.RegisterValidation("endDate", endDateValidator); err != nil {
// 		return err
// 	}
// 	if err := validate.RegisterValidation("dateDiff", dateDiffValidator); err != nil {
// 		return err
// 	}

// 	return seasonValidationErrorHandler(validate.Struct(dto))
// }

// func (dto *UpdateSeasonDto) Validate() error {
// 	var validate = validator.New()

// 	if err := validate.RegisterValidation("startDate", startDateValidator); err != nil {
// 		return err
// 	}
// 	if err := validate.RegisterValidation("endDate", endDateValidator); err != nil {
// 		return err
// 	}
// 	if err := validate.RegisterValidation("dateDiff", dateDiffValidator); err != nil {
// 		return err
// 	}

// 	return seasonValidationErrorHandler(validate.Struct(dto))
// }

func seasonValidationErrorHandler(err error) error {
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

		case "StartDate":
			switch tag {
			case "required":
				errorMessage = "Start Date is required."
			case "startDate":
				errorMessage = "Start Date must be a valid date between 1 year ago and 10 years from now."
			}

		case "EndDate":
			switch tag {
			case "required":
				errorMessage = "End Date is required."
			case "endDate":
				errorMessage = "End Date must be a valid date between today and 10 years from now."
			case "gtfield":
				errorMessage = "End Date must be later than the Start Date."
			case "dateDiff":
				errorMessage = "The difference between the Start Date and End Date must be at least 1 month."
			}
		}

		if errorMessage == "" {
			errorMessage = fmt.Sprintf("Validation failed on field %s.", field)
		}
	}

	return errors.New(errorMessage)
}
