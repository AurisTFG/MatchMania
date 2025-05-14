package leagues

import "time"

type CreateLeagueDto struct {
	Name      string    `example:"Summer 2025"                          json:"name"      validate:"required,min=3,max=100"`
	LogoUrl   *string   `example:"https://example.com/logo.png"         json:"logoUrl"   validate:"omitnil,url,max=255"`
	StartDate time.Time `example:"2025-06-01T00:00:000Z"                json:"startDate" validate:"required,minDate=-30,maxDate=3650"`
	EndDate   time.Time `example:"2025-08-31T00:00:000Z"                json:"endDate"   validate:"required,maxDate=3650,minDateDiff=7"`

	TrackIds []string `example:"550e8400-e29b-41d4-a716-446655440000" json:"trackIds"  validate:"required,min=1,max=100,dive,uuid"`
}
