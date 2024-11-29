package middlewares

import (
	r "MatchManiaAPI/responses"
	"MatchManiaAPI/services"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	authService services.AuthService
}

func NewAuthMiddleware(authService services.AuthService) AuthMiddleware {
	return AuthMiddleware{authService: authService}
}

func (m *AuthMiddleware) RequireAuth(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		r.Unauthorized(ctx, "Authorization header not found")
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		r.Unauthorized(ctx, "Invalid Authorization header format")
		return
	}

	accessToken := parts[1]

	user, err := m.authService.VerifyAccessToken(accessToken)
	if err != nil {
		r.Unauthorized(ctx, err.Error())
		return
	}

	ctx.Set("user", user)
	ctx.Next()
}
