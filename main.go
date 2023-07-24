package main

import (
	"github.com/je-martinez/2023-go-rest-api/config"
	"github.com/je-martinez/2023-go-rest-api/pkg/app"
)

func main() {
	config := config.InitConfig()
	app.New(config)
}
