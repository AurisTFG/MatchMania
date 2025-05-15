package teams

type CreateTeamDto struct {
	Name    string  `example:"BIG Clan"                     json:"name"    validate:"required,min=3,max=100"`
	LogoUrl *string `example:"https://example.com/logo.png" json:"logoUrl" validate:"omitnil,url,max=255"`

	LeagueIds []string `example:"550e8400-e29b-41d4-a716-446655440000" json:"leagueIds" validate:"required,min=1,max=10,dive,uuid"`
	PlayerIds []string `example:"550e8400-e29b-41d4-a716-446655440000" json:"playerIds" validate:"required,min=1,max=10,dive,uuid"`
}
