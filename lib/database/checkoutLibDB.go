package libdb

import (
	"altastore/config"
	"altastore/models"
	"time"
	"fmt"
	"errors"
	//"strings"
)

func GetCheckoutByUserId(userId int) (models.CheckoutAPI, error){
	cart := []models.CartAPI{}
	cartSearchRes := config.Db.Table("carts").Select("products.product_name, products.price, carts.quantity").Joins("left join products on carts.product_name = products.product_name").Where(`user_id = ?`, userId).Find(&cart)

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

func AddCheckoutByUserId(payment *models.PaymentMethodAPI, userId int) (string, int64, error) {
	carts := []models.Cart{}

	findCartRes := config.Db.Table("carts").Where("user_id = ?", userId).Find(&carts)

	if findCartRes.Error != nil {
		return "", findCartRes.RowsAffected, findCartRes.Error
	}

	if findCartRes.RowsAffected == 0 {
		return "", int64(0), errors.New("No products found in the cart. Add products first before checking out")
	}

	paymentMethodCheck := config.Db.Table("payment_methods").Where("payment_method_name = ?", payment.PaymentMethodName).Find(payment)

	if paymentMethodCheck.Error != nil {
		return  "", paymentMethodCheck.RowsAffected , paymentMethodCheck.Error
	}

	if paymentMethodCheck.RowsAffected == 0 {
		return  "", paymentMethodCheck.RowsAffected , errors.New("Payment method is not supported")
	}

	deletedCart := models.Cart{}

	deleteRes := config.Db.Table("carts").Where("user_id = ?", userId).Unscoped().Delete(&deletedCart)

	if deleteRes.Error != nil {
		return "", deleteRes.RowsAffected, deleteRes.Error
	}

	if deleteRes.RowsAffected == 0 {
		return "", deleteRes.RowsAffected, errors.New("Failed to delete user.")
	}

	transaction := models.Transaction{}
	transaction.UserID = uint(userId)
	invoiceId := fmt.Sprintf("USER_%v:%v", userId, time.Now().String()[0:19])
	transaction.InvoiceID = invoiceId
	transaction.PaymentMethodName = payment.PaymentMethodName

	transactionCreation := config.Db.Create(&transaction)

	if transactionCreation.Error != nil {
		return "", transactionCreation.RowsAffected , transactionCreation.Error
	}

	if transactionCreation.RowsAffected == 0{
		return "", transactionCreation.RowsAffected , errors.New("Failed to add transaction")
	}

	transactionDetail := models.TransactionDetail{}
	transactionDetail.InvoiceID = invoiceId

	for _, cartItem := range carts {
		transactionDetail.ProductName = cartItem.ProductName
		transactionDetailCreation := config.Db.Create(&transactionDetail)

		if transactionDetailCreation.Error != nil {
			return "", transactionDetailCreation.RowsAffected, transactionDetailCreation.Error
		}

		if transactionDetailCreation.RowsAffected == 0 {
			return "", transactionDetailCreation.RowsAffected, errors.New("Failed to add transaction detail")
		}
	}

	return invoiceId, transactionCreation.RowsAffected, nil
}