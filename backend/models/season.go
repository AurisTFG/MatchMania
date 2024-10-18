package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Season struct {
	gorm.Model
	Name      string    `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`

	Teams []Team `gorm:"foreignKey:SeasonID"`
}

type SeasonDto struct {
	ID        uint      `json:"id" example:"5"`
	Name      string    `json:"name" example:"Summer 2025"`
	StartDate time.Time `json:"startDate" example:"2025-06-01T00:00:00Z"`
	EndDate   time.Time `json:"endDate" example:"2025-08-31T00:00:00Z"`
}

type CreateSeasonDto struct {
	Name      string    `json:"name" example:"Summer 2025" validate:"required,min=3,max=100"`
	StartDate time.Time `json:"startDate" example:"2025-06-01T00:00:00Z" validate:"required,startDate"`
	EndDate   time.Time `json:"endDate" example:"2025-08-31T00:00:00Z" validate:"required,endDate,dateDiff,gtfield=StartDate"`
}

type UpdateSeasonDto struct {
	Name      string    `json:"name" example:"Summer 2025" validate:"required,min=3,max=100"`
	StartDate time.Time `json:"startDate" example:"2025-06-01T00:00:00Z" validate:"required,startDate"`
	EndDate   time.Time `json:"endDate" example:"2025-08-31T00:00:00Z" validate:"required,endDate,dateDiff,gtfield=StartDate"`
}

func (s *Season) ToDto() SeasonDto {
	return SeasonDto{
		ID:        s.ID,
		Name:      s.Name,
		StartDate: s.StartDate,
		EndDate:   s.EndDate,
	}
}

func ToSeasonDtos(seasons []Season) []SeasonDto {
	seasonDtos := make([]SeasonDto, len(seasons))

	for i, season := range seasons {
		seasonDtos[i] = season.ToDto()
	}

	return seasonDtos
}

func (s *CreateSeasonDto) ToSeason() Season {
	return Season{
		Name:      s.Name,
		StartDate: s.StartDate,
		EndDate:   s.EndDate,
	}
}

func (s *UpdateSeasonDto) ToSeason() Season {
	return Season{
		Name:      s.Name,
		StartDate: s.StartDate,
		EndDate:   s.EndDate,
	}
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

func (dto *CreateSeasonDto) Validate() error {
	var validate = validator.New()

	validate.RegisterValidation("startDate", startDateValidator)
	validate.RegisterValidation("endDate", endDateValidator)
	validate.RegisterValidation("dateDiff", dateDiffValidator)

	return seasonValidationErrorHandler(validate.Struct(dto))
}

func (dto *UpdateSeasonDto) Validate() error {
	var validate = validator.New()

	validate.RegisterValidation("startDate", startDateValidator)
	validate.RegisterValidation("endDate", endDateValidator)
	validate.RegisterValidation("dateDiff", dateDiffValidator)

	return seasonValidationErrorHandler(validate.Struct(dto))
}

func seasonValidationErrorHandler(err error) error {
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
