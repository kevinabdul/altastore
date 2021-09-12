package routes

import(
	login "altastore/controllers/login"
)

func registerLoginRoutes() {
	e.POST("/login/users", login.LoginUserController)

	// e.POST("/login/admins", user.LoginUserController)
}

