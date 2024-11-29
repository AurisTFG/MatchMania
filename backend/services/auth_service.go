package services

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService interface {
	CreateAccessToken(user *models.User) (string, error)
	CreateRefreshToken(user *models.User) (string, error)
	VerifyAccessToken(token string) (*models.User, error)
	VerifyRefreshToken(token string) (*models.User, error)
	CreateUser(signUpDto *models.SignUpDto) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
}

type authService struct {
	userService UserService
	env         *config.Env
}

func NewAuthService(userService UserService, env *config.Env) AuthService {
	return &authService{userService: userService, env: env}
}

func (s *authService) CreateAccessToken(user *models.User) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"jti":  uuid.New().String(),
		"sub":  user.UUID,
		"role": user.Role,
		"iss":  s.env.JWTIssuer,
		"aud":  s.env.JWTAudience,
		"iat":  time.Now().Unix(),
		"nbf":  time.Now().Unix(),
		"exp":  time.Now().AddDate(0, 0, s.env.JWTAccessTokenExpirationDays).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(s.env.JWTAccessTokenSecret))

	return accessTokenString, err
}

func (s *authService) CreateRefreshToken(user *models.User) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"jti":  uuid.New().String(),
		"sub":  user.UUID,
		"role": user.Role,
		"iss":  s.env.JWTIssuer,
		"aud":  s.env.JWTAudience,
		"iat":  time.Now().Unix(),
		"nbf":  time.Now().Unix(),
		"exp":  time.Now().AddDate(0, 0, s.env.JWTRefreshTokenExpirationDays).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(s.env.JWTRefreshTokenSecret))

	return refreshTokenString, err
}

func (s *authService) VerifyAccessToken(accessToken string) (*models.User, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.env.JWTAccessTokenSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid access token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid access token claims")
	}

	if claims["aud"] != s.env.JWTAudience {
		return nil, fmt.Errorf("invalid audience")
	}

	if claims["iss"] != s.env.JWTIssuer {
		return nil, fmt.Errorf("invalid issuer")
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, fmt.Errorf("access token expired")
	}

	if claims["nbf"].(float64) > float64(time.Now().Unix()) {
		return nil, fmt.Errorf("access token not valid yet")
	}

	if claims["exp"].(float64)-claims["iat"].(float64) != float64(s.env.JWTAccessTokenExpirationDays*24*60*60) {
		return nil, fmt.Errorf("access token expiration date is invalid")
	}

	user, err := s.userService.GetUserByID(claims["sub"].(string))
	if err != nil {
		return nil, fmt.Errorf("invalid user")
	}

	roleClaim := models.Role(claims["role"].(string))
	if user.Role != roleClaim {
		return nil, fmt.Errorf("user role mismatch")
	}

	return nil, nil
}

func (s *authService) VerifyRefreshToken(refreshToken string) (*models.User, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.env.JWTRefreshTokenSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid refresh token claims")
	}

	if claims["aud"] != s.env.JWTAudience {
		return nil, fmt.Errorf("invalid audience")
	}

	if claims["iss"] != s.env.JWTIssuer {
		return nil, fmt.Errorf("invalid issuer")
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, fmt.Errorf("refresh token expired")
	}

	user, err := s.userService.GetUserByID(claims["sub"].(string))
	if err != nil {
		return nil, fmt.Errorf("invalid user")
	}

	roleClaim := models.Role(claims["role"].(string))
	if user.Role != roleClaim {
		return nil, fmt.Errorf("user role mismatch")
	}

	return user, nil
}

func (s *authService) CreateUser(signUpDto *models.SignUpDto) (*models.User, error) {
	return s.userService.CreateUser(signUpDto)
}

func (s *authService) GetUserByEmail(email string) (*models.User, error) {
	return s.userService.GetUserByEmail(email)
}
