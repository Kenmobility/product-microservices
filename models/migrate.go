package models

import (
	"gorm.io/gorm"
)

// Migrate function to create tables with updated schema
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&DigitalProduct{})
	db.AutoMigrate(&PhysicalProduct{})
	db.AutoMigrate(&SubscriptionProduct{})
	db.AutoMigrate(&SubscriptionPlan{})
}
