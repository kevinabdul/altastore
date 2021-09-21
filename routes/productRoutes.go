package routes

import(
	product "altastore/controllers/product"
)

func registerProductRoutes() map[string][]interface{} {
	productRoutesMap := map[string][]interface{}{}

	getProduct := e.GET("/products", product.GetProductsController)
	productRoutesMap["GET"] = append(productRoutesMap["GET"], getProduct.Name)

	return productRoutesMap
}

