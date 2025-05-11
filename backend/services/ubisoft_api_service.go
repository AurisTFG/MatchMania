package services

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/constants"
	"MatchManiaAPI/models/dtos/responses"
	"MatchManiaAPI/models/enums"
	"MatchManiaAPI/utils"
	"context"
	"fmt"
	"net/http"
	"time"
)

type UbisoftApiService interface {
	GetSession() (*responses.UbisoftSessionDto, error)
}

type ubisoftApiService struct {
	env               *config.Env
	appSettingService AppSettingService

	client            *http.Client
	session           *responses.UbisoftSessionDto
	sessionExpireDate time.Time
}

func NewUbisoftApiService(
	env *config.Env,
	appSettingService AppSettingService,
) UbisoftApiService {
	return &ubisoftApiService{
		env:               env,
		appSettingService: appSettingService,
		client:            &http.Client{},
		session:           nil,
		sessionExpireDate: time.Time{},
	}
}

func (s *ubisoftApiService) GetSession() (*responses.UbisoftSessionDto, error) {
	if s.session == nil {
		if err := s.applyTokenFromDatabase(); err != nil {
			return nil, fmt.Errorf("applying token from database: %w", err)
		}

		fmt.Println("Session from database:", s.session)
		fmt.Print("Session expiration date from database:", s.sessionExpireDate)
		fmt.Print("Time:", time.Now().UTC())
	}

	if s.session != nil && s.sessionExpireDate.After(time.Now().UTC().Add(10*time.Minute)) {
		return s.session, nil
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		constants.UbisoftApiSessionURL,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", utils.GetBasicAuthHeader(s.env.TrackmaniaApiEmail, s.env.TrackmaniaApiPassword))
	req.Header.Set("Ubi-Appid", constants.TrackmaniaAppId)
	req.Header.Set("User-Agent", constants.UserAgent)
	req.Header.Set("Content-Type", constants.ContentType)

	ubisoftSessionDto, err := utils.HttpRequest[responses.UbisoftSessionDto](s.client, req, nil)
	if err != nil {
		return nil, fmt.Errorf("getting session: %w", err)
	}

	sessionExpirationDate, err := getUbisoftSessionExpirationDate(ubisoftSessionDto)
	if err != nil {
		return nil, fmt.Errorf("getting session expiration date: %w", err)
	}

	if err = s.appSettingService.Set(enums.AppSettingUbisoftAuthResponse, ubisoftSessionDto); err != nil {
		return nil, fmt.Errorf("saving ubisoft auth response: %w", err)
	}

	s.session = ubisoftSessionDto
	s.sessionExpireDate = *sessionExpirationDate

	return s.session, nil
}

func getUbisoftSessionExpirationDate(ubisoftSessionDto *responses.UbisoftSessionDto) (*time.Time, error) {
	expiration, err := time.Parse(time.RFC3339, ubisoftSessionDto.Expiration)
	if err != nil {
		return nil, fmt.Errorf("parsing expiration time: %w", err)
	}

	expirationDate := expiration.UTC()

	return &expirationDate, nil
}

func (s *ubisoftApiService) applyTokenFromDatabase() error {
	ubisoftSessionDto, _ := GetSettingValue[responses.UbisoftSessionDto](
		s.appSettingService,
		enums.AppSettingUbisoftAuthResponse,
	)
	if ubisoftSessionDto == nil {
		return nil
	}

	sessionExpirationDate, err := getUbisoftSessionExpirationDate(ubisoftSessionDto)
	if err != nil {
		return err
	}

	s.session = ubisoftSessionDto
	s.sessionExpireDate = *sessionExpirationDate

	return nil
}
