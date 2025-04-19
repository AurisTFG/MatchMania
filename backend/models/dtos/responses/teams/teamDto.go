package responses

import "github.com/google/uuid"

type TeamDto struct {
	Id   uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000" json:"id"`
	Name string    `example:"BIG Clan" json:"name"`
	Elo  uint      `example:"1200"     json:"elo"`

	SeasonId uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000" json:"seasonId"`
	UserId   uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000" json:"userId"`
}
