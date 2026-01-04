package repository

import (
	"ecovantory/internal/database"
	"ecovantory/internal/models"
)

func CreateCategory(name string) error {
	newCategory := models.Category{
		CategoryName: name,
	}

	result := database.DB.Create(&newCategory)

	return result.Error
}

func GetAllCategorY() ([]models.Category, error) {
	var categories []models.Category

	result := database.DB.Find(&categories)

	return  categories, result.Error
}