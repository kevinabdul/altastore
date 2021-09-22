package routes

import(
	handler "altastore/controllers"
	"altastore/middlewares"
)

func registerAdminRoutes() map[string][]interface{} {
	adminRoutesMap := map[string][]interface{}{}

	postAdmin := e.POST("/admins", handler.AddAdminController)
	adminRoutesMap["POST"] = append(adminRoutesMap["POST"], postAdmin.Name)

	r := e.Group("/admins/:id")

	r.Use(middlewares.AuthenticateUser)

	r.Use(middlewares.CheckId)

	getAdmin := r.GET("", handler.GetAdminByUserIdController, middlewares.AuthenticateUser)
	adminRoutesMap["GET"] = append(adminRoutesMap["GET"], getAdmin.Name)

	return adminRoutesMap
}

