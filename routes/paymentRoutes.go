package routes

import(
	payment "altastore/controllers/payment"
	"altastore/middlewares"
)

func registerPaymentRoutes() {
	paymentGroup := e.Group("/payments")

	paymentGroup.Use(middlewares.AuthenticateUser)

	paymentGroup.GET("", payment.GetPendingPaymentsController)

	paymentGroup.POST("", payment.AddPendingPaymentController)
}

