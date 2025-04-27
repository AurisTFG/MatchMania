package services

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/constants"
	"MatchManiaAPI/models"
	"MatchManiaAPI/repositories"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService interface {
	CreateAccessToken(user *models.User) (string, error)
	CreateRefreshToken(sessionId uuid.UUID, user *models.User) (string, error)
	VerifyAccessToken(token string) (*models.User, error)
	VerifyRefreshToken(token string) (user *models.User, sessionID uuid.UUID, err error)
	CreateSession(sessionId uuid.UUID, UserId uuid.UUID, refreshToken string) error
	ExtendSession(sessionId uuid.UUID, refreshToken string) error
	InvalidateSession(sessionId uuid.UUID) error
	IsSessionValid(sessionId uuid.UUID, refreshToken string) bool
	SetCookies(ctx *gin.Context, accessToken string, refreshToken string)
	DeleteCookies(ctx *gin.Context)
}

type authService struct {
	sessionRepo repositories.SessionRepository
	userRepo    repositories.UserRepository
	env         *config.Env
}

func NewAuthService(
	sessionRepo repositories.SessionRepository,
	userRepo repositories.UserRepository,
	env *config.Env,
) AuthService {
	return &authService{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
		env:         env,
	}
}

func (s *authService) CreateAccessToken(user *models.User) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"iss": s.env.JWTIssuer,
		"aud": s.env.JWTAudience,
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(s.env.JWTAccessTokenDuration).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(s.env.JWTAccessTokenSecret))

	return accessTokenString, err
}

func (s *authService) CreateRefreshToken(sessionId uuid.UUID, user *models.User) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sessionId": sessionId,
		"sub":       user.Id,
		"iss":       s.env.JWTIssuer,
		"aud":       s.env.JWTAudience,
		"iat":       time.Now().Unix(),
		"nbf":       time.Now().Unix(),
		"exp":       time.Now().Add(s.env.JWTRefreshTokenDuration).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(s.env.JWTRefreshTokenSecret))

	return refreshTokenString, err
}

func (s *authService) VerifyAccessToken(accessToken string) (*models.User, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.env.JWTAccessTokenSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid access token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid access token claims")
	}

	if claims["aud"] != s.env.JWTAudience {
		return nil, errors.New("invalid audience")
	}

	if claims["iss"] != s.env.JWTIssuer {
		return nil, errors.New("invalid issuer")
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, errors.New("access token expired")
	}

	if claims["nbf"].(float64) > float64(time.Now().Unix()) {
		return nil, errors.New("access token not valid yet")
	}

	if claims["exp"].(float64)-claims["iat"].(float64) != s.env.JWTAccessTokenDuration.Seconds() {
		return nil, errors.New("access token expiration date is invalid")
	}

	subStr, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("invalid sub claim type")
	}

	userId, err := uuid.Parse(subStr)
	if err != nil {
		return nil, errors.New("invalid sub claim value")
	}

	user, err := s.userRepo.FindById(userId)
	if err != nil {
		return nil, errors.New("invalid user")
	}

	return user, nil
}

func (s *authService) VerifyRefreshToken(refreshToken string) (*models.User, uuid.UUID, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.env.JWTRefreshTokenSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, uuid.Nil, errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, uuid.Nil, errors.New("invalid refresh token claims")
	}

	if claims["jti"] == "" {
		return nil, uuid.Nil, errors.New("invalid session id")
	}

	if claims["aud"] != s.env.JWTAudience {
		return nil, uuid.Nil, errors.New("invalid audience")
	}

	if claims["iss"] != s.env.JWTIssuer {
		return nil, uuid.Nil, errors.New("invalid issuer")
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, uuid.Nil, errors.New("refresh token expired")
	}

	if claims["nbf"].(float64) > float64(time.Now().Unix()) {
		return nil, uuid.Nil, errors.New("refresh token not valid yet")
	}

	if claims["exp"].(float64)-claims["iat"].(float64) != s.env.JWTRefreshTokenDuration.Seconds() {
		return nil, uuid.Nil, errors.New("refresh token expiration date is invalid")
	}

	subStr, ok := claims["sub"].(string)
	if !ok {
		return nil, uuid.Nil, errors.New("invalid sub claim type")
	}

	userId, err := uuid.Parse(subStr)
	if err != nil {
		return nil, uuid.Nil, errors.New("invalid sub claim value")
	}

	user, err := s.userRepo.FindById(userId)
	if err != nil {
		return nil, uuid.Nil, errors.New("invalid user")
	}

	sessionIdStr, ok := claims["sessionId"].(string)
	if !ok {
		return nil, uuid.Nil, errors.New("invalid sessionId claim type")
	}

	sessionId, err := uuid.Parse(sessionIdStr)
	if err != nil {
		return nil, uuid.Nil, errors.New("invalid sessionId claim value")
	}

	return user, sessionId, nil
}

func (s *authService) CreateSession(sessionId uuid.UUID, userId uuid.UUID, refreshToken string) error {
	session := &models.Session{
		BaseModel:        models.BaseModel{Id: sessionId},
		UserId:           userId,
		LastRefreshToken: refreshToken,
		ExpiresAt:        time.Now().Add(s.env.JWTRefreshTokenDuration),
		InitiatedAt:      time.Now(),
		IsRevoked:        false,
	}

	if err := session.HashToken(); err != nil {
		return err
	}

	if err := s.sessionRepo.Create(session); err != nil {
		return err
	}

	return nil
}

func (s *authService) ExtendSession(sessionId uuid.UUID, refreshToken string) error {
	session, err := s.sessionRepo.FindById(sessionId)
	if err != nil {
		return err
	}

	session.LastRefreshToken = refreshToken
	session.ExpiresAt = time.Now().Add(s.env.JWTRefreshTokenDuration)

	if err = session.HashToken(); err != nil {
		return err
	}

	if err = s.sessionRepo.Update(session); err != nil {
		return err
	}

	return nil
}

func (s *authService) InvalidateSession(sessionId uuid.UUID) error {
	session, err := s.sessionRepo.FindById(sessionId)
	if err != nil {
		return err
	}

	session.IsRevoked = true

	if err = s.sessionRepo.Update(session); err != nil {
		return err
	}

	return nil
}

func (s *authService) IsSessionValid(sessionId uuid.UUID, refreshToken string) bool {
	session, err := s.sessionRepo.FindById(sessionId)
	if err != nil {
		return false
	}

	isTokenValid := session.CompareToken(refreshToken)
	if !isTokenValid {
		return false
	}

	if session.IsRevoked {
		return false
	}

	if session.ExpiresAt.Before(time.Now()) {
		return false
	}

	return true
}

func (s *authService) SetCookies(ctx *gin.Context, accessToken string, refreshToken string) {
	ctx.SetSameSite(http.SameSiteNoneMode)

	ctx.SetCookie(
		constants.AccessTokenName, accessToken, int(s.env.JWTAccessTokenDuration.Seconds()),
		"/api/v1", s.env.ClientURL, true, true)
	ctx.SetCookie(
		constants.RefreshTokenName, refreshToken, int(s.env.JWTRefreshTokenDuration.Seconds()),
		"/api/v1/auth", s.env.ClientURL, true, true)
}

func (s *authService) DeleteCookies(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteNoneMode)

	ctx.SetCookie(constants.AccessTokenName, "", -1, "/api/v1", s.env.ClientURL, true, true)
	ctx.SetCookie(constants.RefreshTokenName, "", -1, "/api/v1/auth", s.env.ClientURL, true, true)
}
