package configs

import (
	"github.com/gin-gonic/gin"
	"os"
	"urlzipper/internal/configs/clients"
	"urlzipper/internal/configs/env"
	"urlzipper/internal/v1/zipper/mappers"
	"urlzipper/internal/v1/zipper/repositories"
	"urlzipper/internal/v1/zipper/services"

	"urlzipper/internal/v1/zipper/controllers"
)

type App interface {
	Setup(engine *gin.Engine)
}

type app struct {
}

func NewApp() App {
	return &app{}
}

func (*app) Setup(engine *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)

	scope := os.Getenv("SCOPE")
	envConfig := env.GetEnvConfig(scope)

	//clients
	redisClient := clients.NewRedisClient(&envConfig.RedisConfig)

	// Mappers
	urlMapper := mappers.NewURLMapper(&envConfig.RedisConfig)

	// Repositories
	urlRepo := repositories.NewURLRepository(&envConfig.RedisConfig, redisClient)

	// Services
	urlService := services.NewURLService(urlMapper, urlRepo)

	// Controllers
	urlController := controllers.NewURLController(urlService)

	urlController.Setup(engine)
}
