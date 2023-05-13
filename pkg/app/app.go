package app

import (
	"github.com/je-martinez/2023-go-rest-api/config"
	"github.com/je-martinez/2023-go-rest-api/pkg/bucket_manager"
	"github.com/je-martinez/2023-go-rest-api/pkg/cache"
	constants "github.com/je-martinez/2023-go-rest-api/pkg/constants"
	"github.com/je-martinez/2023-go-rest-api/pkg/database"
	"github.com/je-martinez/2023-go-rest-api/pkg/logger"
	"github.com/je-martinez/2023-go-rest-api/pkg/router"
	router_types "github.com/je-martinez/2023-go-rest-api/pkg/types/router"

	"github.com/gin-gonic/gin"
)

var app *Server = (new(Server))

type Server struct {
	Router        *router.RouterApiInstance
	Logger        *logger.ApiLogger
	Cache         *cache.RedisApiInstance
	Config        *config.Config
	Database      *database.DatabaseApiInstance
	BucketManager *bucket_manager.MinioApiInstance
}

func New(cfg *config.Config) *Server {
	if !cfg.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	//Initialize Packages
	app.Config = cfg
	app.Logger = logger.New(cfg)
	app.Logger.InitLogger()
	app.Database, _ = database.New(&app.Config.Database, app.Logger)
	app.BucketManager, _ = bucket_manager.New(&app.Config.AWS, app.Logger)
	app.Router = router.New(cfg.Server.Address, app.Logger, &router_types.RouterHandlerProps{
		Database:      app.Database,
		BucketManager: app.BucketManager,
		Redis:         app.Cache,
		Config:        &app.Config.Server,
	})
	//Initialize Router And Run Server
	app.Router.RegisterRoutes()
	app.Router.Start()
	app.Logger.Infof(constants.API_RUNNING, cfg.Server.Port)
	return app
}

func App() *Server {
	return app
}

func Config() *config.Config {
	return app.Config
}

func Database() *database.DatabaseApiInstance {
	return app.Database
}

func Cache() *cache.RedisApiInstance {
	return app.Cache
}

func Logger() *logger.ApiLogger {
	return app.Logger
}

func BucketManager() *bucket_manager.MinioApiInstance {
	return app.BucketManager
}
