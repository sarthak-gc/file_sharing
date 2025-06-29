package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	fmt.Println("Connecting to the database....")
	dsn := "host=localhost user=sarthakgc dbname=file_sharing_service port=5432 sslmode=disable"

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while connecting to the db", err)
	}
	db = d

}
func GetDB() *gorm.DB {
	return db
}
