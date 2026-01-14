package repository

import (
	"ecovantory/internal/database"
	"ecovantory/internal/models"
)

func CreateItem(item *models.Category) error {
	result := database.DB.Create(item)
	return result.Error 
}

func GetAllItems() ([]models.Item, error) {
	var items []models.Item
	result := database.DB.Preload("Category").Preload("Unit").Find(&items)
	return  items, result.Error
}

func GetItemById(id uint) (models.Item, error) {
	var item models.Item
	result := database.DB.Preload("Category").Preload("Unit").First(&item, id)
	return  item, result.Error
}

func UpdateItem(item *models.Item) error {
	result := database.DB.Save(item);
	return result.Error
}

func DeleteItem(id uint) error {
	result := database.DB.Delete(&models.Item{}, id)
	return result.Error
}