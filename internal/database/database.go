package database

import (
	"ecovantory/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB	

func InitDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("ecovantory.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Unit{},
		&models.Item{},
		&models.ItemCategory{},
		&models.Stock{},
		&models.StockMutations{},
		&models.Transaction{},
		&models.TransactionDetail{},
		&models.Log{},
	)

	if err != nil {
		log.Fatal("Failed to migrate to database: ", err)
	}

	log.Println("Database successfully initialized and migrated")
}