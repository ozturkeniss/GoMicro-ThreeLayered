package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Order represents the order model
type Order struct {
	ID         uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID     uuid.UUID      `gorm:"type:uuid;not null"`
	ProductID  uuid.UUID      `gorm:"type:uuid;not null"`
	Quantity   int            `gorm:"not null"`
	TotalPrice float64        `gorm:"not null"`
	CreatedAt  time.Time      `gorm:"not null"`
	UpdatedAt  time.Time      `gorm:"not null"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

// TableName specifies the table name for the Order model
func (Order) TableName() string {
	return "orders"
} 