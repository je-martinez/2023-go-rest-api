package server

import (
	router "main/api/v1/router"
	"main/config"
	"main/pkg/database"
	types "main/pkg/types"

	"github.com/gin-gonic/gin"
)

func Start(cfg *config.Config) *types.Server {
	if !cfg.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	server := &types.Server{Gin: gin.New(), Config: cfg, Database: database.Start(cfg)}
	router.Start(server.Gin)
	server.Gin.Run(cfg.Server.Address)
	return server
}
