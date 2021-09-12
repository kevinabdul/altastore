package models

type CheckoutAPI struct {
	UserID				uint 		`json:"user_id"`
	Products			[]CartAPI 	`json:"products"`
	Total				uint 		`json:"total"`
	// Shipment			string		`gorm:"type:enum('jne', 'pos', 'tiki');default:'jne'" json:"shipment"`
	PaymentOptions		[]string	`json:"payment_options"`
}
