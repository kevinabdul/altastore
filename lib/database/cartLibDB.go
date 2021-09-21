package libdb

import (
	"errors"
	"fmt"
	"strings"
	"gorm.io/gorm"

	"altastore/config"
	"altastore/models"
)

func GetCartByUserId(userId int) ([]models.CartAPI, error) {
	var cart []models.CartAPI

	res := config.Db.Table("carts").Select("products.product_id, products.product_name, products.price, carts.quantity").Joins("left join products on carts.product_id = products.product_id").Where(`user_id = ?`, userId).Find(&cart)

	if res.Error != nil {
		return []models.CartAPI{}, res.Error
	}

	if res.RowsAffected == 0 {
		return []models.CartAPI{}, errors.New("No product found in the cart")
	}

	return cart, nil
}

// This function assumes the userId is still exist. That check should be handled by another auth functionality, not by this function.
func UpdateCartByUserId(userCart []models.Cart, userId int)  error {
	err := config.Db.Transaction(func(tx *gorm.DB) error {
		for _, cartItem := range userCart {
			// Request body binding done in the controller should already "convert" any integer less than zero to zero
			if cartItem.Quantity == 0 || cartItem.ProductID == 0 {
				continue
			}

			targetCart := models.Cart{}

			/* Just found about this awesome and convenient method the night before presentation */
			res := tx.Where(models.Cart{UserID: uint(userId), ProductID: cartItem.ProductID}).Assign(models.Cart{Quantity: cartItem.Quantity}).FirstOrCreate(&targetCart)

			if res.Error != nil {
				// Error 1452 means we try to change a child table with invalid parent's table primary key
				if strings.HasPrefix(res.Error.Error(), "Error 1452") {
					return errors.New(fmt.Sprintf("No product id %v found in the product table", cartItem.ProductID))
				}

				return res.Error
			}
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func DeleteCartByUserId(itemIds []int, userId int) (error) {
	if len(itemIds) == 0 {
		return errors.New("No item found in delete list. Please specify before deleting")
	}

	deletedCart := models.Cart{}

	err := config.Db.Transaction(func(tx *gorm.DB) error {
		for _, itemId := range itemIds {
			deleteRes := tx.Table("carts").Where("user_id = ? and product_id = ?", userId, itemId).Unscoped().Delete(&deletedCart)
			
			if deleteRes.Error != nil {
				return deleteRes.Error
			}

			if deleteRes.RowsAffected == 0 {
				return errors.New(fmt.Sprintf("No product with id %v is found in user's cart.", itemId))
			}
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}