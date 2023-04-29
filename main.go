package main

import (
	"main/config"
	"main/pkg/cache"
	"main/pkg/logger"
	"main/pkg/server"
)

func main() {
	config := config.InitConfig()
	appLogger := logger.NewApiLogger(config)
	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v", config.Server.AppVersion, config.Logger.Level, config.Server.Mode, config.Server.SSL)
	cache.InitRedisClient(config)
	server.Start(config)
}
