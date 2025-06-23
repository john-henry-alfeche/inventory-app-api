package services

import (
	"github.com/john-henry-alfeche/inventory-app-api/config"
	"github.com/john-henry-alfeche/inventory-app-api/models"
)

func GetAllSuppliers() ([]models.Supplier, error) {
	var suppliers []models.Supplier
	err := config.DB.Find(&suppliers).Error
	return suppliers, err
}

func GetSupplierById(id string) (models.Supplier, error) {
	var supplier models.Supplier
	err := config.DB.Find(&supplier, "supplier_id = ?", id).Error

	if err != nil {
		return supplier, err
	}

	return supplier, nil
}

func CreateSupplier(product models.Supplier) (models.Supplier, error) {
	err := config.DB.Create(&product).Error

	return product, err
}

func UpdateSupplierById(id string, updated models.Supplier) (models.Supplier, error) {
	var supplier models.Supplier
	err := config.DB.Find(&supplier, "supplier_id = ?", id).Error

	if err != nil {
		return supplier, err
	}
	err = config.DB.Model(&supplier).Updates(updated).Where("supplier_id =?", id).Error

	if err != nil {
		return supplier, err
	}

	err = config.DB.First(&supplier, "supplier_id = ?", id).Error
	return supplier, err

}

func DeleteSupplierById(id string) error {
	err := config.DB.Delete(&models.Supplier{}, id).Error
	return err
}
