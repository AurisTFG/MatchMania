package services

import (
	"MatchManiaAPI/constants"
	"MatchManiaAPI/models/dtos/responses"
	"MatchManiaAPI/models/enums"
	"MatchManiaAPI/utils"
	"context"
	"fmt"
	"net/http"
	"time"
)

type NadeoApiService interface {
	GetAccessToken(ubisoftTicket string) (*responses.NadeoAuthDto, error)
}

type nadeoApiService struct {
	appSettingService AppSettingService

	client         *http.Client
	auth           *responses.NadeoAuthDto
	authExpireDate time.Time
}

func NewNadeoApiService(
	appSettingService AppSettingService,
) NadeoApiService {
	return &nadeoApiService{
		appSettingService: appSettingService,
		client:            &http.Client{},
		auth:              nil,
		authExpireDate:    time.Time{},
	}
}

func (s *nadeoApiService) GetAccessToken(ubisoftTicket string) (*responses.NadeoAuthDto, error) {
	if s.auth == nil {
		if err := s.applyTokenFromDatabase(); err != nil {
			return nil, fmt.Errorf("applying token from database: %w", err)
		}
	}

	if s.auth != nil && s.authExpireDate.After(time.Now().Add(10*time.Minute)) {
		return s.auth, nil
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		constants.NadeoApiAuthURL,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", utils.GetUbisoftAuthHeader(ubisoftTicket))
	req.Header.Set("Content-Type", constants.ContentType)

	body := map[string]string{"audience": constants.NadeoLiveServicesAudience}

	nadeoAuthDto, err := utils.HttpRequest[responses.NadeoAuthDto](s.client, req, body)
	if err != nil {
		return nil, fmt.Errorf("getting access token: %w", err)
	}

	expireDate, err := getNadeoSessionExpirationDate(nadeoAuthDto)
	if err != nil {
		return nil, fmt.Errorf("getting expiration date: %w", err)
	}

	if err = s.appSettingService.Set(enums.AppSettingNadeoAuthResponse, nadeoAuthDto); err != nil {
		return nil, fmt.Errorf("saving nadeo auth response: %w", err)
	}

	s.auth = nadeoAuthDto
	s.authExpireDate = *expireDate

	return s.auth, nil
}

func getNadeoSessionExpirationDate(nadeoAuthDto *responses.NadeoAuthDto) (*time.Time, error) {
	exp, err := utils.GetJwtClaim[float64](nadeoAuthDto.AccessToken, "exp")
	if err != nil {
		return nil, fmt.Errorf("extracting 'exp' from JWT: %w", err)
	}

	expirationDate := time.Unix(int64(exp), 0)

	return &expirationDate, nil
}

func (s *nadeoApiService) applyTokenFromDatabase() error {
	nadeoAuthDto, _ := GetSettingValue[responses.NadeoAuthDto](s.appSettingService, enums.AppSettingNadeoAuthResponse)
	if nadeoAuthDto == nil {
		return nil
	}

	sessionExpirationDate, err := getNadeoSessionExpirationDate(nadeoAuthDto)
	if err != nil {
		return err
	}

	s.auth = nadeoAuthDto
	s.authExpireDate = *sessionExpirationDate

	return nil
}
