package libdb

import (
	"altastore/config"
	"altastore/models"
)

func GetProducts(categoryName string) ([]models.ProductAPI, int64, error) {
	var products []models.ProductAPI

	var rowsAffected int64

	if categoryName == "" {
		prodSearchRes := config.Db.Table("products").Select("products.product_id, products.product_name, categories.category_name, products.price").Joins("left join categories on categories.category_id = products.category_id").Scan(&products)	

		if prodSearchRes.Error != nil ||  prodSearchRes.RowsAffected == 0 {
			return []models.ProductAPI{}, prodSearchRes.RowsAffected, prodSearchRes.Error
		}
		rowsAffected = prodSearchRes.RowsAffected
	} else {
		prodSearchRes := config.Db.Table("products").Select("products.product_id, products.product_name, categories.category_name, products.price").Joins("left join categories on categories.category_id = products.category_id").Where("categories.category_name = ?", categoryName).Scan(&products)	
		
		if prodSearchRes.Error != nil ||  prodSearchRes.RowsAffected == 0 {
			return []models.ProductAPI{}, prodSearchRes.RowsAffected, prodSearchRes.Error
		}		
		rowsAffected = prodSearchRes.RowsAffected
	}
	
	return products, rowsAffected, nil
}