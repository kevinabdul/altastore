package libdb

import (
	"altastore/config"
	"altastore/models"
	"errors"
)

func GetProducts(categoryName string) ([]models.ProductAPI, error) {
	var products []models.ProductAPI

	if categoryName == "" {
		prodSearchRes := config.Db.Table("products").Select("products.product_id, products.product_name, categories.category_name, products.price").Joins("left join categories on categories.category_id = products.category_id").Scan(&products)	
		
		if prodSearchRes.Error != nil {
			return []models.ProductAPI{}, prodSearchRes.Error
		}		

		if prodSearchRes.RowsAffected == 0 {
			return []models.ProductAPI{}, errors.New("No product found in the product table")
		}
	} else {
		prodSearchRes := config.Db.Table("products").Select("products.product_id, products.product_name, categories.category_name, products.price").Joins("left join categories on categories.category_id = products.category_id").Where("categories.category_name = ?", categoryName).Scan(&products)	
		
		if prodSearchRes.Error != nil {
			return []models.ProductAPI{}, prodSearchRes.Error
		}		

		if prodSearchRes.RowsAffected == 0 {
			return []models.ProductAPI{}, errors.New("No product found for the given cateogory")
		}	

	}
	
	return products, nil
}