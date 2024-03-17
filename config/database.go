package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host   = "DB_HOST"
	user   = "DB_USER"
	pass   = "DB_PASS"
	dbname = "DB_NAME"
	dbport = "DB_PORT"
)

func InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv(host), os.Getenv(user), os.Getenv(pass), os.Getenv(dbname), os.Getenv(dbport))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database : %v", err)
	}

	return db
}
