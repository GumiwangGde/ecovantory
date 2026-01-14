package repository

import (
	"ecovantory/internal/database"
	"ecovantory/internal/models"
)

func CreateUnit(unit *models.Unit) error {
	return database.DB.Create(unit).Error
}

func GetAllUnits() ([]models.Unit, error) {
	var units []models.Unit
	result := database.DB.Find(&units)
	return units, result.Error
}

func GetUnitById(id uint) (models.Unit, error) {
	var unit models.Unit
	result := database.DB.First(&unit, id)
	return  unit, result.Error
}

func UpdateUnit(id uint, newName string) error {
	return database.DB.Model(&models.Unit{}).Where("id = ?", id).Update("unit_name", newName).Error
}

func DeleteUnit(id uint) error {
	return database.DB.Delete(&models.Unit{}, id).Error
}