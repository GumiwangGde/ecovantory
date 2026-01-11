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

func GetAllCategory() ([]models.Category, error) {
	var categories []models.Category
	result := database.DB.Find(&categories)
	return  categories, result.Error
}

func GetCategoryById(id uint) (models.Category, error) {
	var category models.Category
	result := database.DB.First(&category, id)
	return  category, result.Error
}

func UpdateCategory(id uint, newName string) error {
	result := database.DB.Model(&models.Category{}).Where("id = ?", id).Update("category_name", newName)
	return  result.Error
}

func DeleteCategory(id uint) error {
	result := database.DB.Delete(&models.Category{}, id)
	return  result.Error
}