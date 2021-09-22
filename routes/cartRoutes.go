package routes

import(
	handler "altastore/controllers"
	"altastore/middlewares"
)

func registerCartRoutes() map[string][]interface{} {
	cartGroup := e.Group("/carts")

	cartGroup.Use(middlewares.AuthenticateUser)

	cartRoutesMap := map[string][]interface{}{}

	getCartByUserId := cartGroup.GET("", handler.GetCartByUserIdController)
	cartRoutesMap["GET"] = append(cartRoutesMap["GET"], getCartByUserId.Name)

	editcart := cartGroup.PUT("", handler.UpdateCartByUserIdController)
	cartRoutesMap["PUT"] = append(cartRoutesMap["GET"], editcart.Name)

	deleteCart := cartGroup.DELETE("", handler.DeleteCartByUserIdController)
	cartRoutesMap["DELETE"] = append(cartRoutesMap["GET"], deleteCart.Name)

	return cartRoutesMap
}

