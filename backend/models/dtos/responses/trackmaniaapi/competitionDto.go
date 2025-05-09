package trackmaniaapi

type CompetitionCreateResponseDto struct {
	Id                      int         `json:"id"`
	ClubId                  int         `json:"clubId"`
	ActivityId              int         `json:"activityId"`
	MaxPlayers              int         `json:"maxPlayers"`
	Type                    string      `json:"type"`
	Competition             Competition `json:"competition"`
	ExternalRegistrationUrl *string     `json:"externalRegistrationUrl"`
}

type Competition struct {
	Id                           int      `json:"id"`
	LiveId                       string   `json:"liveId"`
	Creator                      string   `json:"creator"`
	Name                         string   `json:"name"`
	ParticipantType              string   `json:"participantType"`
	Description                  string   `json:"description"`
	RegistrationStart            *int64   `json:"registrationStart"`
	RegistrationEnd              *int64   `json:"registrationEnd"`
	StartDate                    int64    `json:"startDate"`
	EndDate                      int64    `json:"endDate"`
	MatchesGenerationDate        int64    `json:"matchesGenerationDate"`
	NbPlayers                    int      `json:"nbPlayers"`
	SpotStructure                string   `json:"spotStructure"`
	LeaderboardID                int      `json:"leaderboardId"`
	Manialink                    *string  `json:"manialink"`
	RulesUrl                     *string  `json:"rulesUrl"`
	StreamUrl                    *string  `json:"streamUrl"`
	WebsiteUrl                   *string  `json:"websiteUrl"`
	LogoUrl                      *string  `json:"logoUrl"`
	VerticalUrl                  *string  `json:"verticalUrl"`
	AllowedZones                 []string `json:"allowedZones"`
	DeletedOn                    *int64   `json:"deletedOn"`
	AutoNormalizeSeeds           bool     `json:"autoNormalizeSeeds"`
	Region                       *string  `json:"region"`
	AutoGetParticipantSkillLevel string   `json:"autoGetParticipantSkillLevel"`
	MatchAutoMode                string   `json:"matchAutoMode"`
	Partition                    string   `json:"partition"`
	ClubId                       int      `json:"clubId"`
	ActivityId                   int      `json:"activityId"`
}
