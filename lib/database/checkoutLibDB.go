package libdb

import (
	"altastore/config"
	"altastore/models"
	"time"
	"fmt"
	"errors"
	"strings"
)

func GetCheckoutByUserId(userId int) (models.CheckoutAPI, error){
	cart := []models.CartAPI{}
	res := config.Db.Table("carts").Select("products.product_name, products.price, carts.quantity").Joins("left join products on carts.product_name = products.product_name").Where(`user_id = ?`, userId).Find(&cart)

	if res.Error != nil {
		return models.CheckoutAPI{}, res.Error
	}

	var total uint

	for _, cartItem := range cart {
		total += (uint(cartItem.Quantity) * uint(cartItem.Price))
	}

	var checkout models.CheckoutAPI

	checkout.UserID = uint(userId)
	checkout.Products = cart
	checkout.Total = total
	checkout.PaymentOptions = []string{"alfamart", "indomaret", "bank transfer", "gopay", "ovo", "link aja", "dana"}
	return checkout, nil
}

func AddCheckoutByUserId(paymentMethod string, userId int) (string, int64, error) {
	invoice := models.Invoice{}
	invoice.UserID = uint(userId)
	invoiceId := fmt.Sprintf("USER_%v:%v", userId, time.Now().String()[0:19])
	invoice.InvoiceID = invoiceId

	payment := models.Payment{}
	payment.PaymentMethod = paymentMethod
	paymentCheck := config.Db.Model(&models.Payment{}).Where("payment_method = ?", paymentMethod).Find(&payment)

	if paymentCheck.Error != nil {
		return  "", paymentCheck.RowsAffected , paymentCheck.Error
	}

	if paymentCheck.RowsAffected == 0 {
		return  "", paymentCheck.RowsAffected , errors.New("Payment is not supported")
	}

	invoice.PaymentMethod = strings.ToLower(paymentMethod)

	invoiceCreation := config.Db.Create(&invoice)

	if invoiceCreation.Error != nil || invoiceCreation.RowsAffected == 0{
		return "", invoiceCreation.RowsAffected , invoiceCreation.Error
	}

	return invoiceId, invoiceCreation.RowsAffected, nil
}