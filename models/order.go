package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerName string     `json:"customer_name" gorm:"column:customer_name;type:varchar(155);not null"`
	Items        []Item     `json:"items"`
	OrderedAt    *time.Time `json:"ordered_at"`
}

type UpdateOrder struct {
	OrderedAt    string       `json:"ordered_at"`
	CustomerName string       `json:"customer_name"`
	Items        []UpdateItem `json:"items"`
}
