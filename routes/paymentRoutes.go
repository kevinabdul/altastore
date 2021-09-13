package routes

import(
	payment "altastore/controllers/payment"
	"altastore/middlewares"
)

func registerPaymentRoutes() {
	paymentGroup := e.Group("/payments/:id")

	paymentGroup.Use(middlewares.AuthenticateUser)

	paymentGroup.Use(middlewares.CheckId)

	paymentGroup.GET("", payment.GetPendingPaymentsByUserIdController)

	paymentGroup.POST("", payment.AddPendingPaymentByUserIdController)
}

