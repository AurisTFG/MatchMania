package middlewares

import (
	"MatchManiaAPI/initializers"
	"MatchManiaAPI/models"
	r "MatchManiaAPI/responses"
	"MatchManiaAPI/services"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("AccessToken")
	if err != nil {
		r.Unauthorized(c, "Token not found")
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(initializers.Cfg.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		r.Unauthorized(c, "Invalid token")
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		r.Unauthorized(c, "Invalid token claims")
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
		r.Unauthorized(c, "Token expired")
		return
	}

	if claims["nbf"].(float64) > float64(time.Now().Unix()) {
		r.Unauthorized(c, "Token not valid yet")
		return
	}

	if claims["exp"].(float64)-claims["iat"].(float64) != float64(initializers.Cfg.JWTRefreshTokenExpirationDays*24*60*60) {
		r.Unauthorized(c, "Token expiration date is invalid")
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

	c.Set("user", user)
	c.Next()
}
