package routes

import(
	checkout "altastore/controllers/checkout"
	"altastore/middlewares"
)

func registerCheckoutRoutes() {
	e.GET("/checkout", checkout.GetCheckoutByUserIdController, middlewares.AuthenticateUser)
}

