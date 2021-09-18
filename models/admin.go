package models

import (
	"time"
)

type Admin struct {
	UserID     	uint 		`json:"user_id" gorm:"primaryKey"`
	CreatedAt 	time.Time
	UpdatedAt	time.Time
	User        User    	`gorm:"foreignKey:"UserID"`
}