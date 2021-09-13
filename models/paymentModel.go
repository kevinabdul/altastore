package models

import (
	"time"
)

type PaymentMethod struct {
	PaymentMethodID		uint 				`gorm:"primaryKey" json:"payment_method_id"`
	PaymentMethodName	string				`gorm:"unique;type:varchar(25)" json:"payment_method_name"`
	Description     	string 				`gorm:"type:varchar(1000)" json:"description"`
}

type PaymentMethodAPI struct {
	PaymentMethodName 	string				`json:"payment_method_name"`
	Description     	string 				`gorm:"type:varchar(1000)" json:"description"`
}

type PendingPaymentAPI struct {
	InvoiceID			string				`json:"invoice_id"`
	Total 				uint 				`json:"total"`
	PaymentMethodName 	string  			`json:"payment_method_name"`
	Description   		string 				`json:"payment_method_description"` 
	Detail  			[]PaymentDetailAPI	`json:"detail"`
}

type PaymentDetailAPI struct {
	InvoiceID			string				`json:"-"`
	ProductName 		string				`json:"product_name"`
	ProductPrice		uint 				`json:"product_price"`
	Quantity  			uint  				`json:"quantity"`
	PaymentMethodName 	string  			`json:"-"`
	Description   		string 				`json:"-"` 
}

type UserPaymentAPI struct {
	InvoiceID			string 				`json:"invoice_id"`
	Amount    			uint  				`json:"amount"`
	PaymentMethodName	string 				`json:"payment_method_name"`
}

type ReceiptAPI struct {
	UserID				string 				`json:"user_id"`
	InvoiceID			string 				`json:"invoice_id"`
	Amount    			uint  				`json:"amount"`
	PaymentMethodName	string 				`json:"payment_method_name"`
	UpdatedAt 			time.Time 			`json:"payment_date"`
}

type ReceiptDetailAPI struct {
	UserID				string 				`json:"user_id"`
	InvoiceID			string				`json:"invoice_id"`
	Status   			string  			`json:"status"`
	ProductName 		string				`json:"product_name"`
	ProductPrice		uint 				`json:"product_price"`
	Quantity  			uint  				`json:"quantity"`
	PaymentMethodName 	string  			`json:"payment_method_name"`
	Description   		string 				`json:"description"` 
}