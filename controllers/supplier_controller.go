package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/john-henry-alfeche/inventory-app-api/dto"
	"github.com/john-henry-alfeche/inventory-app-api/models"
	"github.com/john-henry-alfeche/inventory-app-api/services"
	"gorm.io/gorm"
	"net/http"
)

func GetAllSuppliers(c *gin.Context) {
	suppliers, err := services.GetAllSuppliers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": suppliers})
}

func GetSupplierById(c *gin.Context) {
	supplierId := c.Param("supplier_id")

	if _, err := uuid.Parse(supplierId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	supplier, err := services.GetSupplierById(supplierId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "supplier not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": supplier})
}

func CreateSupplier(c *gin.Context) {
	var input dto.CreateSupplierInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	supplier := models.Supplier{
		Name:          input.Name,
		ContactNumber: input.ContactNumber,
		Email:         input.Email,
		Address:       input.Address,
	}

	created, err := services.CreateSupplier(supplier)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": created})
}

func UpdateSupplier(c *gin.Context) {
	supplierId := c.Param("supplier_id")

	if _, err := uuid.Parse(supplierId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input dto.UpdateSupplierInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := models.Supplier{
		Name:          input.Name,
		ContactNumber: input.ContactNumber,
		Email:         input.Email,
		Address:       input.Address,
	}

	updateSupplier, err := services.UpdateSupplierById(supplierId, update)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "supplier not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updateSupplier})
}

func DeleteSupplier(c *gin.Context) {
	supplierId := c.Param("supplier_id")
	if _, err := uuid.Parse(supplierId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.DeleteSupplierById(supplierId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "supplier not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "supplier deleted"})
}
