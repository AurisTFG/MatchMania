package responses

import (
	"time"

	"github.com/google/uuid"
)

type SeasonDto struct {
	Id        uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000" json:"id"`
	Name      string    `example:"Summer 2025"          json:"name"`
	StartDate time.Time `example:"2025-06-01T00:00:00Z" json:"startDate"`
	EndDate   time.Time `example:"2025-08-31T00:00:00Z" json:"endDate"`
	UserId    uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000" json:"userId"`
}
