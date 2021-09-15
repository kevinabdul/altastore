package routes

import(
	cart "altastore/controllers/cart"
	"altastore/middlewares"
)

func registerCartRoutes() {
	cartGroup := e.Group("/carts")

	cartGroup.Use(middlewares.AuthenticateUser)
	
	cartGroup.GET("", cart.GetCartByUserIdController)

	cartGroup.PUT("", cart.UpdateCartByUserIdController)

	cartGroup.DELETE("", cart.DeleteCartByUserIdController)
}

