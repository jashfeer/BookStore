package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB() (*gorm.DB, error) {
	godotenv.Load()
	dbhost := os.Getenv("dbHost")
	dbuser := os.Getenv("dbUser")
	dbpassword := os.Getenv("dbPassword")
	dbname := os.Getenv("dbName")
	dbport := os.Getenv("dbPort")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s", dbhost, dbuser, dbpassword, dbname, dbport)
	Db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	return Db, nil
}