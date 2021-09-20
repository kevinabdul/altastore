package libdb

import (
	"altastore/config"
	"altastore/models"
	"time"
	"fmt"
	"errors"
	"strings"
	"gorm.io/gorm"
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

func AddCheckoutByUserId(payment *models.PaymentMethodAPI, userId int) (models.TransactionAPI, error) {
	carts := []models.CartAPI{}

	findCartRes := config.Db.Table("carts").Select("products.product_id, products.product_name, products.price, carts.quantity").Joins("left join products on carts.product_id = products.product_id").Where(`user_id = ?`, userId).Find(&carts)

	if findCartRes.Error != nil {
		return models.TransactionAPI{}, findCartRes.Error
	}

	if findCartRes.RowsAffected == 0 {
		return models.TransactionAPI{}, errors.New("No products found in the cart. Add products first before checking out")
	}

	/* transactionAPI, transactionCreation, transaction, and transactionDetail here means:
	   the display struct used for response, query result from gorm method, model, and model,
	   not the mysql transaction statement executed in the function below */
	transactionAPI := models.TransactionAPI{}
	paymentChoiceAPI := models.PaymentMethodAPI{}

	var transactionCreation *gorm.DB

	err := config.Db.Transaction(func(tx *gorm.DB) error {

		deletedCart := models.Cart{}

		deleteRes := tx.Table("carts").Where("user_id = ?", userId).Unscoped().Delete(&deletedCart)

		if deleteRes.Error != nil {
			return deleteRes.Error
		}

		if deleteRes.RowsAffected == 0 {
			return errors.New("Failed to delete items in user's cart.")
		}

		invoiceId := fmt.Sprintf("USER_%v:%v", userId, time.Now().String()[0:19])

		transaction := models.Transaction{}
		transaction.UserID = uint(userId)
		transaction.InvoiceID = invoiceId
		transaction.PaymentMethodID = payment.PaymentMethodID

		transactionCreation = tx.Create(&transaction)

		if transactionCreation.Error != nil {
			if strings.HasPrefix(transactionCreation.Error.Error(), "Error 1452") {
					return errors.New(fmt.Sprintf("No payment_method_id '%v' found in the payment method table", payment.PaymentMethodID))
				}
			return transactionCreation.Error
		}

		if transactionCreation.RowsAffected == 0 {
			return errors.New("Failed to add transaction")
		}

		transactionDetail := models.TransactionDetail{}
		transactionDetail.InvoiceID = invoiceId
		var total uint

		for _, cartItem := range carts {
			transactionDetail.ProductID = cartItem.ProductID
			transactionDetail.ProductPrice = cartItem.Price
			transactionDetail.Quantity = cartItem.Quantity
			total += uint(cartItem.Price) * uint(cartItem.Quantity)

			transactionDetailCreation := tx.Create(&transactionDetail)

			if transactionDetailCreation.Error != nil {
				if strings.HasPrefix(transactionDetailCreation.Error.Error(), "Error 1452") {
					return errors.New(fmt.Sprintf("No invoice id '%v' found in the transaction table", transactionDetail.InvoiceID))
				}
				return transactionDetailCreation.Error
			}

			if transactionDetailCreation.RowsAffected == 0 {
				return errors.New("Failed to add transaction detail")
			}
		}

		paymentMethodCheck := config.Db.Table("payment_methods").Where("payment_method_id = ?", payment.PaymentMethodID).Find(&paymentChoiceAPI)

		if paymentMethodCheck.Error != nil {
			return  paymentMethodCheck.Error
		}

		if paymentMethodCheck.RowsAffected == 0 {
			return  errors.New("Payment method is not supported")
		}

		transactionAPI.InvoiceID = invoiceId
		transactionAPI.Total = uint(total)
		transactionAPI.PaymentMethodID = payment.PaymentMethodID
		transactionAPI.Description = paymentChoiceAPI.Description

		return nil
	})

	if err != nil {
		if transactionCreation.RowsAffected == 0 {
			return transactionAPI, errors.New("Failed to create invoice")
		}
		return transactionAPI, err
	}

	return transactionAPI, nil
}