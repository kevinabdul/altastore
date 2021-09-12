package models

import (
	"time"
)

type CheckoutAPI struct {
	UserID				uint 		`json:"user_id"`
	Products			[]CartAPI 	`json:"products"`
	Total				uint 		`json:"total"`
	// Shipment			string		`gorm:"type:enum('jne', 'pos', 'tiki');default:'jne'" json:"shipment"`
	PaymentOptions		[]string	`json:"payment_options"`
}

type Transaction struct {
	UserID				uint		`gorm:"primaryKey;autoIncrement:false"`
	InvoiceID			string		`gorm:"primaryKey;not null;type:varchar(60)" json:"invoice_id"`
	Status				string		`gorm:"type:enum('pending payment', 'cancelled', 'expired', 'resolved');default:'pending payment'" json:"status"`
	Paid				string		`gorm:"type:enum('false', 'true');default:'false'" json:"paid"`
	PaymentMethodName	string		`gorm:"type:varchar(25)" json:"payment_method_name"`
	CreatedAt 			time.Time
	UpdatedAt			time.Time
}

type TransactionDetail struct {
	InvoiceID			string		`gorm:"primaryKey;not null;type:varchar(60)" json:"invoice_id"`
	ProductName 		string		`gorm:"primaryKey;type:varchar(100)" json:"product_name"`
	ProductPrice		uint 		`gorm:"type:uint" json:"product_price"`	
	CreatedAt 			time.Time
	UpdatedAt			time.Time
}

type TransactionAPI struct {
	InvoiceID			string		`json:"invoice_id"`
	Total 				uint 		`json:"total"`
	PaymentMethodName 	string		`json:"payment_method_name"`
	Description     	string 		`json:"description"` 		
}

type PaymentMethod struct {
	PaymentMethodID		uint 		`gorm:"primaryKey" json:"payment_method_id"`
	PaymentMethodName	string		`gorm:"unique;type:varchar(25)" json:"payment_method_name"`
	Description     	string 		`gorm:"type:varchar(1000)" json:"description"`
}

type PaymentMethodAPI struct {
	PaymentMethodName 	string		`json:"payment_method_name"`
	Description     	string 		`gorm:"type:varchar(1000)" json:"description"`
}
