package models

import "github.com/google/uuid"

type Product struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
}
