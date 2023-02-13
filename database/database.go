package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Database *gorm.DB

func Connect() {
	var err error
	Database, err = gorm.Open(postgres.Open("user=postgres password=admin host=localhost port=5432 dbname=activeCitizen sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("database error")

	} else {

		log.Println("Succesfully connected to database")
	}

}
