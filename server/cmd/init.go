package main

import (
	"log"

	"github.com/phzeng0726/go-server-template/internal/config"
	"github.com/phzeng0726/go-server-template/internal/database"
	"github.com/phzeng0726/go-server-template/internal/repository"
	"github.com/phzeng0726/go-server-template/internal/service"
	"github.com/phzeng0726/go-server-template/pkg/logger"

	delivery "github.com/phzeng0726/go-server-template/internal/delivery/http"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var (
	repos    *repository.Repositories
	services *service.Services
	handlers *delivery.Handler
	router   *gin.Engine
)

func initServiceDeps(repos *repository.Repositories, loggerManager logger.LoggerManager) service.Deps {
	// tokenManager, err := auth.NewManager(nil, "./public.pem")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	deps := service.Deps{
		Repos:         repos,
		LoggerManager: loggerManager,
		// TokenManager:  tokenManager,
	}

	return deps
}

func init() {
	config.InitConfig()

	loggerManager, err := logger.NewManager("your_server_name_for_logger", config.Env.Env, config.Env.DataPath.LogFolder)
	if err != nil {
		log.Fatalln(err)
	}

	if err := loggerManager.InitLogger(); err != nil {
		log.Fatalln(err)
	}

	// Database
	conn := database.Connect(loggerManager.GetLogger())
	database.SyncDatabase(conn)

	// Others
	repos = repository.NewRepositories(conn)
	deps := initServiceDeps(repos, loggerManager)
	services = service.NewServices(deps)
	handlers = delivery.NewHandler(services, deps.TokenManager)
	router = handlers.Init()

}
