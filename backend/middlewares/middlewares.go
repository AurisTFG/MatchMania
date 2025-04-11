package middlewares

import (
	"MatchManiaAPI/services"
)

type Middlewares struct {
	AuthMiddleware AuthMiddleware
}

func SetupMiddlewares(
	services *services.Services,
) *Middlewares {
	authMiddleware := NewAuthMiddleware(services.AuthService)

	return &Middlewares{
		AuthMiddleware: authMiddleware,
	}
}
