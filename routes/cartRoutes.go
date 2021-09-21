package routes

import(
	cart "altastore/controllers/cart"
	"altastore/middlewares"
)

func registerCartRoutes() map[string][]interface{} {
	cartGroup := e.Group("/carts")

	cartGroup.Use(middlewares.AuthenticateUser)

	cartRoutesMap := map[string][]interface{}{}

	getCartByUserId := cartGroup.GET("", cart.GetCartByUserIdController)
	cartRoutesMap["GET"] = append(cartRoutesMap["GET"], getCartByUserId.Name)

	editcart := cartGroup.PUT("", cart.UpdateCartByUserIdController)
	cartRoutesMap["PUT"] = append(cartRoutesMap["GET"], editcart.Name)

	deleteCart := cartGroup.DELETE("", cart.DeleteCartByUserIdController)
	cartRoutesMap["DELETE"] = append(cartRoutesMap["GET"], deleteCart.Name)

	return cartRoutesMap
}

