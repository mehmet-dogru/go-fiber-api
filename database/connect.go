package database

import (
	"fmt"
	"github.com/mehmet-dogru/go-fiber-api/config"
	"github.com/mehmet-dogru/go-fiber-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

func ConnectDB() {
	var err error

	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("Failed to parse database port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Fail connection to database")
	}
	fmt.Println("Connection open to database")

	DB.AutoMigrate(&models.Task{})
	fmt.Println("Database migrated")
}
