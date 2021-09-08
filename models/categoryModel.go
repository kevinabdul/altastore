package models

import (
	"time"
)

type Category struct {
	CategoryID		uint 		`gorm:"primaryKey"`
	CategoryName	string		`gorm:"type:varchar(70);not null;unique" json:"category_name" form:"category_name"`
	CreatedAt		time.Time 	
	UpdatedAt		time.Time	
}

type CategoryAPI struct {
	CategoryID		uint 		`gorm:"primaryKey"`
	CategoryName	string		`gorm:"type:varchar(70);not null;unique" json:"category_name" form:"category_name"`
}