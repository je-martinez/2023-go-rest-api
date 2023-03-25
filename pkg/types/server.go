package types

import (
	"main/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Gin      *gin.Engine
	Config   *config.Config
	Database *gorm.DB
}
