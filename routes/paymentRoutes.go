package routes

import(
	payment "altastore/controllers/payment"
	"altastore/middlewares"
)

func registerPaymentRoutes() map[string][]interface{} {
	paymentGroup := e.Group("/payments")

	paymentGroup.Use(middlewares.AuthenticateUser)

	paymentRoutesMap := map[string][]interface{}{}

	getPayment := paymentGroup.GET("", payment.GetPendingPaymentsController)
	paymentRoutesMap["GET"] = append(paymentRoutesMap["GET"], getPayment.Name)

	paymentGroup.POST("", payment.AddPendingPaymentController)
	postPayment := paymentGroup.POST("", payment.AddPendingPaymentController)
	paymentRoutesMap["POST"] = append(paymentRoutesMap["POST"], postPayment.Name)	

	return paymentRoutesMap
}

