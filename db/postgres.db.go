package db

import (
	"fmt"

	"github.com/kenmobility/product-microservice/config"
	"github.com/kenmobility/product-microservice/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgresDb(config config.Config) (*gorm.DB, error) {
	conString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		config.DatabaseHost,
		config.DatabasePort,
		config.DatabaseUser,
		config.DatabaseName,
		config.DatabasePassword,
	)
	fmt.Println("con string: ", conString)
	if helpers.IsLocal() {
		conString += " sslmode=disable"
	}

	return gorm.Open(postgres.Open(conString), &gorm.Config{})
}
