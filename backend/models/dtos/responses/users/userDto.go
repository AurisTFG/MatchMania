package users

import (
	"MatchManiaAPI/models/dtos/responses/roles"
	"MatchManiaAPI/models/dtos/responses/trackmanioauth"

	"github.com/google/uuid"
)

type UserDto struct {
	Id               uuid.UUID                            `example:"526432ea-822b-4b5b-b1b3-34e8ab453e03" json:"id"`
	Username         string                               `example:"john_doe_123"                         json:"username"`
	Email            string                               `example:"email@example.com"                    json:"email"`
	ProfilePhotoUrl  string                               `example:"https://example.com/profile.jpg"      json:"profilePhotoUrl"`
	Country          string                               `example:"FR"                                   json:"country"`
	Roles            []roles.RoleDto                      `                                               json:"roles"`
	Permissions      []string                             `                                               json:"permissions"`
	TrackmaniaId     uuid.UUID                            `example:"526432ea-822b-4b5b-b1b3-34e8ab453e03" json:"trackmaniaId"`
	TrackmaniaName   string                               `example:"JohnDoe"                              json:"trackmaniaName"`
	TrackmaniaTracks []trackmanioauth.TrackmaniaTracksDto `                                               json:"tracks"`
}
