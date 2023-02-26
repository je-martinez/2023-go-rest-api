package server

import (
	router "main/api/v1/router"
	"main/config"
	types "main/pkg/types"

	"github.com/gin-gonic/gin"
)

func Start(cfg *config.Config) *types.Server {
	if !cfg.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	server := &types.Server{Gin: gin.New(), Config: cfg}
	router.Start(server.Gin)
	server.Gin.Run(cfg.Server.Address)
	return server
}
