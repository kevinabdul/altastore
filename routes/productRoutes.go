package routes

import(
	handler "altastore/controllers"
)

func registerProductRoutes() map[string][]interface{} {
	productRoutesMap := map[string][]interface{}{}

	getProduct := e.GET("/products", handler.GetProductsController)
	productRoutesMap["GET"] = append(productRoutesMap["GET"], getProduct.Name)

	return productRoutesMap
}

