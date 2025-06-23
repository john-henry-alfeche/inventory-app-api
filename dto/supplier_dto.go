package dto

type Supplier struct {
	Name          string `json:"name" binding:"required"`
	Email         string `json:"email" binding:"required"`
	ContactNumber string `json:"contactNumber" binding:"required"`
	Address       string `json:"address" binding:"required"`
}

type CreateSupplierInput struct {
	Supplier `binding:"required"`
}

type UpdateSupplierInput struct {
	Supplier `binding:"required"`
}
