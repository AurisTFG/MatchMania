package trackmaniaapi

import (
	"MatchManiaAPI/constants"
	"fmt"
	"time"
)

type CreateCompetitionDto struct {
	ClubId                int           `json:"clubId"`
	Name                  string        `json:"name"`
	Description           string        `json:"description"`
	ParticipantType       string        `json:"participantType"`
	RegistrationStartDate *time.Time    `json:"registrationStartDate"`
	RegistrationEndDate   *time.Time    `json:"registrationEndDate"`
	RulesURL              *string       `json:"rulesUrl"`
	MaxPlayers            int           `json:"maxPlayers"`
	AllowedZone           *string       `json:"allowedZone"`
	Rounds                []Round       `json:"rounds"`
	SpotStructure         SpotStructure `json:"spotStructure"`
}

type Round struct {
	Name                string     `json:"name"`
	StartDate           string     `json:"startDate"`
	EndDate             string     `json:"endDate"`
	LeaderboardType     string     `json:"leaderboardType"`
	TeamLeaderboardType string     `json:"teamLeaderboardType"`
	Qualifier           *Qualifier `json:"qualifier"`
	Config              Config     `json:"config"`
}

type Config struct {
	Name           string  `json:"name"`
	Title          *string `json:"title"`
	Script         string  `json:"script"`
	ScriptSettings any     `json:"scriptSettings"`
	MaxPlayers     int     `json:"maxPlayers"`
	MaxSpectators  int     `json:"maxSpectators"`
	Plugin         string  `json:"plugin"`
	PluginSettings any     `json:"pluginSettings"`
	// Password          string     `json:"password"`
	Maps []string `json:"maps"`
	// PlayerRestriction string     `json:"playerRestriction"`
	// VoteRatios *VoteRatios `json:"voteRatios"`
}

type VoteRatios struct {
	Ban                   int `json:"ban"`
	Kick                  int `json:"kick"`
	RestartMap            int `json:"restartMap"`
	NextMap               int `json:"nextMap"`
	JumpToMapIndex        int `json:"jumpToMapIndex"`
	JumpToMapIdent        int `json:"jumpToMapIdent"`
	SetNextMapIndex       int `json:"setNextMapIndex"`
	SetNextMapIdent       int `json:"setNextMapIdent"`
	AutoTeamBalance       int `json:"autoTeamBalance"`
	SetModeScriptSettings int `json:"setModeScriptSettings"`
}

type Qualifier struct {
	Name            string    `json:"name"`
	Position        int       `json:"position"`
	Id              int       `json:"id"`
	EndDate         time.Time `json:"endDate"`
	StartDate       time.Time `json:"startDate"`
	LeaderboardType string    `json:"leaderboardType"`
	Config          Config    `json:"config"`
}

type SpotStructure struct {
	Version int         `json:"version"`
	Rounds  []SpotRound `json:"rounds"`
}

type SpotRound struct {
	Name               string             `json:"name"`
	MatchGeneratorType string             `json:"matchGeneratorType"`
	MatchGeneratorData MatchGeneratorData `json:"matchGeneratorData"`
}

type MatchGeneratorData struct {
	Matches []MatchGeneratorMatches `json:"matches"`
}

type MatchGeneratorMatches struct {
	Spots []MatchGeneratorSpots `json:"spots"`
}

type MatchGeneratorSpots struct {
	Seed     int    `json:"seed"`
	SpotType string `json:"spotType"`
}

type TMWTPluginSettings struct {
	AdImageUrls              string  `json:"S_AdImageUrls"`
	MessageTimer             string  `json:"S_MessageTimer"`
	AutoStartMode            string  `json:"S_AutoStartMode"`
	AutoStartDelay           int     `json:"S_AutoStartDelay"`
	PickBanStartAuto         bool    `json:"S_PickBanStartAuto"`
	PickBanOrder             string  `json:"S_PickBanOrder"`
	UseAutoReady             bool    `json:"S_UseAutoReady"`
	EnableReadyManager       bool    `json:"S_EnableReadyManager"`
	PickBanUseGamepadVersion bool    `json:"S_PickBanUseGamepadVersion"`
	ReadyStartRatio          float64 `json:"S_ReadyStartRatio"`
	ReadyMinimumTeamSize     int     `json:"S_ReadyMinimumTeamSize"`
}

