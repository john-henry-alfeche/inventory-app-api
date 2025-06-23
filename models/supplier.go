package models

import (
	"github.com/google/uuid"
	"time"
)

type Supplier struct {
	Id            uuid.UUID `gorm:"primaryKey;column:supplier_id;type:uuid;default:uuid_generate_v4()" json:"supplierId"`
	Name          string    `gorm:"column:name;type:varchar(255);not null" json:"name"`
	ContactNumber string    `gorm:"column:contact_number;type:varchar(255);not null" json:"contactNumber"`
	Email         string    `gorm:"column:email;type:varchar(255);" json:"email"`
	Address       string    `gorm:"column:address;type:text;not null" json:"address"`
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp;not null" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updatedAt"`
}

func (Supplier) TableName() string {
	return "suppliers"
}
