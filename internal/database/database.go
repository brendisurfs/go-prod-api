package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

// NewDatabase - returns a pointer to a new database object.
func NewDatabase() (*gorm.DB, error) {
	fmt.Println("setting up new db connection")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("some error occured. Err: %s", err)
	}

	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUserName, dbTable, dbPassword)
	// now open up a database.
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return db, err
	}

	if err := db.DB().Ping(); err != nil {
		return db, err
	}
	return db, nil
}
