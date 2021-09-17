package models

import (
	"time"
)

type User struct {
	UserID     	uint 		`json:"-" gorm:"primaryKey"`
	Name   		string		`gorm:"type:varchar(50); not null" json:"name" form:"name"`
	Email 		string		`gorm:"unique;type:varchar(50);not null" json:"email" form:"email"`
	Password 	string		`gorm:"type:varchar(100);not null" json:"password" form: "password"`
	CreatedAt 	time.Time
	UpdatedAt	time.Time
}

type UserAPI struct {
	UserID 		uint 		`json:"user_id"`
	Name 		string 		`json:"name"`
	Email 		string 		`json:"email"`
}