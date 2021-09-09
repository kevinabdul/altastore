package models

import (
	"time"
)

type CheckoutAPI struct {
	UserID			uint 		`gorm:"primaryKey;autoIncrement:false" json:"user_id"`
	Products		[]CartAPI 	`gorm:"primaryKey;not null;type:varchar(50)" json:"products"`
	Total			uint 		`gorm:"not null;type:smallint" json:"total"`
	Shippment		string
	Payment_method	string
}

type Invoice struct {
	UserID			uint		`gorm:"primaryKey;autoIncrement:false"`
	InvoiceId		string		`gorm:"primaryKey;not null;type:varchar(60)" json:"invoice_id"`
	Paid			string		`gorm:"type:enum('false', 'true');default:'false'" json:"paid"`
	CreatedAt 		time.Time
}

type InvoiceAPI struct {
	UserID			uint		`gorm:"primaryKey;autoIncrement:false" json:"user_id"`
	InvoiceId		string		`gorm:"primaryKey;not null;type:varchar(60)" json:"invoice_id"`
}
