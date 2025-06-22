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

func GetAllProducts(c *gin.Context) {
	products, err := services.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func GetProductById(c *gin.Context) {
	productId := c.Param("product_id")

	if _, err := uuid.Parse(productId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	product, err := services.GetProductById(productId)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error while fetching product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func CreateProduct(c *gin.Context) {
	var input dto.CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	product := models.Product{
		CategoryId:  input.CategoryId,
		Name:        input.Name,
		Description: input.Description,
	}

	created, err := services.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": created})
}

func UpdateProductById(c *gin.Context) {
	productId := c.Param("product_id")

	if _, err := uuid.Parse(productId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	var input dto.UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	update := models.Product{
		Name:        input.Name,
		Description: input.Description,
	}

	updateProduct, err := services.UpdateProductById(productId, update)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updateProduct})
}

func DeleteProductById(c *gin.Context) {
	productId := c.Param("product_id")

	if _, err := uuid.Parse(productId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	err := services.DeleteProductById(productId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}
