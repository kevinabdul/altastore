package routes

import(
	checkout "altastore/controllers/checkout"
	"altastore/middlewares"
)

func registerCheckoutRoutes() {
	checkoutGroup := e.Group("/checkout")

	checkoutGroup.Use(middlewares.AuthenticateUser)

	checkoutGroup.GET("", checkout.GetCheckoutByUserIdController)

	checkoutGroup.POST("", checkout.AddCheckoutByUserIdController)
}

