package libdb

import (
	"errors"
	"fmt"

	"altastore/config"
	"altastore/models"
)

func GetCartByUserId(userId int) ([]models.CartAPI, error) {
	var cart []models.CartAPI

	res := config.Db.Table("carts").Select("products.product_name, products.price, carts.quantity").Joins("left join products on carts.product_name = products.product_name").Where(`user_id = ?`, userId).Find(&cart)

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
			return prodSearchRes.RowsAffected, errors.New(fmt.Sprintf("No product named %s found in the product list", cartItem.ProductName))
		}
		
		cartItemSearchRes := config.Db.Where(`user_id = ? AND product_name = ?`, userId, cartItem.ProductName).Find(&cart)

		if cartItemSearchRes.Error != nil {
			return cartItemSearchRes.RowsAffected, cartItemSearchRes.Error
		} else if cartItemSearchRes.RowsAffected == 0 {
			insertRes := config.Db.Select("UserID", "ProductName", "Quantity").Create(&cartItem)

			if insertRes.Error != nil {
				return insertRes.RowsAffected, insertRes.Error
			}

			if insertRes.RowsAffected == 0 {
				return insertRes.RowsAffected, errors.New(fmt.Sprintf("Failed to add item %s to user's cart", cartItem.ProductName))
			}

			rowsAffected++
		} else if cartItemSearchRes.RowsAffected != 0 && cart.Quantity != cartItem.Quantity {
			updateRes := config.Db.Model(&cart).Select("quantity").Updates(cartItem)
			
			if updateRes.Error != nil {
				return updateRes.RowsAffected, updateRes.Error
			}

			if updateRes.RowsAffected == 0{
				return updateRes.RowsAffected, errors.New(fmt.Sprintf("Failed to update item %s in user's cart", cartItem.ProductName))
			}

			rowsAffected++
		}
	}

	return rowsAffected, nil
}

func DeleteCartByUserId(items []string, userId int) (int, error) {
	if len(items) == 0 {
		return 0, errors.New("No item found in delete list. Please specify before deleting")
	}

	deletedCart := models.Cart{}
	deletedItem := 0

	for _, item := range items {
		deleteRes := config.Db.Table("carts").Where("user_id = ? and product_name = ?", userId, item).Unscoped().Delete(&deletedCart)
		
		if deleteRes.Error != nil {
			return 0, deleteRes.Error
		}

		if deleteRes.RowsAffected == 0 {
			return 0, errors.New(fmt.Sprintf("No item named %s is found in user's cart.", item))
		}

		deletedItem++
	}

	return deletedItem, nil
}