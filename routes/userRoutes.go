package routes

import(
	user "altastore/controllers/user"
	"altastore/middlewares"
	//"fmt"
	//"reflect"
)

func registerUserRoutes() map[string][]interface{} {
	userRoutesMap := map[string][]interface{}{}

	getUsers := e.GET("/users", user.GetUsersController, middlewares.AuthenticateUser)
	userRoutesMap["GET"] = append(userRoutesMap["GET"], getUsers.Name)

	postUser := e.POST("/users", user.AddUserController)
	userRoutesMap["POST"] = append(userRoutesMap["POST"], postUser.Name)

	r := e.Group("/users/:id")

	r.Use(middlewares.AuthenticateUser)

	r.Use(middlewares.CheckId)

	getUserById := r.GET("", user.GetUserByIdController)
	userRoutesMap["GET"] = append(userRoutesMap["GET"], getUserById.Name)

	editUserById := r.PUT("", user.EditUserController)
	userRoutesMap["PUT"] = append(userRoutesMap["GET"], editUserById.Name)

	deleteUserById := r.DELETE("", user.DeleteUserController)
	userRoutesMap["DELETE"] = append(userRoutesMap["GET"], deleteUserById.Name)
	
	return userRoutesMap
}

