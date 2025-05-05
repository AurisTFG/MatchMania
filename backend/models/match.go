package models

type Match struct {
	BaseModel

	MatchId string `gorm:"unique;not null"`
}
