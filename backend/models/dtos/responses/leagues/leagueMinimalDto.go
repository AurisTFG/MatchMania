package leagues

import (
	"github.com/google/uuid"
)

type LeagueMinimalDto struct {
	Id      uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000" json:"id"`
	Name    string    `example:"Summer 2025"                          json:"name"`
	LogoUrl string    `example:"https://example.com/logo.png"         json:"logoUrl"`
}
