package middlewares

import (
	"MatchManiaAPI/constants"
	"MatchManiaAPI/services"
	r "MatchManiaAPI/utils/httpresponses"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	authService services.AuthService
}

func NewAuthMiddleware(authService services.AuthService) AuthMiddleware {
	return AuthMiddleware{authService: authService}
}

func (m *AuthMiddleware) RequireAuth(ctx *gin.Context) {
	accessToken, err := ctx.Cookie(constants.AccessTokenName)
	if err != nil || accessToken == "" {
		r.Unauthorized(ctx)
		return
	}

	user, err := m.authService.VerifyAccessToken(accessToken)
	if err != nil {
		r.Unauthorized(ctx)
		return
	}

	ctx.Set("userId", user.Id)
	ctx.Next()
}
