package services

import (
	"github.com/john-henry-alfeche/inventory-app-api/config"
	"github.com/john-henry-alfeche/inventory-app-api/models"
)

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := config.DB.Preload("Category").Find(&products).Error
	return products, err
}

func GetProductById(id string) (models.Product, error) {
	var product models.Product
	err := config.DB.Where("product_id = ?", id).First(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}

func CreateProduct(product models.Product) (models.Product, error) {
	err := config.DB.Create(&product).Error

	err = config.DB.Preload("Category").First(&product, "product_id = ?", product.Id).Error
	return product, err
}

func UpdateProductById(id string, updated models.Product) (models.Product, error) {
	var product models.Product
	err := config.DB.Where("product_id = ?", id).First(&product).Error
	if err != nil {
		return product, err
	}

	product.Name = updated.Name
	product.Description = updated.Description

	err = config.DB.Save(&product).Error
	return product, err
}

func DeleteProductById(id string) error {
	err := config.DB.Where("product_id = ?", id).Delete(&models.Product{}).Error
	return err
}
