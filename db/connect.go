// DB: https://api.elephantsql.com/console/08526c2d-f401-4194-b66d-a2c45e6bcc97/details

package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database...")
	} else {
		fmt.Println("Successfully connected to database")
	}
}
