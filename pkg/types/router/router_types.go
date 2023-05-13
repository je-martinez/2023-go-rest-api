package router_types

import (
	"github.com/je-martinez/2023-go-rest-api/config"
	"github.com/je-martinez/2023-go-rest-api/pkg/bucket_manager"
	"github.com/je-martinez/2023-go-rest-api/pkg/cache"
	"github.com/je-martinez/2023-go-rest-api/pkg/database"
)

type RouterHandlerProps struct {
	Database      *database.DatabaseApiInstance
	BucketManager *bucket_manager.MinioApiInstance
	Redis         *cache.RedisApiInstance
	Config        *config.ServerConfig
}
