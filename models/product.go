package models

import (
	"time"
)

type Product struct {
	ID          uint   `gorm:"primary_key"`               // Serial integer as primary key
	PublicID    string `gorm:"type:varchar;unique_index"` // UUID as public ID
	Name        string
	Description string
	Price       float32
	CreatedAt   time.Time
	UpdatedAt   time.Time

	DigitalProduct      *DigitalProduct
	PhysicalProduct     *PhysicalProduct
	SubscriptionProduct *SubscriptionProduct
}

type DigitalProduct struct {
	ID           uint `gorm:"primary_key"`
	ProductID    uint `gorm:"not null;unique_index"` // Foreign key
	FileSize     int32
	DownloadLink string
}

type PhysicalProduct struct {
	ID         uint `gorm:"primary_key"`
	ProductID  uint `gorm:"not null;unique_index"` // Foreign key
	Weight     float32
	Dimensions string
}

type SubscriptionProduct struct {
	ID                 uint `gorm:"primary_key"`
	ProductID          uint `gorm:"not null;unique_index"` // Foreign key
	SubscriptionPeriod string
	RenewalPrice       float32
}

type SubscriptionPlan struct {
	ID        uint `gorm:"primary_key"`    // Serial integer as primary key
	ProductID uint `gorm:"not null;index"` // Foreign key
	PlanName  string
	Duration  int
	Price     float64
}
