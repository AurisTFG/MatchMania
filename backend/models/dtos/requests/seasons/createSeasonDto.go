package seasons

import "time"

type CreateSeasonDto struct {
	Name      string    `example:"Summer 2025"           json:"name"      validate:"required,min=3,max=100"`
	StartDate time.Time `example:"2025-06-01T00:00:000Z" json:"startDate" validate:"required,minDate=-30,maxDate=3650"`
	EndDate   time.Time `example:"2025-08-31T00:00:000Z" json:"endDate"   validate:"required,maxDate=3650,minDateDiff=7"`
}
