package routes

import(
	admin "altastore/controllers/admin"
	"altastore/middlewares"
)

func registerAdminRoutes() {
	e.POST("/admins", admin.AddAdminController)

	r := e.Group("/admins/:id")

	r.Use(middlewares.AuthenticateUser)

	r.Use(middlewares.CheckId)

	r.GET("", admin.GetAdminByUserIdController)

	// r.PUT("", admin.EditAdminController)

	// r.DELETE("", admin.DeleteAdminController)
}

