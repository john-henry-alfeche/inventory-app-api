package models

import (
	"github.com/google/uuid"
	"time"
)

type UserRole string

const (
	Admin  UserRole = "admin"
	Teller UserRole = "teller"
)

type User struct {
	Id        uuid.UUID `gorm:"primaryKey; column:user_id; type:uuid; not null; default: uuid_generate_v4()" json:"userId"`
	Username  string    `gorm:"column:username; type:varchar(255); not null" json:"username"`
	Password  string    `gorm:"column:password; type:varchar(255); not null" json:"password"`
	Name      string    `gorm:"column:name; type:varchar(255); not null" json:"name"`
	Role      UserRole  `gorm:"column:role; type:user_role_enum; not null" json:"role"`
	CreatedAt time.Time `gorm:"column:created_at; type:timestamp; not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at; type:timestamp; not null" json:"updatedAt"`
}

func (User) TableName() string { return "users" }
