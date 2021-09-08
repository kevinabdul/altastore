package services

import (
	"altastore/config"
	"altastore/models"
)

func GetCartByUserId(userId int) ([]models.CartAPI, error) {
	var cart []models.CartAPI

	res := config.Db.Model(&models.Cart{}).Where(`user_id = ?`, userId).Find(&cart)

	if res.Error != nil {
		return []models.CartAPI{}, res.Error
	}

	return cart, nil
}

func UpdateCartByUserId(userCart []models.Cart, userId int)  (int64, error) {
	var rowsAffected int64

	for _, cartItem := range userCart {
		if cartItem.Quantity == 0 || cartItem.ProductName == ""{
			continue
		}
		
		cartItem.UserID = uint(userId)
		
		cart := models.Cart{}
		
		productTarget := models.Product{}
		prodSearchRes := config.Db.Where(`product_name = ?`, cartItem.ProductName).Find(&productTarget)

		if prodSearchRes.Error != nil {
			return prodSearchRes.RowsAffected, prodSearchRes.Error
		}

		if prodSearchRes.RowsAffected == 0 {
			continue
		}
		
		cartItemSearchRes := config.Db.Where(`user_id = ? AND product_name = ?`, userId, cartItem.ProductName).Find(&cart)

		if cartItemSearchRes.Error != nil {
			return cartItemSearchRes.RowsAffected, cartItemSearchRes.Error
		} else if cartItemSearchRes.RowsAffected == 0 {
			insertRes := config.Db.Select("UserID", "ProductName", "Quantity").Create(&cartItem)

			if insertRes.Error != nil || insertRes.RowsAffected == 0 {
				return insertRes.RowsAffected, insertRes.Error
			}

			rowsAffected++
		} else if cartItemSearchRes.RowsAffected != 0 && cart.Quantity != cartItem.Quantity{
			updateRes := config.Db.Model(&cart).Select("quantity").Updates(cartItem)
			
			if updateRes.Error != nil || updateRes.RowsAffected == 0{
				return updateRes.RowsAffected, updateRes.Error
			}

			rowsAffected++
		}
	}

	return rowsAffected, nil
}