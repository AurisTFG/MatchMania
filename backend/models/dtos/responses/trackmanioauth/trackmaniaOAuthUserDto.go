package trackmanioauth

import "github.com/google/uuid"

type TrackmaniaOAuthUserDto struct {
	AccountId   uuid.UUID `json:"accountId"`
	DisplayName string    `json:"displayName"`
}
