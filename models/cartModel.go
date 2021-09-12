package models

import (
	"time"
)

type Cart struct {
	UserID			uint 		`gorm:"primaryKey;autoIncrement:false"`
	ProductName		string 		`gorm:"primaryKey;not null;type:varchar(50)" json:"product_name"`
	Quantity		uint 		`gorm:"not null;type:smallint" json:"quantity"`
	CreatedAt 		time.Time
	UpdatedAt		time.Time
}

type CartAPI struct {
	ProductName 	string		`json:"product_name"`
	Price			uint 		`json:"price_per_product"`
	Quantity		uint 		`json:"quantity"`
}