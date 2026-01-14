package service

import (
	"ecovantory/internal/models"
	"ecovantory/internal/repository"
	"errors"
)

func CreateItemService(item *models.Item, categoryIDs []uint) error {
	if item.ItemName == "" {
		return errors.New("item name cannot be empty")
	}

	var categories []models.Category
	for _, id := range categoryIDs {
		cat, err := repository.GetCategoryById(id)
		if err != nil {
			return errors.New("one of the categories was not found")
		}
		categories = append(categories, cat)
	}

	item.Categories = categories

	return repository.CreateItem(item)
}

func GetAllItemsService() ([]models.Item, error) {
	items, err := repository.GetAllItems()
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, errors.New("item not found")
	}

	return items, nil
}

func GetItemByIdService(id uint) (models.Item, error) {
	item, err := repository.GetItemById(id)
	if err != nil {
		return models.Item{}, errors.New("item not found")
	}

	return item, nil
}

func UpdateItemService(id uint, updatedData *models.Item, categoryIDs []uint) error {
	item, err := repository.GetItemById(id)
	if err != nil {
		return  errors.New("item not found")
	}

	var newCategories []models.Category
	for _, catId := range categoryIDs {
		cat, err := repository.GetCategoryById(catId)
		if err != nil {
			return  errors.New("one of the categories not found")
		}
		newCategories = append(newCategories, cat)
	}
	
	item.ItemName = updatedData.ItemName
	item.ItemCode = updatedData.ItemCode
	item.BasePrice = updatedData.BasePrice
	item.SellPrice = updatedData.SellPrice
	item.Description = updatedData.Description
	item.UnitId = updatedData.UnitId
	item.Categories = newCategories

	return  repository.UpdateItem(&item)
}

func DeleteItemService(id uint) error {
	_, err := repository.GetItemById(id)
	if err != nil {
		return  errors.New("item not found")
	}

	return  repository.DeleteCategory(id)
}
