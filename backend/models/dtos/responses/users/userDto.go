package users

import (
	"MatchManiaAPI/models/dtos/responses/trackmanioauth"

	"github.com/google/uuid"
)

type UserDto struct {
	Id                    uuid.UUID                                 `example:"526432ea-822b-4b5b-b1b3-34e8ab453e03" json:"id"`
	Username              string                                    `example:"john_doe_123"                         json:"username"`
	Email                 string                                    `example:"email@example.com"                    json:"email"`
	Country               string                                    `example:"FR"                                   json:"country"`
	Permissions           []string                                  `                                               json:"permissions"`
	TrackmaniaId          uuid.UUID                                 `example:"526432ea-822b-4b5b-b1b3-34e8ab453e03" json:"trackmaniaId"`
	TrackmaniaName        string                                    `example:"JohnDoe"                              json:"trackmaniaName"`
	TrackmaniaOauthTracks []trackmanioauth.TrackmaniaOAuthTracksDto `                                               json:"tracks"`
}
