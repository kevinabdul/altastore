package routes

import(
	handler "altastore/controllers"
)

func registerLoginRoutes() {
	e.POST("/login", handler.LoginUserController)

	// e.POST("/login/admins", user.LoginUserController)
}

