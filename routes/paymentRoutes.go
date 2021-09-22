package routes

import(
	handler "altastore/controllers"
	"altastore/middlewares"
)

func registerPaymentRoutes() map[string][]interface{} {
	paymentGroup := e.Group("/payments")

	paymentGroup.Use(middlewares.AuthenticateUser)

	paymentRoutesMap := map[string][]interface{}{}

	getPayment := paymentGroup.GET("", handler.GetPendingPaymentsController)
	paymentRoutesMap["GET"] = append(paymentRoutesMap["GET"], getPayment.Name)

	paymentGroup.POST("", handler.AddPendingPaymentController)
	postPayment := paymentGroup.POST("", handler.AddPendingPaymentController)
	paymentRoutesMap["POST"] = append(paymentRoutesMap["POST"], postPayment.Name)	

	return paymentRoutesMap
}

