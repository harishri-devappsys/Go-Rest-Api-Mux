package db

import (
	"gorestapi/pkg/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:551551551@localhost:5433/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
