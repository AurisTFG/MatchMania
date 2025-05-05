package models

type MatchTeam struct {
	BaseModel

	MatchId string `gorm:"not null"`
	TeamId  string `gorm:"not null"`

	Team  *Team
	Match Match
}
