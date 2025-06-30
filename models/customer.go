package models

import (
	"github.com/google/uuid"
	"time"
)

type Customer struct {
	Id        uuid.UUID `gorm:"primaryKey; column:customer_id; type:uuid; default:uuid_generate_v4(); not null" json:"customerId"`
	Name      string    `gorm:"column:name; type:varchar(255); unique; not null" json:"name"`
	Email     string    `gorm:"column:email; type:varchar(255); unique;" json:"email"`
	Phone     string    `gorm:"column:phone; type:varchar(255); unique;" json:"phone"`
	Address   string    `gorm:"column:address; type:varchar(255);" json:"address"`
	CreatedAt time.Time `gorm:"column:created_at; type:timestamp;not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at; type:timestamp;not null" json:"updatedAt"`
}

func (Customer) TableName() string { return "customers" }
