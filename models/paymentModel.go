package models

type PaymentMethod struct {
	PaymentMethodID		uint 		`gorm:"primaryKey" json:"payment_method_id"`
	PaymentMethodName	string		`gorm:"unique;type:varchar(25)" json:"payment_method_name"`
	Description     	string 		`gorm:"type:varchar(1000)" json:"description"`
}

type PaymentMethodAPI struct {
	PaymentMethodName 	string		`json:"payment_method_name"`
	Description     	string 		`gorm:"type:varchar(1000)" json:"description"`
}