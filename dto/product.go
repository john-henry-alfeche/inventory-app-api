package dto

import "github.com/google/uuid"

type Product struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
type CreateProductInput struct {
	CategoryId uuid.UUID `json:"categoryId" binding:"required,uuid"`
	Product
}

type UpdateProductInput struct {
	Product `binding:"required"`
}
