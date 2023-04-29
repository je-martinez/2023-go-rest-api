package server

import (
	"main/api/v1/router"
	"main/config"
	"main/pkg/bucket_manager"
	constants "main/pkg/constants"
	"main/pkg/database"
	"main/pkg/logger"
	types "main/pkg/types"

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
	Server.Database = database.Start(cfg)
	Server.BucketManager = bucket_manager.Start(cfg)

	//Initialize Router And Run Server
	router.Start(Server.Gin)
	logger.ApiLogger.Infof(constants.API_RUNNING, cfg.Server.Port)
	err := Server.Gin.Run(cfg.Server.Address)

	if err != nil {
		logger.ApiLogger.Fatalf(constants.API_RUNNING, err)
	}

}
