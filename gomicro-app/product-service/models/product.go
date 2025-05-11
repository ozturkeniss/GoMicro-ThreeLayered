package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Product represents the product model
type Product struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string         `gorm:"size:255;not null"`
	Description string         `gorm:"size:1000;not null"`
	Price       float64        `gorm:"not null"`
	Stock       int            `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"not null"`
	UpdatedAt   time.Time      `gorm:"not null"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// TableName specifies the table name for the Product model
func (Product) TableName() string {
	return "products"
} 