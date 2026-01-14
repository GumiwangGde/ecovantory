package service

import (
	"ecovantory/internal/models"
	"ecovantory/internal/repository"
	"errors"
)

func CreateCategoryService(name string) error {
	if name == "" {
		return  errors.New("category name cannot be empty")
	}

	return  repository.CreateCategory(name)
}

func GetAllCategoriesService() ([]models.Category, error) {
	categories, err := repository.GetAllCategory()
	if err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return nil, errors.New("there are no categories registered in the system yet")
	}

	return categories, nil
}

func GetCategoryByIDService(id uint) (models.Category, error) {
	category, err := repository.GetCategoryById(id)
	if err != nil {
		return  models.Category{}, errors.New("category not found")
	}

	return category, nil
}

func UpdateCategoryService(id uint, newName string) error {
	if newName == "" {
		return  errors.New("category name cannot be empty")
	}

	_, err := repository.GetCategoryById(id)
	if err != nil {
		return  errors.New("category not found")
	}

	return  repository.UpdateCategory(id, newName)
}

func DeleteCategoryService(id uint) error {
	_, err := repository.GetCategoryById(id)
	if err != nil {
		return  errors.New("category not found")
	}

	return repository.DeleteCategory(id)
}