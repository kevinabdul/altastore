package routes

import(
	checkout "altastore/controllers/checkout"
	"altastore/middlewares"
)

func registerCheckoutRoutes() {
	s := e.Group("/checkout")

	s.Use(middlewares.AuthenticateUser)

	s.GET("", checkout.GetCheckoutByUserIdController)

	s.POST("", checkout.AddCheckoutByUserIdController)
}

