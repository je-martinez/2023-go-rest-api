package database

import (
	"fmt"
	"log"
	"main/config"
	e "main/pkg/database/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

func Start(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(getConnectionString(cfg)), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect with database", db)
	}
	db.AutoMigrate(&e.User{})
	return db
}
