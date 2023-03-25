package database

import (
	"fmt"
	"main/config"
	constants "main/pkg/constants"
	e "main/pkg/database/entities"
	"main/pkg/database/repository"
	l "main/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Instance
var Database *gorm.DB

// Repositories Instance

type ABC struct {
}

func Start(cfg *config.Config) *gorm.DB {
	database, err := gorm.Open(postgres.Open(getConnectionString(cfg)), &gorm.Config{})
	if err != nil {
		l.ApiLogger.Fatal(constants.CONNECT_DB_ERROR, err)
	}
	errAutoMigrate := database.AutoMigrate(&e.User{}, &e.Profile{}, e.Post{}, e.Comment{}, e.File{})
	if errAutoMigrate != nil {
		l.ApiLogger.Fatal(constants.DB_MIGRATION_ERROR, err)
	}
	initRepositories(database)
	Database = database
	l.ApiLogger.Info(constants.DB_RUNNING)
	return database
}

func getConnectionString(cfg *config.Config) string {

	db := cfg.Database
	var ssl string
	if db.PostgresqlSSLMode {
		ssl = "require"
	} else {
		ssl = "disable"
	}
	cnx := fmt.Sprintf("postgresql://%s:%s/%s?sslmode=%s&user=%s&password=%s", db.PostgresqlHost, db.PostgresqlPort, db.PostgresqlDbname, ssl, db.PostgresqlUser, db.PostgresqlPassword)
	return cnx
}

func initRepositories(db *gorm.DB) {
	UserRepository = repository.NewRepository[e.User](db, nil)
}
