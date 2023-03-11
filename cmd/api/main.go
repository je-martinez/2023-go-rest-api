package main

import (
	"log"
	"main/config"
	"main/pkg/server"
	"main/pkg/utils"
	"os"
)

func main() {
	log.Println("Starting 2023 Go Rest API")
	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	server.Start(cfg)
}
