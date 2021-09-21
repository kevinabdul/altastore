package models

import (
	"time"
)

type PaymentMethod struct {
	PaymentMethodID		uint 				`gorm:"primaryKey" json:"payment_method_id"`
	PaymentMethodName	string				`gorm:"unique;type:varchar(25)" json:"payment_method_name"`
	Description     	string 				`gorm:"type:varchar(1000)" json:"description"`
}

// Placeholder of information sent by the user when they do a checkout through post checkout endpoint
// Can be used for placeholder when finding about paymentmethod information which as of now only consist of its description
type PaymentMethodAPI struct {
	PaymentMethodID 	uint				`json:"payment_method_id"`
	Description     	string 				`json:"description"`
}

// Response struct to be returned for all transaction with pending payments status from a given user
// Used in GetPendingPaymentsByUserId
type PendingPaymentAPI struct {
	InvoiceID			string				`json:"invoice_id"`
	Total 				uint 				`json:"total"`
	PaymentMethodID 	uint  				`json:"payment_method_id"`
	Description   		string 				`json:"payment_method_description"` 
	Detail  			[]PaymentDetailAPI	`json:"detail"`
}

// Used to hold information about details of a transaction
// This is essentialy like the TransactionDetailAPI struct but instead of for query result placeholder, its used for response struct
type PaymentDetailAPI struct {
	InvoiceID			string				`json:"-"`
	ProductID	 		uint				`json:"product_id"`
	ProductName 		string				`json:"product_name"`
	ProductPrice		uint 				`json:"product_price"`
	Quantity  			uint  				`json:"quantity"`
	PaymentMethodID 	uint  				`json:"-"`
	Description   		string 				`json:"-"` 
}


// Placeholder of information sent by the user when they do a payment through post payments endpoint
type UserPaymentAPI struct {
	InvoiceID			string 				`json:"invoice_id"`
	Amount    			uint  				`json:"amount"`
	PaymentMethodID		uint 				`json:"payment_method_id"`
}


// Response struct returned to user after they completed a payment
type ReceiptAPI struct {
	UserID				string 				`json:"user_id"`
	InvoiceID			string 				`json:"invoice_id"`
	Amount    			uint  				`json:"amount"`
	PaymentMethodID		uint 				`json:"payment_method_id"`
	CreatedAt 			time.Time 			`json:"payment_date"`
}