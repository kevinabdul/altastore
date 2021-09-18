package libdb

import (
	"altastore/config"
	"altastore/models"
	"errors"
	"fmt"
	//"strings"
	"strconv"
)

func GetPendingPaymentsByUserId(userId int) ([]models.PendingPaymentAPI, error){
	invoiceMap := map[string]map[string][]models.PaymentDetailAPI{}

	paymentDetails := []models.PaymentDetailAPI{}

	pendingPaymentSearchRes := config.Db.Table("transactions").Select("transactions.invoice_id, transaction_details.product_id, products.product_name, transaction_details.product_price, transaction_details.quantity, transactions.payment_method_id, payment_methods.description").Joins("left join transaction_details on transactions.invoice_id = transaction_details.invoice_id").Joins("left join payment_methods on transactions.payment_method_id = payment_methods.payment_method_id").Joins("left join products on products.product_id = transaction_details.product_id").Where("user_id = ? and status = ?", userId, "pending payment").Find(&paymentDetails)

	if pendingPaymentSearchRes.Error != nil {
		return []models.PendingPaymentAPI{}, pendingPaymentSearchRes.Error
	}

	if pendingPaymentSearchRes.RowsAffected == 0 {
		return []models.PendingPaymentAPI{}, errors.New("No pending payments found")
	}

	for _, paymentDetail := range paymentDetails {
		if _, ok := invoiceMap[paymentDetail.InvoiceID]; !ok {
			newDetail := map[string][]models.PaymentDetailAPI{"detail": []models.PaymentDetailAPI{}}
			invoiceMap[paymentDetail.InvoiceID] = newDetail
			invoiceMap[paymentDetail.InvoiceID]["detail"] = append(invoiceMap[paymentDetail.InvoiceID]["detail"], paymentDetail)
		} else {
			invoiceMap[paymentDetail.InvoiceID]["detail"] = append(invoiceMap[paymentDetail.InvoiceID]["detail"], paymentDetail)
		}
	}
	
	res := []models.PendingPaymentAPI{}

	for id, invoice := range invoiceMap {
		result := models.PendingPaymentAPI{}
		result.InvoiceID = id
		result.PaymentMethodID = invoice["detail"][0].PaymentMethodID
		result.Description = invoice["detail"][0].Description
		result.Detail = invoice["detail"]

		for _, detail := range invoice["detail"] {
			result.Total += (detail.ProductPrice * detail.Quantity)
		}

		res = append(res, result)
	}


	return res, nil
}

func AddPendingPaymentByUserId(payment models.UserPaymentAPI, userId int) (models.ReceiptAPI, error) {
	transactionTarget := []models.ReceiptDetailAPI{}

	findPayment := config.Db.Table("transactions").Select("transactions.user_id, transactions.invoice_id, transactions.status, transaction_details.product_id, transaction_details.product_price, transaction_details.quantity, transactions.payment_method_id, payment_methods.payment_method_name, payment_methods.description").Joins("left join transaction_details on transactions.invoice_id = transaction_details.invoice_id").Joins("left join payment_methods on transactions.payment_method_id = payment_methods.payment_method_id").Where("transactions.user_id = ? and transactions.invoice_id = ?", userId, payment.InvoiceID).Find(&transactionTarget)

	if findPayment.Error != nil {
		return models.ReceiptAPI{}, findPayment.Error
	}

	if findPayment.RowsAffected == 0 {
		return models.ReceiptAPI{}, errors.New("No invoice_id found")
	}

	if transactionTarget[0].Status != "pending payment" {
		return models.ReceiptAPI{}, errors.New(fmt.Sprintf("Specified invoice id: %v has been %v", payment.InvoiceID ,transactionTarget[0].Status))	
	}

	if payment.PaymentMethodID != transactionTarget[0].PaymentMethodID {
		return models.ReceiptAPI{}, errors.New(fmt.Sprintf("Specified payment method doesnt match. Please pay using payment_method_id: %v", transactionTarget[0].PaymentMethodID))
	}

	var total uint

	for _, transactionDetail := range transactionTarget {
		total += (transactionDetail.ProductPrice * transactionDetail.Quantity)
	}

	if total != payment.Amount {
		return models.ReceiptAPI{}, errors.New(fmt.Sprintf("Specified amount doesnt match. Amount to be paid: %v", total))
	}

	newTransaction := models.Transaction{}

	updateRes := config.Db.Model(&newTransaction).Where("user_id = ? and invoice_id = ?", userId, payment.InvoiceID).Updates(models.Transaction{Status:"resolved", Paid:"true"})

	if updateRes.Error != nil {
		return models.ReceiptAPI{}, updateRes.Error
	}

	if updateRes.RowsAffected == 0 {
		return models.ReceiptAPI{}, errors.New(fmt.Sprintf("Failed to update invoice: %v", payment.InvoiceID))
	}

	receipt := models.ReceiptAPI{}

	receipt.UserID = strconv.Itoa(userId)
	receipt.InvoiceID = payment.InvoiceID
	receipt.Amount = payment.Amount
	receipt.PaymentMethodID = payment.PaymentMethodID
	receipt.CreatedAt = newTransaction.UpdatedAt

	return receipt, nil
}