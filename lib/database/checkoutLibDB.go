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
	cartSearchRes := config.Db.Table("carts").Select("products.product_name, products.price, carts.quantity").Joins("left join products on carts.product_name = products.product_name").Where(`user_id = ?`, userId).Find(&cart)

	if cartSearchRes.Error != nil {
		return models.CheckoutAPI{}, cartSearchRes.Error
	}

	if cartSearchRes.RowsAffected == 0 {
		return models.CheckoutAPI{}, errors.New("No item found in the cart")
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
	count := 0

	config.Db.Model(&models.Cart{}).Select("count(user_id)").Where("user_id = ?", userId).Find(&count)

	if count == 0 {
		return "", int64(0), errors.New("No products found in the cart. Add products first before checking out")
	}

	payment := models.Payment{}
	payment.PaymentMethod = paymentMethod

	paymentCheck := config.Db.Model(&models.Payment{}).Where("payment_method = ?", paymentMethod).Find(&payment)

	if paymentCheck.Error != nil {
		return  "", paymentCheck.RowsAffected , paymentCheck.Error
	}

	if paymentCheck.RowsAffected == 0 {
		return  "", paymentCheck.RowsAffected , errors.New("Payment is not supported")
	}

	carts := []models.Cart{}

	findCartRes := config.Db.Where("user_id = ?", userId).Find(&carts)

	if findCartRes.Error != nil || findCartRes.RowsAffected == 0 {
		return "", findCartRes.RowsAffected, findCartRes.Error
	}

	deletedCart := models.Cart{}
	deleteRes := config.Db.Where("user_id = ?", userId).Unscoped().Delete(&deletedCart)
	fmt.Println(deletedCart)

	if deleteRes.Error != nil {
		return "", deleteRes.RowsAffected, deleteRes.Error
	}

	if deleteRes.RowsAffected == 0 {
		return "", deleteRes.RowsAffected, errors.New("Failed to delete user.")
	}

	invoice := models.Invoice{}
	invoice.UserID = uint(userId)
	invoiceId := fmt.Sprintf("USER_%v:%v", userId, time.Now().String()[0:19])
	invoice.InvoiceID = invoiceId
	invoice.PaymentMethod = strings.ToLower(paymentMethod)

	invoiceCreation := config.Db.Create(&invoice)

	if invoiceCreation.Error != nil || invoiceCreation.RowsAffected == 0{
		return "", invoiceCreation.RowsAffected , invoiceCreation.Error
	}

	return invoiceId, invoiceCreation.RowsAffected, nil
}