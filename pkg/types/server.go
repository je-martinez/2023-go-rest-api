package types

import (
	"github.com/je-martinez/2023-go-rest-api/config"
	"github.com/je-martinez/2023-go-rest-api/pkg/bucket_manager"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Gin           *gin.Engine
	Config        *config.Config
	Database      *gorm.DB
	BucketManager *bucket_manager.MinioApiInstance
}
