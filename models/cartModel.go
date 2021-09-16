package models

import (
	"time"
)

type Cart struct {
	UserID			uint 		`gorm:"primaryKey;autoIncrement:false"`
	ProductID		uint 		`gorm:"primaryKey;not null" json:"product_id"`
	Quantity		uint 		`gorm:"not null;type:int" json:"quantity"`
	CreatedAt 		time.Time
	UpdatedAt		time.Time
	User            User 		`gorm:"foreignKey:UserID"`
	Product  		Product  	`gorm:"foreignKey:ProductID"`
}

type CartAPI struct {
	ProductID   	uint 		`json:"product_id"`
	ProductName 	string		`json:"product_name"`
	Price			uint 		`json:"price_per_product"`
	Quantity		uint 		`json:"quantity"`
}