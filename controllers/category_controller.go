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

func GetCategories(c *gin.Context) {
	categories, err := services.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func GetCategoryById(c *gin.Context) {
	categoryId := c.Param("category_id")

	if _, err := uuid.Parse(categoryId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	category, err := services.GetCategoryByID(categoryId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error while fetching category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func CreateCategory(c *gin.Context) {
	var input dto.CreateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	category := models.Category{
		Name:        input.Name,
		Description: input.Description,
	}

	created, err := services.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": created})
}

func UpdateCategoryById(c *gin.Context) {
	categoryId := c.Param("category_id")

	if _, err := uuid.Parse(categoryId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	var input dto.UpdateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	update := models.Category{
		Name:        input.Name,
		Description: input.Description,
	}

	updatedCategory, err := services.UpdateCategoryByID(categoryId, update)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedCategory})
}

func DeleteCategory(c *gin.Context) {
	categoryId := c.Param("category_id")

	if _, err := uuid.Parse(categoryId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	err := services.DeleteCategoryById(categoryId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "category deleted"})
}
