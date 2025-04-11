package config

import (
	"MatchManiaAPI/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupSwagger(server *gin.Engine, env *Env) {
	if env.IsDev {
		docs.SwaggerInfo.Host = env.ServerURL

		server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
