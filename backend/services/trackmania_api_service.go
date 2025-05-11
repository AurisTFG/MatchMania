package services

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/constants"
	"MatchManiaAPI/models"
	"MatchManiaAPI/models/dtos/requests/trackmaniaapi"
	"MatchManiaAPI/models/dtos/responses"
	"MatchManiaAPI/models/enums"
	"MatchManiaAPI/utils"
	"context"
	"fmt"
	"net/http"
)

type TrackmaniaApiService interface {
	CreateCompetition(label string, trackUids []string) (*responses.CompetitionCreateResponseDto, error)
	DeleteCompetition(competitionId int) error
	AddTeamsToCompetition(*models.Team, *models.Team, int) error
	GetTeamsResults(competitionId int) (*responses.TeamsResultsDto, error)
}

type trackmaniaApiService struct {
	env               *config.Env
	ubisoftApiService UbisoftApiService
	nadeoApiService   NadeoApiService
	appSettingService AppSettingService

	client      *http.Client
	accessToken string
}

func NewTrackmaniaApiService(
	env *config.Env,
	ubisoftApiService UbisoftApiService,
	nadeoApiService NadeoApiService,
	appSettingService AppSettingService,
) TrackmaniaApiService {
	return &trackmaniaApiService{
		env:               env,
		ubisoftApiService: ubisoftApiService,
		nadeoApiService:   nadeoApiService,
		appSettingService: appSettingService,
		client:            &http.Client{},
		accessToken:       "",
	}
}

func (s *trackmaniaApiService) CreateCompetition(
	label string,
	trackUids []string,
) (*responses.CompetitionCreateResponseDto, error) {
	if err := s.authenticate(); err != nil {
		return nil, fmt.Errorf("authenticating: %w", err)
	}

	matchCount := 1
	lastMatchCount, err := GetSettingValue[int](s.appSettingService, enums.AppSettingMatchCount)
	if err == nil {
		matchCount = *lastMatchCount + 1
	}

	if err = s.appSettingService.Set(enums.AppSettingMatchCount, matchCount); err != nil {
		return nil, fmt.Errorf("saving match count: %w", err)
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		constants.TrackmaniaApiCreateCompetitionURL,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", utils.GetNadeoAuthHeader(s.accessToken))
	req.Header.Set("Content-Type", constants.ContentType)

	createCompetitionDto := trackmaniaapi.MakeCompetition(
		s.env.TrackmaniaApiClubId,
		matchCount,
		label,
		trackUids,
	)

	responseDto, err := utils.HttpRequest[responses.CompetitionCreateResponseDto](s.client, req, createCompetitionDto)
	if err != nil {
		return nil, fmt.Errorf("creating competition: %w", err)
	}

	return responseDto, nil
}

func (s *trackmaniaApiService) DeleteCompetition(competitionId int) error {
	if err := s.authenticate(); err != nil {
		return fmt.Errorf("authenticating: %w", err)
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		fmt.Sprintf(constants.TrackmaniaApiDeleteCompetitionURL, competitionId),
		nil,
	)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", utils.GetNadeoAuthHeader(s.accessToken))
	req.Header.Set("Content-Type", constants.ContentType)

	_, err = utils.HttpRequest[any](s.client, req, nil)
	if err != nil {
		return fmt.Errorf("deleting competition: %w", err)
	}

	return nil
}

func (s *trackmaniaApiService) AddTeamsToCompetition(teamA *models.Team, teamB *models.Team, competitionId int) error {
	if err := s.authenticate(); err != nil {
		return fmt.Errorf("authenticating: %w", err)
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		fmt.Sprintf(constants.TrackmaniaApiAddTeamsURL, competitionId),
		nil,
	)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", utils.GetNadeoAuthHeader(s.accessToken))
	req.Header.Set("Content-Type", constants.ContentType)

	addTeamsDto := trackmaniaapi.MakeCompetitionTeams(
		teamA,
		teamB,
	)

	for _, addTeamDto := range addTeamsDto {
		_, err = utils.HttpRequest[any](s.client, req, addTeamDto)
		if err != nil {
			return fmt.Errorf("adding teams: %w", err)
		}
	}

	return nil
}

func (s *trackmaniaApiService) GetTeamsResults(competitionId int) (*responses.TeamsResultsDto, error) {
	if err := s.authenticate(); err != nil {
		return nil, fmt.Errorf("authenticating: %w", err)
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		fmt.Sprintf(constants.TrackmaniaApiGetTeamsLeaderboardURL, competitionId),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", utils.GetNadeoAuthHeader(s.accessToken))
	req.Header.Set("Content-Type", constants.ContentType)

	responseDto, err := utils.HttpRequest[responses.TeamsResultsDto](s.client, req, nil)
	if err != nil {
		return nil, fmt.Errorf("getting teams leaderboard: %w", err)
	}

	return responseDto, nil
}

func (s *trackmaniaApiService) authenticate() error {
	session, err := s.ubisoftApiService.GetSession()
	if err != nil {
		return fmt.Errorf("getting session: %w", err)
	}

	authDto, err := s.nadeoApiService.GetAccessToken(session.Ticket)
	if err != nil {
		return fmt.Errorf("getting access token: %w", err)
	}

	s.accessToken = authDto.AccessToken
	return nil
}
