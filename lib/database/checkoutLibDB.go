package libdb

import (
	"altastore/config"
	"altastore/models"
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

	return checkout, nil
}

// func UpdateCartByUserId(userCart []models.Cart, userId int)  (int64, error) {
// 	var rowsAffected int64

// 	for _, cartItem := range userCart {
// 		if cartItem.Quantity == 0 || cartItem.ProductName == ""{
// 			continue
// 		}
		
// 		cartItem.UserID = uint(userId)
		
// 		cart := models.Cart{}
		
// 		productTarget := models.Product{}
// 		prodSearchRes := config.Db.Where(`product_name = ?`, cartItem.ProductName).Find(&productTarget)

// 		if prodSearchRes.Error != nil {
// 			return prodSearchRes.RowsAffected, prodSearchRes.Error
// 		}

// 		if prodSearchRes.RowsAffected == 0 {
// 			continue
// 		}
		
// 		cartItemSearchRes := config.Db.Where(`user_id = ? AND product_name = ?`, userId, cartItem.ProductName).Find(&cart)

// 		if cartItemSearchRes.Error != nil {
// 			return cartItemSearchRes.RowsAffected, cartItemSearchRes.Error
// 		} else if cartItemSearchRes.RowsAffected == 0 {
// 			insertRes := config.Db.Select("UserID", "ProductName", "Quantity").Create(&cartItem)

// 			if insertRes.Error != nil || insertRes.RowsAffected == 0 {
// 				return insertRes.RowsAffected, insertRes.Error
// 			}

// 			rowsAffected++
// 		} else if cartItemSearchRes.RowsAffected != 0 && cart.Quantity != cartItem.Quantity{
// 			updateRes := config.Db.Model(&cart).Select("quantity").Updates(cartItem)
			
// 			if updateRes.Error != nil || updateRes.RowsAffected == 0{
// 				return updateRes.RowsAffected, updateRes.Error
// 			}

// 			rowsAffected++
// 		}
// 	}

// 	return rowsAffected, nil
// }