package middlewares

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/services"
)

type Middlewares struct {
	AuthMiddleware AuthMiddleware
}

func NewMiddlewares(
	authMiddleware AuthMiddleware,
) Middlewares {
	return Middlewares{
		AuthMiddleware: authMiddleware,
	}
}

func SetupMiddlewares(
	db *config.DB,
	env *config.Env,
) *Middlewares {
	sessionRepository := repositories.NewSessionRepository(db)
	userRepository := repositories.NewUserRepository(db)

	authService := services.NewAuthService(sessionRepository, userRepository, env)

	authMiddleware := NewAuthMiddleware(authService)

	return &Middlewares{
		AuthMiddleware: authMiddleware,
	}
}