type TMWTScriptSettings struct {
	ChatTime              int     `json:"S_ChatTime"`
	ForceLapsNumber       int     `json:"S_ForceLapsNb"`
	RespawnBehavior       int     `json:"S_RespawnBehaviour"`
	WarmupDuration        int     `json:"S_WarmUpDuration"`
	WarmupNumber          int     `json:"S_WarmUpNb"`
	WarmupTimeout         int     `json:"S_WarmUpTimeout"`
	PickBanEnable         bool    `json:"S_PickAndBan_Enable"`
	TeamsURL              *string `json:"S_TeamsUrl,omitempty"`
	MatchPointsLimit      int     `json:"S_MatchPointsLimit"`
	MatchInfo             *string `json:"S_MatchInfo,omitempty"`
	MapPointsLimit        int     `json:"S_MapPointsLimit"`
	FinishTimeout         int     `json:"S_FinishTimeout"`
	LoadingScreenImageURL string  `json:"S_LoadingScreenImageUrl"`
	SponsorsURL           string  `json:"S_SponsorsUrl"`
	HeaderLogoURL         string  `json:"S_HeaderLogoUrl"`
	IntroBackgroundURL    string  `json:"S_IntroBackgroundUrl"`
	IntroLogoURL          string  `json:"S_IntroLogoUrl"`
	Sign2x3DefaultURL     string  `json:"S_Sign2x3DefaultUrl"`
	Sign16x9DefaultURL    string  `json:"S_Sign16x9DefaultUrl"`
	Sign64x10DefaultURL   string  `json:"S_Sign64x10DefaultUrl"`
	DisableMatchIntro     bool    `json:"S_DisableMatchIntro"`
	ForceRoadSpectatorsNb int     `json:"S_ForceRoadSpectatorsNb"`
	EnableDossardColor    bool    `json:"S_EnableDossardColor"`
	IsMatchmaking         bool    `json:"S_IsMatchmaking"`
	// PickBanStyle           `json:"S_PickAndBanStyle"`
	ApiURL                 string `json:"S_ApiUrl"`
	ApiCompetitionUid      string `json:"S_ApiCompetitionUid"`
	ApiAuthorizationHeader string `json:"S_ApiAuthorizationHeader"`
}

type PickBanStyle struct {
	Background   string `json:"Background"`
	TopLeftLogo  string `json:"TopLeftLogo"`
	TopRightLogo string `json:"TopRightLogo"`
	BottomLogo   string `json:"BottomLogo"`
}

func MakeCompetition(clubId int, matchNumber int, label string, trackUids []string) CreateCompetitionDto {
	const iso8601Format = "2006-01-02T15:04:05.000Z"

	return CreateCompetitionDto{
		ClubId:                clubId,
		Name:                  fmt.Sprintf("MatchMania - #%d", matchNumber),
		Description:           "Competition generated by MatchMania, a matchmaking service for Trackmania",
		ParticipantType:       constants.ParticipantTeam,
		RegistrationStartDate: nil,
		RegistrationEndDate:   nil,
		RulesURL:              nil,
		MaxPlayers:            10000,
		AllowedZone:           nil,
		Rounds: []Round{
			{
				Name:                label,
				StartDate:           time.Now().Add(1 * time.Second).UTC().Format(iso8601Format),
				EndDate:             time.Now().Add(2 * time.Hour).UTC().Format(iso8601Format),
				LeaderboardType:     constants.LeaderboardTypeBracket,
				TeamLeaderboardType: constants.TeamLeaderboardTypeScore,
				Qualifier:           nil,
				Config: Config{
					Name:          label,
					Title:         nil,
					MaxPlayers:    4,
					MaxSpectators: 4,
					Maps:          trackUids,
					Script:        constants.ScriptTMWT2025,
					ScriptSettings: TMWTScriptSettings{
						ChatTime:               10,
						ForceLapsNumber:        -1,
						RespawnBehavior:        constants.RespawnNeverGiveUp,
						WarmupDuration:         0,
						WarmupNumber:           0,
						WarmupTimeout:          -1,
						PickBanEnable:          true,
						TeamsURL:               nil,
						MatchPointsLimit:       2,
						MatchInfo:              nil,
						MapPointsLimit:         10,
						FinishTimeout:          -1,
						LoadingScreenImageURL:  "",
						SponsorsURL:            "",
						HeaderLogoURL:          "",
						IntroBackgroundURL:     "",
						IntroLogoURL:           "",
						Sign2x3DefaultURL:      "file://Media/Manialinks/Nadeo/Trackmania/Modes/TMWT/Sign2x3/Default.dds",
						Sign16x9DefaultURL:     "",
						Sign64x10DefaultURL:    "file://Media/Manialinks/Nadeo/Trackmania/Modes/TMWT/Sign64x10/Default.dds",
						DisableMatchIntro:      false,
						ForceRoadSpectatorsNb:  -1,
						EnableDossardColor:     true,
						IsMatchmaking:          true,
						ApiURL:                 "",
						ApiCompetitionUid:      "",
						ApiAuthorizationHeader: "",
					},
					Plugin: constants.PluginClub,
					PluginSettings: TMWTPluginSettings{
						AdImageUrls:              "",
						MessageTimer:             "",
						AutoStartMode:            constants.AutoStartDelay,
						AutoStartDelay:           600,
						PickBanStartAuto:         true,
						PickBanOrder:             "b:0,b:1,p:0,p:1,p:0,p:1,p:0,p:1,p:r",
						UseAutoReady:             true,
						EnableReadyManager:       true,
						PickBanUseGamepadVersion: true,
						ReadyStartRatio:          1.0,
						ReadyMinimumTeamSize:     1,
					},
				},
			},
		},
		SpotStructure: SpotStructure{
			Version: 1,
			Rounds: []SpotRound{
				{
					Name:               label,
					MatchGeneratorType: constants.MatchGeneratorType,
					MatchGeneratorData: MatchGeneratorData{
						Matches: []MatchGeneratorMatches{
							{
								Spots: []MatchGeneratorSpots{
									{
										Seed:     1,
										SpotType: constants.SpotTeam,
									},
									{
										Seed:     2,
										SpotType: constants.SpotTeam,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
