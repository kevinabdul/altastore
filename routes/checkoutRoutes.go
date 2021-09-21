package routes

import(
	checkout "altastore/controllers/checkout"
	"altastore/middlewares"
)

func registerCheckoutRoutes() map[string][]interface{}{
	checkoutGroup := e.Group("/checkout")

	checkoutGroup.Use(middlewares.AuthenticateUser)

	checkoutMap := map[string][]interface{}{}

	checkoutGroup.GET("", checkout.GetCheckoutByUserIdController)
	getCheckout := checkoutGroup.GET("", checkout.GetCheckoutByUserIdController)
	checkoutMap["GET"] = append(checkoutMap["GET"], getCheckout.Name)

	checkoutGroup.POST("", checkout.AddCheckoutByUserIdController)
	postCheckout := checkoutGroup.POST("", checkout.AddCheckoutByUserIdController)
	checkoutMap["POST"] = append(checkoutMap["POST"], postCheckout.Name)	

	return checkoutMap
}

