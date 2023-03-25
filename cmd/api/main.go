package main

import (
	"log"
	"main/config"
	constants "main/pkg/constants"
	"main/pkg/server"
	"main/pkg/utils"
	"os"
)

func main() {
	configPath := utils.GetConfigPath(os.Getenv("config"))
	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf(constants.LOADING_CONFIG_ERROR, err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf(constants.PARSING_CONFIG_ERROR, err)
	}

	server.Start(cfg)
}
