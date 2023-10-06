package database

import (
	"assignment-2/models"
	"fmt"
)

func Migrate() {
	fmt.Println("Processing to Migrate Databases")

	DB.AutoMigrate(&models.Order{}, &models.Item{})

	fmt.Println("Success to Migrate Databases")
}
