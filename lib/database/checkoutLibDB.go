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
	cartSearchRes := config.Db.Table("carts").Select("products.product_id, products.product_name, products.price, carts.quantity").Joins("left join products on carts.product_id = products.product_id").Where(`user_id = ?`, userId).Find(&cart)

	if cartSearchRes.Error != nil {
		return models.CheckoutAPI{}, cartSearchRes.Error
	}

	if cartSearchRes.RowsAffected == 0 {
		return models.CheckoutAPI{}, errors.New("No product found in the cart")
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

func AddCheckoutByUserId(payment *models.PaymentMethodAPI, userId int) (models.TransactionAPI, int64, error) {
	payment.PaymentMethodName = strings.ToLower(payment.PaymentMethodName)
	carts := []models.CartAPI{}

	findCartRes := config.Db.Table("carts").Select("products.product_id, products.product_name, products.price, carts.quantity").Joins("left join products on carts.product_id = products.product_id").Where(`user_id = ?`, userId).Find(&carts)

	if findCartRes.Error != nil {
		return models.TransactionAPI{}, findCartRes.RowsAffected, findCartRes.Error
	}

	if findCartRes.RowsAffected == 0 {
		return models.TransactionAPI{}, findCartRes.RowsAffected, errors.New("No products found in the cart. Add products first before checking out")
	}

	paymentChoice := models.PaymentMethodAPI{}
	paymentMethodCheck := config.Db.Table("payment_methods").Where("payment_method_name = ?", payment.PaymentMethodName).Find(&paymentChoice)

	if paymentMethodCheck.Error != nil {
		return  models.TransactionAPI{}, paymentMethodCheck.RowsAffected , paymentMethodCheck.Error
	}

	if paymentMethodCheck.RowsAffected == 0 {
		return  models.TransactionAPI{}, paymentMethodCheck.RowsAffected , errors.New("Payment method is not supported")
	}

	deletedCart := models.Cart{}

	deleteRes := config.Db.Table("carts").Where("user_id = ?", userId).Unscoped().Delete(&deletedCart)

	if deleteRes.Error != nil {
		return models.TransactionAPI{}, deleteRes.RowsAffected, deleteRes.Error
	}

	if deleteRes.RowsAffected == 0 {
		return models.TransactionAPI{}, deleteRes.RowsAffected, errors.New("Failed to delete items in user's cart.")
	}

	invoiceId := fmt.Sprintf("USER_%v:%v", userId, time.Now().String()[0:19])
	transactionDetail := models.TransactionDetail{}
	transactionDetail.InvoiceID = invoiceId
	var total uint

	for _, cartItem := range carts {
		transactionDetail.ProductID = cartItem.ProductID
		transactionDetail.ProductPrice = cartItem.Price
		transactionDetail.Quantity = cartItem.Quantity
		total += uint(cartItem.Price) * uint(cartItem.Quantity)

		transactionDetailCreation := config.Db.Create(&transactionDetail)

		if transactionDetailCreation.Error != nil {
			return models.TransactionAPI{}, transactionDetailCreation.RowsAffected, transactionDetailCreation.Error
		}

		if transactionDetailCreation.RowsAffected == 0 {
			return models.TransactionAPI{}, transactionDetailCreation.RowsAffected, errors.New("Failed to add transaction detail")
		}
	}

	transaction := models.Transaction{}
	transaction.UserID = uint(userId)
	transaction.InvoiceID = invoiceId
	transaction.PaymentMethodName = payment.PaymentMethodName

	transactionCreation := config.Db.Create(&transaction)

	if transactionCreation.Error != nil {
		return models.TransactionAPI{}, transactionCreation.RowsAffected , transactionCreation.Error
	}

	if transactionCreation.RowsAffected == 0{
		return models.TransactionAPI{}, transactionCreation.RowsAffected , errors.New("Failed to add transaction")
	}


	transactionAPI := models.TransactionAPI{}
	transactionAPI.InvoiceID = invoiceId
	transactionAPI.Total = uint(total)
	transactionAPI.PaymentMethodName = payment.PaymentMethodName
	transactionAPI.Description = paymentChoice.Description

	return transactionAPI, transactionCreation.RowsAffected, nil
}