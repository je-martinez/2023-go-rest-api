package types

import (
	"main/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Gin    *gin.Engine
	Config *config.Config
}
