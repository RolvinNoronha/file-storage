package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDb() (*gorm.DB, error) {

	connStr := os.Getenv("DB_STRING")
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal("Could not connect to the database: ", err)
	}

	return db, err
}
