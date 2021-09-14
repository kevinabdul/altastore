package models

import (
	"time"
)

type Product struct {
	ProductID 		uint 		`gorm:"primaryKey"`
	ProductName		string		`gorm:"type:varchar(70);not null;unique" json:"product_name" form:"product_name"`
	CategoryID		uint		`gorm:"type:varchar(70);not null" json:"category" form:"category"`
	Price 			uint 		`gorm:"type:int unsigned;not null" json:"price" form:"price"`
	CreatedAt		time.Time 	
	UpdatedAt		time.Time	
}

type ProductAPI struct {
	ProductID   	uint 		`json:"product_id"`
	ProductName		string		`json:"product_name" form:"product_name"`
	CategoryName	string		`json:"category_name" form:"category_name"`
	Price 			uint 		`json:"price" form:"price"`
}