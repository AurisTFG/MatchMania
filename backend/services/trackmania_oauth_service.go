package services

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/constants"
	"MatchManiaAPI/models"
	"MatchManiaAPI/models/dtos/responses"
	"MatchManiaAPI/repositories"
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type TrackmaniaOAuthService interface {
	GenerateRandomState() string
	SaveState(state string, userId uuid.UUID) error
	GetAuthorizationUrl(state string) string
	VerifyCallbackResponse(state string, code string) bool
	GetUserIdByState(state string) (uuid.UUID, error)
	GetAccessToken(code string) (*responses.TrackmaniaOAuthAccessTokenDto, error)
	GetProfilePageUrl() string
	GetUserInfo(accessToken string) (*responses.TrackmaniaOAuthUserDto, error)
	GetUserFavoriteMaps(accessToken string) ([]responses.TrackmaniaOAuthFavoritesDto, error)
}

type trackmaniaOAuthService struct {
	trackmaniaOAuthStateRepository repositories.TrackmaniaOAuthStateRepository
	env                            *config.Env
}

func NewTrackmaniaOAuthService(
	trackmaniaOAuthStateRepository repositories.TrackmaniaOAuthStateRepository,
	env *config.Env,
) TrackmaniaOAuthService {
	return &trackmaniaOAuthService{
		trackmaniaOAuthStateRepository: trackmaniaOAuthStateRepository,
		env:                            env,
	}
}

func (s *trackmaniaOAuthService) GenerateRandomState() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 32)

	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			panic(fmt.Sprintf("failed to generate random number: %v", err))
		}
		b[i] = letters[n.Int64()]
	}

	return string(b)
}

func (s *trackmaniaOAuthService) SaveState(state string, userId uuid.UUID) error {
	err := s.trackmaniaOAuthStateRepository.DeleteStateByUserId(userId)
	if err != nil {
		return fmt.Errorf("failed to delete existing state: %w", err)
	}

	oauthState := &models.TrackmaniaOauthState{
		State:  state,
		UserId: userId,
	}

	err = s.trackmaniaOAuthStateRepository.SaveState(oauthState)
	if err != nil {
		return fmt.Errorf("failed to save state: %w", err)
	}

	return nil
}

func (s *trackmaniaOAuthService) GetAuthorizationUrl(state string) string {
	query := url.Values{}
	query.Set("response_type", "code")
	query.Set("client_id", s.env.TrackmaniaOAuthClientID)
	query.Set("state", state)
	query.Set("scope", s.env.TrackmaniaOAuthScopes)
	query.Set("redirect_uri", s.env.TrackmaniaOAuthRedirectURL)

	baseUrl := constants.TrackmaniaOAuthURL
	var url = fmt.Sprintf("%s?%s", baseUrl, query.Encode())

	return url
}

func (s *trackmaniaOAuthService) VerifyCallbackResponse(state string, code string) bool {
	if state == "" || code == "" {
		return false
	}

	exists := s.trackmaniaOAuthStateRepository.DoesStateExist(state)

	return exists
}

func (s *trackmaniaOAuthService) GetUserIdByState(state string) (uuid.UUID, error) {
	userId, err := s.trackmaniaOAuthStateRepository.GetUserIdByState(state)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to get user ID by state: %w", err)
	}

	return userId, nil
}

func (s *trackmaniaOAuthService) GetAccessToken(code string) (*responses.TrackmaniaOAuthAccessTokenDto, error) {
	var baseUrl = constants.TrackmaniaOAuthTokenURL

	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("client_id", s.env.TrackmaniaOAuthClientID)
	form.Set("client_secret", s.env.TrackmaniaOAuthClientSecret)
	form.Set("redirect_uri", s.env.TrackmaniaOAuthRedirectURL)
	form.Set("code", code)

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		baseUrl,
		bytes.NewBufferString(form.Encode()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			fmt.Printf("warning: failed to close response body: %v\n", cerr)
		}
	}()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get access token: %s", response.Status)
	}

	var tokenDto responses.TrackmaniaOAuthAccessTokenDto
	err = json.Unmarshal(body, &tokenDto)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if tokenDto.TokenType != "Bearer" {
		return nil, fmt.Errorf("invalid token type: %s", tokenDto.TokenType)
	}

	return &tokenDto, nil
}

func (s *trackmaniaOAuthService) GetProfilePageUrl() string {
	return fmt.Sprintf("%s/profile", s.env.ClientURL)
}

func (s *trackmaniaOAuthService) GetUserInfo(accessToken string) (*responses.TrackmaniaOAuthUserDto, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, constants.TrackmaniaOAuthUserURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			fmt.Printf("warning: failed to close response body: %v\n", cerr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var user responses.TrackmaniaOAuthUserDto
	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *trackmaniaOAuthService) GetUserFavoriteMaps(
	accessToken string,
) ([]responses.TrackmaniaOAuthFavoritesDto, error) {
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		constants.TrackmaniaOAuthFavoritesURL,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			fmt.Printf("warning: failed to close response body: %v\n", cerr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var favoritesWrapper struct {
		List []responses.TrackmaniaOAuthFavoritesDto `json:"list"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&favoritesWrapper); err != nil {
		return nil, err
	}

	return favoritesWrapper.List, nil
}
