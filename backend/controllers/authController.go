package controllers

import (
	"MatchManiaAPI/initializers"
	"MatchManiaAPI/models"
	r "MatchManiaAPI/responses"
	"MatchManiaAPI/services"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func UserSignUp(c *gin.Context) {
	var bodyDto models.SignUpDto

	if err := c.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(c, err.Error())
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(c, err.Error())
		return
	}

	user, err := services.CreateUser(&bodyDto)
	if err != nil {
		r.UnprocessableEntity(c, err.Error())
		return
	}

	r.Created(c, models.UserSignUpResponse{User: user.ToDto()})
}

func UserLogIn(c *gin.Context) {
	var bodyDto models.LoginDto

	if err := c.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(c, err.Error())
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(c, err.Error())
		return
	}

	user, err := services.GetUserByEmail(bodyDto.Email)
	if err != nil {
		r.UnprocessableEntity(c, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(bodyDto.Password))
	if err != nil {
		r.UnprocessableEntity(c, "Invalid email or password")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"jti":  uuid.New().String(),
		"sub":  user.ID,
		"role": user.Role,
		"iss":  initializers.Cfg.JWTIssuer,
		"aud":  initializers.Cfg.JWTAudience,
		"iat":  time.Now().Unix(),
		"nbf":  time.Now().Unix(),
		"exp":  time.Now().AddDate(0, 0, initializers.Cfg.JWTTokenExpirationDays).Unix(),
	})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"jti":  uuid.New().String(),
		"sub":  user.ID,
		"role": user.Role,
		"iss":  initializers.Cfg.JWTIssuer,
		"aud":  initializers.Cfg.JWTAudience,
		"iat":  time.Now().Unix(),
		"nbf":  time.Now().Unix(),
		"exp":  time.Now().AddDate(0, 0, initializers.Cfg.JWTRefreshTokenExpirationDays).Unix(),
	})

	tokenString, err := token.SignedString([]byte(initializers.Cfg.JWTSecret))
	if err != nil {
		r.UnprocessableEntity(c, err.Error())
		return
	}

	refreshTokenString, err := refreshToken.SignedString([]byte(initializers.Cfg.JWTSecret))
	if err != nil {
		r.UnprocessableEntity(c, err.Error())
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("AccessToken", tokenString, initializers.Cfg.JWTTokenExpirationDays*24*60*60, "/", "", false, true)
	c.SetCookie("RefreshToken", refreshTokenString, initializers.Cfg.JWTRefreshTokenExpirationDays*24*60*60, "/", "", false, true)

	r.OK(c, models.UserLoginResponse{AccessToken: tokenString, RefreshToken: refreshTokenString})
}

func UserLogOut(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("AccessToken", "", -1, "/", "", false, true)
	c.SetCookie("RefreshToken", "", -1, "/", "", false, true)

	r.Deleted(c)
}

func UserRefreshToken(c *gin.Context) {
	tokenString, err := c.Cookie("RefreshToken")
	if err != nil {
		r.Unauthorized(c, "Refresh token not found")
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(initializers.Cfg.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		r.Unauthorized(c, "Invalid refresh token")
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		r.Unauthorized(c, "Invalid refresh token claims")
		return
	}

	if claims["aud"] != initializers.Cfg.JWTAudience {
		r.Unauthorized(c, "Invalid audience")
		return
	}

	if claims["iss"] != initializers.Cfg.JWTIssuer {
		r.Unauthorized(c, "Invalid issuer")
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		r.Unauthorized(c, "Refresh token expired")
		return
	}

	user, err := services.GetUserByID(uint(claims["sub"].(float64)))
	if err != nil {
		r.Unauthorized(c, "Invalid user")
		return
	}

	roleClaim := models.Role(claims["role"].(string))
	if user.Role != roleClaim {
		r.Unauthorized(c, "User role mismatch")
		return
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"jti":  uuid.New().String(),
		"sub":  user.ID,
		"role": user.Role,
		"iss":  initializers.Cfg.JWTIssuer,
		"aud":  initializers.Cfg.JWTAudience,
		"iat":  time.Now().Unix(),
		"nbf":  time.Now().Unix(),
		"exp":  time.Now().AddDate(0, 0, initializers.Cfg.JWTTokenExpirationDays).Unix(),
	})

	tokenString, err = token.SignedString([]byte(initializers.Cfg.JWTSecret))
	if err != nil {
		r.UnprocessableEntity(c, err.Error())
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("AccessToken", tokenString, initializers.Cfg.JWTTokenExpirationDays*24*60*60, "/", "", false, true)

	r.OK(c, models.UserRefreshTokenResponse{AccessToken: tokenString})
}
