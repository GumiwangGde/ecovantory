package service

import (
	"ecovantory/internal/models"
	"ecovantory/internal/repository"
	"errors"
)

func CreateStockService(stock *models.Stock) error {
	_, err := repository.GetItemById(stock.ItemID)
	if err != nil {
		return errors.New("item not found, can't create stock")
	}

	if stock.Quantity < 0 {
		return errors.New("stock cannot be minus")
	}

	return repository.CreateStock(stock)
}

func GetAllStockService() ([]models.Stock, error) {
	return repository.GetAllStock()
}

func GetStockByIdService(id uint) (models.Stock, error) {
	stock, err := repository.GetStockById(id)
	if err != nil {
		return models.Stock{}, errors.New("stock not found")
	}

	return stock, nil
}

func UpdateStockService(id uint, updatedData *models.Stock) error {
	oldStock, err := repository.GetStockById(id)
	if err != nil {
		return errors.New("stock not found")
	}

	oldStock.PhysicalStatus = updatedData.PhysicalStatus
	oldStock.PositionStatus = updatedData.PositionStatus
	oldStock.Quantity = updatedData.Quantity

	return repository.UpdateStock(&oldStock)
}

func DeleteStockService(id uint) error {
	_, err := repository.GetStockById(id)
	if err != nil {
		return  errors.New("stock not found")
	}

	return repository.DeleteCategory(id)
}