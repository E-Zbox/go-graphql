package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBInstance *gorm.DB

func ConnectDb() {
	// create a new PostgreSQL database connection
	// "host=localhost user=dibia password=password123 dbname=todo_app port=5432 timezone=UTC"
	POSTGRES_DBNAME := os.Getenv("POSTGRES_DBNAME")
	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_PORT := os.Getenv("POSTGRES_PORT")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")

	dsn := "host=" + POSTGRES_HOST + " user=" + POSTGRES_USER + " password=" + POSTGRES_PASSWORD + " dbname=" + POSTGRES_DBNAME + " port=" + POSTGRES_PORT + " timezone=UTC"

	// openn a connection to the database
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	if err := db.AutoMigrate(Todo{}); err != nil {
		panic("Failed to perform migrations: " + err.Error())
	}

	if err := db.AutoMigrate(User{}); err != nil {
		panic("Failed to perform migrations: " + err.Error())
	}

	DBInstance = db
}
