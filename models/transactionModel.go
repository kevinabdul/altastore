package models

import (
	"time"
)

type Transaction struct {
	UserID				uint	
	InvoiceID			string		`gorm:"primaryKey;not null;type:varchar(60)" json:"invoice_id"`
	Status				string		`gorm:"type:enum('pending payment', 'cancelled', 'expired', 'resolved');default:'pending payment'" json:"status"`
	Paid				string		`gorm:"type:enum('false', 'true');default:'false'" json:"paid"`
	PaymentMethodID		uint		`gorm:"type:uint" json:"payment_method_id"`
	CreatedAt 			time.Time
	UpdatedAt			time.Time
	User 				User  			`gorm:"foreignKey:UserID"`
	PaymentMethod   	PaymentMethod 	`gorm:"foreignKey:PaymentMethodID"`
}

type TransactionDetail struct {
	InvoiceID			string		`gorm:"primaryKey;not null;type:varchar(60)" json:"invoice_id"`
	ProductID			uint		`gorm:"primaryKey;type:uint" json:"product_id"`
	ProductPrice		uint 		`gorm:"type:uint" json:"product_price"`
	Quantity  			uint  		`gorm:"not null;type:smallint" json:"quantity"`
	CreatedAt 			time.Time
	UpdatedAt			time.Time
	Transaction   		Transaction `gorm:"foreignKey:InvoiceID"`
	Product   			Product  	`gorm:"foreignKey:ProductID"`
}

// Response struct used in case of a succesful checkout in post checkout endpoint
// Succesful checkout means we are able to delete data from carts table, creating new data in transactions table,
// and moving the deleted data into transaction_details table. Any failure in those step will fail whole transaction
type TransactionAPI struct {
	InvoiceID			string		`json:"invoice_id"`
	Total 				uint 		`json:"total"`
	PaymentMethodID 	uint		`json:"payment_method_id"`
	Description     	string 		`json:"description"` 		
}

// Commonly used when user try to do a payment. 
// AddPaymentByUserId will try to find corresponding transaction in a database based on UserId and information provided in UserPaymentAPI struct.
// This struct will be used as a placeholder of above query result. 
type TransactionDetailAPI struct {
	UserID				string 				`json:"user_id"`
	InvoiceID			string				`json:"invoice_id"`
	Status   			string  			`json:"status"`
	ProductName 		string				`json:"product_name"`
	ProductPrice		uint 				`json:"product_price"`
	Quantity  			uint  				`json:"quantity"`
	PaymentMethodID 	uint  				`json:"payment_method_id"`
	Description   		string 				`json:"description"` 
}