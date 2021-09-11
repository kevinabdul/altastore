package models

import (
	"time"
)

type CheckoutAPI struct {
	UserID				uint 		`gorm:"primaryKey;autoIncrement:false" json:"user_id"`
	Products			[]CartAPI 	`gorm:"primaryKey;not null;type:varchar(50)" json:"products"`
	Total				uint 		`gorm:"not null;type:smallint" json:"total"`
	// Shipment			string		`gorm:"type:enum('jne', 'pos', 'tiki');default:'jne'" json:"shipment"`
	PaymentOptions		[]string	`json:"payment_options"`
}

type Transaction struct {
	UserID				uint		`gorm:"primaryKey;autoIncrement:false"`
	InvoiceID			string		`gorm:"primaryKey;not null;type:varchar(60)" json:"invoice_id"`
	Paid				string		`gorm:"type:enum('false', 'true');default:'false'" json:"paid"`
	PaymentMethodName	string		`gorm:"type:varchar(25)" json:"payment_method_name"`
	Status				string		`gorm:"type:enum('pending', 'cancelled', 'expired', 'resolved');default:'pending'" json:"status"`
	CreatedAt 			time.Time
	UpdatedAt			time.Time
}

type TransactionDetail struct {
	InvoiceID			string		`gorm:"primaryKey;not null;type:varchar(60)" json:"invoice_id"`
	ProductName 		string		`gorm:"primaryKey;type:varchar(100)" json:"product_name"`	
	CreatedAt 			time.Time
	UpdatedAt			time.Time
}

type PaymentMethod struct {
	PaymentMethodID		uint 		`gorm:"primaryKey" json:"payment_method_id"`
	PaymentMethodName	string		`gorm:"unique;type:varchar(25)" json:"payment_method_name"`
}

type PaymentMethodAPI struct {
	PaymentMethodName 	string		`json:"payment_method_name"`
}
