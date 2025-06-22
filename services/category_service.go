package services

import (
	"github.com/john-henry-alfeche/inventory-app-api/config"
	"github.com/john-henry-alfeche/inventory-app-api/models"
)

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := config.DB.Find(&categories).Error
	return categories, err
}

func GetCategoryByID(id string) (models.Category, error) {
	var category models.Category
	err := config.DB.Where("category_id = ?", id).First(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func CreateCategory(category models.Category) (models.Category, error) {
	err := config.DB.Create(&category).Error
	return category, err
}

func UpdateCategoryByID(id string, updated models.Category) (models.Category, error) {
	var category models.Category
	err := config.DB.Where("category_id = ?", id).First(&category).Error
	if err != nil {
		return category, err
	}

	category.Name = updated.Name
	category.Description = updated.Description

	err = config.DB.Save(&category).Error
	return category, err
}

func DeleteCategoryById(id string) error {
	err := config.DB.Where("category_id = ?", id).Delete(&models.Category{}).Error
	return err
}
