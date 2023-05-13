package database

import (
	"fmt"

	"github.com/je-martinez/2023-go-rest-api/config"
	constants "github.com/je-martinez/2023-go-rest-api/pkg/constants"
	e "github.com/je-martinez/2023-go-rest-api/pkg/database/entities"
	"github.com/je-martinez/2023-go-rest-api/pkg/database/repository"
	l "github.com/je-martinez/2023-go-rest-api/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Instance
var Database *gorm.DB

// Repositories Instance
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
	PostRepository = repository.NewRepository[e.Post](db, nil)
	FileRepository = repository.NewRepository[e.File](db, nil)

}
