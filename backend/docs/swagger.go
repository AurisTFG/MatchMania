package docs

import (
	"MatchManiaAPI/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupSwagger(server *gin.Engine, env *config.Env) {
	if env.IsDev {
		SwaggerInfo.Host = env.ServerUrl

		server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
