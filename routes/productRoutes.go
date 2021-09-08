package routes

import(
	product "altastore/controllers/product"
)

func registerProductRoutes() {
	e.GET("/products", product.GetProductsController)
}

