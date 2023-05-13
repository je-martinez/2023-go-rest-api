package server

import (
	"github.com/je-martinez/2023-go-rest-api/api/v1/router"
	"github.com/je-martinez/2023-go-rest-api/config"
	"github.com/je-martinez/2023-go-rest-api/pkg/bucket_manager"
	constants "github.com/je-martinez/2023-go-rest-api/pkg/constants"
	"github.com/je-martinez/2023-go-rest-api/pkg/database"
	"github.com/je-martinez/2023-go-rest-api/pkg/logger"
	types "github.com/je-martinez/2023-go-rest-api/pkg/types/server"

	"github.com/gin-gonic/gin"
)

var Server *types.Server

func Start(cfg *config.Config) {

	if !cfg.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	Server = new(types.Server)

	//Initialize Packages
	Server.Gin = gin.New()
	Server.Config = cfg
	Server.Database, _ = database.StartGlobalInstance(&cfg.Database)
	Server.BucketManager, _ = bucket_manager.StartGlobalInstance(&cfg.AWS)

	//Initialize Router And Run Server
	router.Start(Server.Gin)
	logger.ApiLogger.Infof(constants.API_RUNNING, cfg.Server.Port)
	err := Server.Gin.Run(cfg.Server.Address)

	if err != nil {
		logger.ApiLogger.Fatalf(constants.API_RUNNING, err)
	}

}
