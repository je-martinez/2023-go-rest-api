package server_types

import (
	"github.com/je-martinez/2023-go-rest-api/config"
	"github.com/je-martinez/2023-go-rest-api/pkg/bucket_manager"
	"github.com/je-martinez/2023-go-rest-api/pkg/database"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Gin           *gin.Engine
	Config        *config.Config
	Database      *database.DatabaseApiInstance
	BucketManager *bucket_manager.MinioApiInstance
}
