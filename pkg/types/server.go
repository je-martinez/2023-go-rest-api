package types

import (
	"github.com/je-martinez/2023-go-rest-api/config"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type Server struct {
	Gin           *gin.Engine
	Config        *config.Config
	Database      *gorm.DB
	BucketManager *minio.Client
}
