package routes

import(
	admin "altastore/controllers/admin"
	"altastore/middlewares"
)

func registerAdminRoutes() map[string][]interface{} {
	adminRoutesMap := map[string][]interface{}{}

	postAdmin := e.POST("/admins", admin.AddAdminController)
	adminRoutesMap["POST"] = append(adminRoutesMap["POST"], postAdmin.Name)

	r := e.Group("/admins/:id")

	r.Use(middlewares.AuthenticateUser)

	r.Use(middlewares.CheckId)

	getAdmin := r.GET("", admin.GetAdminByUserIdController, middlewares.AuthenticateUser)
	adminRoutesMap["GET"] = append(adminRoutesMap["GET"], getAdmin.Name)

	return adminRoutesMap
}

