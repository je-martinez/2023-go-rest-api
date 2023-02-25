package server

import (
	"main/config"
	router "main/pkg/router"

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
