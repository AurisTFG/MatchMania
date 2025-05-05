package players

import (
	"github.com/google/uuid"
)

type PlayerMinimalDto struct {
	Id              uuid.UUID `example:"526432ea-822b-4b5b-b1b3-34e8ab453e03" json:"id"`
	TrackmaniaName  string    `example:"JohnDoe"                              json:"trackmaniaName"`
	ProfilePhotoUrl string    `example:"https://example.com/profile.jpg"      json:"profilePhotoUrl"`
}
