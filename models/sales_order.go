package models

import (
	"github.com/google/uuid"
	"time"
)

type SalesOrderStatus string

const (
	Pending   SalesOrderStatus = "pending"
	Completed SalesOrderStatus = "completed"
)

type SalesOrder struct {
	Id          uuid.UUID        `gorm:"primaryKey;column:sales_order_id; type:uuid; default:uuid_generate_v4()" json:"saleOrderId"`
	OrNumber    string           `gorm:"column:or_number; type:varchar(255); unique; not null" json:"orNumber"`
	TotalAmount float64          `gorm:"column:total_amount; type:decimal(10,2)" json:"totalAmount"`
	VatAmount   float64          `gorm:"column:vat_amount; type:decimal(10,2)" json:"vatAmount"`
	Status      SalesOrderStatus `gorm:"column:status; type:sales_order_status_enum; not null" json:"status"`
	CreatedAt   time.Time        `gorm:"column:created_at;type:timestamp;not null;default:now()" json:"createdAt"`
	UpdatedAt   time.Time        `gorm:"column:updated_at;type:timestamp;not null;default:now()" json:"updatedAt"`

	CustomerId uuid.UUID `gorm:"column:customer_id; type:uuid;" json:"-"`
	Customer   Customer  `gorm:"foreignKey:CustomerId; references:Id; constraint:OnDelete:CASCADE" json:"customers,omitempty"`

	UserId uuid.UUID `gorm:"column:user_id; type:uuid; not null" json:"-"`
	User   User      `gorm:"foreignKey:UserId; references:Id; constraint:OnDelete:CASCADE" json:"users,omitempty"`
}

func (SalesOrder) TableName() string { return "sales_orders" }
