package models

import (
	"github.com/google/uuid"
	"time"
)

type Category struct {
	Id          uuid.UUID `gorm:"primaryKey;column:category_id;type:uuid;default:uuid_generate_v4()" json:"categoryId"`
	Name        string    `gorm:"column:name;type:varchar(255);unique;not null" json:"name"`
	Description string    `gorm:"column:description;type:text;not null" json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;not null;default:now()" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;not null;default:now()" json:"updatedAt"`

	Products []Product `gorm:"foreignKey:CategoryId" json:"products,omitempty"`
}

func (Category) TableName() string {
	return "categories"
}
