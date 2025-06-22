package models

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	Id          uuid.UUID `gorm:"primaryKey;column:product_id;type:uuid;default:uuid_generate_v4()" json:"productId"`
	Name        string    `gorm:"column:name;type:varchar(255);unique;not null" json:"name"`
	Description string    `gorm:"column:description;text" json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;not null;default:now()" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;not null;default:now()" json:"updatedAt"`

	CategoryId uuid.UUID `gorm:"type:uuid;not null;column:category_id" json:"-"`
	Category   Category  `gorm:"foreignKey:CategoryId;constraint:OnDelete:CASCADE" json:"category,omitempty"`
}

func (Product) TableName() string {
	return "products"
}
