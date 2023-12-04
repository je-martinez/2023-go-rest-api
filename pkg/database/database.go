package database

import (
	"fmt"

	"github.com/je-martinez/2023-go-rest-api/config"
	constants "github.com/je-martinez/2023-go-rest-api/pkg/constants"
	e "github.com/je-martinez/2023-go-rest-api/pkg/database/entities"
	"github.com/je-martinez/2023-go-rest-api/pkg/database/repository"
	"github.com/je-martinez/2023-go-rest-api/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfg *config.DatabaseConfig, logger *logger.ApiLogger) (*DatabaseApiInstance, error) {
	//Establish connection with Database
	database, err := gorm.Open(postgres.Open(getConnectionString(cfg)), &gorm.Config{})
	if err != nil {
		logger.Fatal(constants.CONNECT_DB_ERROR, err)
		return nil, err
	}
	//Auto migrations
	errAutoMigrate := database.AutoMigrate(&e.User{}, &e.Profile{}, e.Post{}, e.Comment{}, e.File{}, e.Reaction{})
	if errAutoMigrate != nil {
		logger.Fatal(constants.DB_MIGRATION_ERROR, err)
		return nil, errAutoMigrate
	}
	//Repositories
	logger.Info(constants.DB_RUNNING)
	return &DatabaseApiInstance{
		db:                 database,
		logger:             logger,
		UserRepository:     repository.NewRepository[e.User](database, nil),
		PostRepository:     repository.NewRepository[e.Post](database, nil),
		FileRepository:     repository.NewRepository[e.File](database, nil),
		ReactionRepository: repository.NewRepository[e.Reaction](database, nil),
	}, nil
}

type DatabaseApiInstance struct {
	db                 *gorm.DB
	logger             *logger.ApiLogger
	UserRepository     *repository.GormRepository[e.User]
	PostRepository     *repository.GormRepository[e.Post]
	FileRepository     *repository.GormRepository[e.File]
	ReactionRepository *repository.GormRepository[e.Reaction]
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
