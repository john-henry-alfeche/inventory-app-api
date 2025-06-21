package models

import (
	"github.com/google/uuid"
	"time"
)

type Category struct {
	ID          uuid.UUID `gorm:"column:category_id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"column:name;type:varchar(255);not null;unique"`
	Description string    `gorm:"column:description;type:text;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;not null"`
}

func (Category) TableName() string {
	return "categories"
}
