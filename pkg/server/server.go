package server

import (
	router "main/api/v1/router"
	"main/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	gin *gin.Engine
	cfg *config.Config
}

func Start(cfg *config.Config) *Server {
	if !cfg.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	s := &Server{gin: gin.New(), cfg: cfg}
	router.Start(s.gin)
	s.gin.Run(cfg.Server.Address)
	return s
}
