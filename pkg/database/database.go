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

var GlobalInstance *DatabaseApiInstance

func StartGlobalInstance(cfg *config.DatabaseConfig) (*DatabaseApiInstance, error) {
	var err error
	GlobalInstance, err = New(cfg)
	return GlobalInstance, err
}

func New(cfg *config.DatabaseConfig) (*DatabaseApiInstance, error) {
	//Establish connection with Database
	database, err := gorm.Open(postgres.Open(getConnectionString(cfg)), &gorm.Config{})
	if err != nil {
		l.ApiLogger.Fatal(constants.CONNECT_DB_ERROR, err)
		return nil, err
	}
	//Auto migrations
	errAutoMigrate := database.AutoMigrate(&e.User{}, &e.Profile{}, e.Post{}, e.Comment{}, e.File{})
	if errAutoMigrate != nil {
		l.ApiLogger.Fatal(constants.DB_MIGRATION_ERROR, err)
		return nil, errAutoMigrate
	}
	//Repositories
	l.ApiLogger.Info(constants.DB_RUNNING)
	return &DatabaseApiInstance{
		db:             database,
		UserRepository: repository.NewRepository[e.User](database, nil),
		PostRepository: repository.NewRepository[e.Post](database, nil),
		FileRepository: repository.NewRepository[e.File](database, nil),
	}, nil
}

type DatabaseApiInstance struct {
	db             *gorm.DB
	UserRepository *repository.GormRepository[e.User]
	PostRepository *repository.GormRepository[e.Post]
	FileRepository *repository.GormRepository[e.File]
}

func getConnectionString(cfg *config.DatabaseConfig) string {

	var ssl string
	if cfg.PostgresqlSSLMode {
		ssl = "require"
	} else {
		ssl = "disable"
	}
	cnx := fmt.Sprintf("postgresql://%s:%s/%s?sslmode=%s&user=%s&password=%s", cfg.PostgresqlHost, cfg.PostgresqlPort, cfg.PostgresqlDbname, ssl, cfg.PostgresqlUser, cfg.PostgresqlPassword)
	return cnx
}
