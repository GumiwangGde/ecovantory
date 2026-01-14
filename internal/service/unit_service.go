package service

import (
	"ecovantory/internal/repository"
	"ecovantory/internal/models"
	"errors"
)

func CreateUnitService(name string) error {
	if name == "" {
		return errors.New("unit name cannot be empty")
	}

	unit := models.Unit{UnitName: name}
	
	return  repository.CreateUnit(&unit)
}

func GetAllUnitService() ([]models.Unit, error) {
	units, err := repository.GetAllUnits()
	if err != nil {
		return  nil, err
	}

	return units, nil
}

func GetUnitByIdService(id uint) (models.Unit, error) {
	unit, err := repository.GetUnitById(id)
	if err != nil {
		return  models.Unit{}, errors.New("unit not found")
	}

	return unit, nil
}

func UpdateUnitService(id uint, newName string) error {
	if newName == "" {
		return errors.New("unit name cannot be empty")
	}

	_, err := repository.GetUnitById(id)
	if err != nil {
		return errors.New("unit not found")
	}

	return repository.UpdateUnit(id, newName)
}

func DeleteUnitService(id uint) error {
	_, err := repository.GetUnitById(id)
	if err != nil {
		return errors.New("unit not found")
	}

	return repository.DeleteUnit(id)
}