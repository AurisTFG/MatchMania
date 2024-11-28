package middlewares

import (
	r "MatchManiaAPI/responses"
	"MatchManiaAPI/services"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	authService services.AuthService
}

func NewAuthMiddleware(authService services.AuthService) AuthMiddleware {
	return AuthMiddleware{authService: authService}
}

func (m *AuthMiddleware) RequireAuth(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("AccessToken")
	if err != nil {
		r.Unauthorized(ctx, "Token not found")
		return
	}

	user, err := m.authService.VerifyAccessToken(tokenString)
	if err != nil {
		r.Unauthorized(ctx, err.Error())
		return
	}

	ctx.Set("user", user)
	ctx.Next()
}
