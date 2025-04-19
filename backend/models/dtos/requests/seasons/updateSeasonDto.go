package requests

import "time"

type UpdateSeasonDto struct {
	Name      string    `example:"Summer 2025"          json:"name"      validate:"required,min=3,max=100"`
	StartDate time.Time `example:"2025-06-01T00:00:00Z" json:"startDate" validate:"required,startDate"`
	EndDate   time.Time `example:"2025-08-31T00:00:00Z" json:"endDate"   validate:"required,endDate,dateDiff,gtfield=StartDate"`
}
