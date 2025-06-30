package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	fmt.Println("Connecting to the database....")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	if host == "" || user == "" || dbname == "" || port == "" {
		log.Fatal("Database environment variables are not set properly")
	}

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s",
		host, user, dbname, port, sslmode)
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while connecting to the db", err)
	}
	fmt.Println("Connected to the database!")
	db = d

}
func GetDB() *gorm.DB {
	return db
}
