package routes

import(
	handler "altastore/controllers"
	"altastore/middlewares"
)

func registerCheckoutRoutes() map[string][]interface{}{
	checkoutGroup := e.Group("/checkout")

	checkoutGroup.Use(middlewares.AuthenticateUser)

	checkoutMap := map[string][]interface{}{}

	checkoutGroup.GET("", handler.GetCheckoutByUserIdController)
	getCheckout := checkoutGroup.GET("", handler.GetCheckoutByUserIdController)
	checkoutMap["GET"] = append(checkoutMap["GET"], getCheckout.Name)

	checkoutGroup.POST("", handler.AddCheckoutByUserIdController)
	postCheckout := checkoutGroup.POST("", handler.AddCheckoutByUserIdController)
	checkoutMap["POST"] = append(checkoutMap["POST"], postCheckout.Name)	

	return checkoutMap
}
