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
	CreateRefreshToken(sessionUUID string, user *models.User) (string, error)
	VerifyAccessToken(token string) (*models.User, error)
	VerifyRefreshToken(token string) (user *models.User, sessionID string, err error)
	CreateSession(sessionUUID uuid.UUID, userUUID uuid.UUID, refreshToken string) error
	ExtendSession(sessionUUID string, refreshToken string) error
	InvalidateSession(sessionUUID string) error
	IsSessionValid(sessionUUID string, refreshToken string) bool
	SetCookie(ctx *gin.Context, name string, value string)
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
		"sub":  user.UUID,
		"role": user.Role,
		"iss":  s.env.JWTIssuer,
		"aud":  s.env.JWTAudience,
		"iat":  time.Now().Unix(),
		"nbf":  time.Now().Unix(),
		"exp":  time.Now().Add(s.env.JWTAccessTokenDuration).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(s.env.JWTAccessTokenSecret))

	return accessTokenString, err
}

func (s *authService) CreateRefreshToken(sessionUUID string, user *models.User) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sessionId": sessionUUID,
		"sub":       user.UUID,
		"role":      user.Role,
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

	user, err := s.userRepo.FindByID(claims["sub"].(string))
	if err != nil {
		return nil, errors.New("invalid user")
	}

	roleClaim := models.Role(claims["role"].(string))
	if user.Role != roleClaim {
		return nil, errors.New("user role mismatch")
	}

	return user, nil
}

func (s *authService) VerifyRefreshToken(refreshToken string) (*models.User, string, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.env.JWTRefreshTokenSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, "", errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, "", errors.New("invalid refresh token claims")
	}

	if claims["jti"] == "" {
		return nil, "", errors.New("invalid session id")
	}

	if claims["aud"] != s.env.JWTAudience {
		return nil, "", errors.New("invalid audience")
	}

	if claims["iss"] != s.env.JWTIssuer {
		return nil, "", errors.New("invalid issuer")
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, "", errors.New("refresh token expired")
	}

	if claims["nbf"].(float64) > float64(time.Now().Unix()) {
		return nil, "", errors.New("refresh token not valid yet")
	}

	if claims["exp"].(float64)-claims["iat"].(float64) != s.env.JWTRefreshTokenDuration.Seconds() {
		return nil, "", errors.New("refresh token expiration date is invalid")
	}

	user, err := s.userRepo.FindByID(claims["sub"].(string))
	if err != nil {
		return nil, "", errors.New("invalid user")
	}

	roleClaim := models.Role(claims["role"].(string))
	if user.Role != roleClaim {
		return nil, "", errors.New("user role mismatch")
	}

	return user, claims["sessionId"].(string), nil
}

func (s *authService) CreateSession(sessionUUID uuid.UUID, userUUID uuid.UUID, refreshToken string) error {
	session := &models.Session{
		UUID:             sessionUUID,
		UserUUID:         userUUID,
		LastRefreshToken: refreshToken,
		ExpiresAt:        time.Now().Add(s.env.JWTRefreshTokenDuration),
		InitiatedAt:      time.Now(),
		IsRevoked:        false,
	}

	if err := session.HashToken(); err != nil {
		return err
	}

	_, err := s.sessionRepo.Create(session)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) ExtendSession(sessionUUID string, refreshToken string) error {
	session, err := s.sessionRepo.FindByID(sessionUUID)
	if err != nil {
		return err
	}

	session.LastRefreshToken = refreshToken
	session.ExpiresAt = time.Now().Add(s.env.JWTRefreshTokenDuration)

	if err = session.HashToken(); err != nil {
		return err
	}

	_, err = s.sessionRepo.Update(session)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) InvalidateSession(sessionUUID string) error {
	session, err := s.sessionRepo.FindByID(sessionUUID)
	if err != nil {
		return err
	}

	session.IsRevoked = true

	_, err = s.sessionRepo.Update(session)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) IsSessionValid(sessionUUID string, refreshToken string) bool {
	session, err := s.sessionRepo.FindByID(sessionUUID)
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

func (s *authService) SetCookie(ctx *gin.Context, name string, value string) {
	maxAge := -1
	path := ""
	domain := ""
	secure := false
	httpOnly := true

	if value != "" {
		switch name {
		case constants.AccessTokenName:
			maxAge = int(s.env.JWTAccessTokenDuration.Seconds())
			path = "/"
		case constants.RefreshTokenName:
			maxAge = int(s.env.JWTRefreshTokenDuration.Seconds())
			path = "/refresh"
		default:
			panic("Invalid cookie name")
		}

		switch {
		case s.env.IsDev:
			domain = "localhost"
			secure = false
			ctx.SetSameSite(http.SameSiteLaxMode)
		case s.env.IsProd:
			domain = s.env.ClientURL
			secure = true
			ctx.SetSameSite(http.SameSiteNoneMode)
		default:
			panic("Invalid environment")
		}
	}

	ctx.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
}
