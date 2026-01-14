package repository

import (
	"ecovantory/internal/models"
	"ecovantory/internal/database"
)

func CreateStock(stock *models.Stock) error {
	result := database.DB.Create(stock)
	return result.Error
}

func GetAllStock() ([]models.Stock, error) {
	var stocks []models.Stock
	result := database.DB.Preload("Item.Unit").Preload("Item.Categories").Find(&stocks)
	return  stocks, result.Error
}

func GetStockById(id uint) (models.Stock, error) {
	var stock models.Stock
	result := database.DB.Preload("Item.Unit").Preload("Item.Categories").First(&stock)
	return  stock, result.Error
}

func UpdateStock(stock *models.Stock) error {
	result := database.DB.Save(stock)
	return result.Error
}

func DeleteStock(id uint) error {
	result := database.DB.Delete(&models.Stock{}, id)
	return  result.Error
}

