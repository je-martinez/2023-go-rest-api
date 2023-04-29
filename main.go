package main

import (
	"github.com/je-martinez/2023-go-rest-api/config"
	"github.com/je-martinez/2023-go-rest-api/pkg/cache"
	"github.com/je-martinez/2023-go-rest-api/pkg/logger"
	"github.com/je-martinez/2023-go-rest-api/pkg/server"
)

func main() {
	config := config.InitConfig()
	appLogger := logger.NewApiLogger(config)
	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v", config.Server.AppVersion, config.Logger.Level, config.Server.Mode, config.Server.SSL)
	cache.InitRedisClient(config)
	server.Start(config)
}
