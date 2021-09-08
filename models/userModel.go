package models

import (
	"time"
)

type User struct {
	ID     		uint 		`json:"-" gorm:"primaryKey`
	Name   		string		`gorm:"type:varchar(50)" json:"name" form:"name"`
	Email 		string		`gorm:"unique;type:varchar(50);not null" json:"email" form:"email"`
	Password 	string		`gorm:"type:varchar(30);not null" json:"password" form: "password"`
	CreatedAt 	time.Time
	UpdatedAt	time.Time
}

type UserAPI struct {
	ID 			uint 		`json:"id"`
	Name 		string 		`json:"name"`
	Email 		string 		`json:"email"`
}