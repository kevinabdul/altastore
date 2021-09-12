package routes

import(
	cart "altastore/controllers/cart"
	"altastore/middlewares"
)

func registerCartRoutes() {
	cartGroup := e.Group("/carts/:id")

	cartGroup.Use(middlewares.AuthenticateUser)

	cartGroup.Use(middlewares.CheckId)
	
	cartGroup.GET("", cart.GetCartByUserIdController)

	cartGroup.PUT("", cart.UpdateCartByUserIdController)

	cartGroup.DELETE("", cart.DeleteCartByUserIdController)
}

