package models

import "github.com/google/uuid"

type QueueTeam struct {
	BaseModel

	TeamId  uuid.UUID `gorm:"not null"`
	QueueId uuid.UUID `gorm:"not null"`

	Team  *Team
	Queue Queue
}
