package routes

import(
	handler "altastore/controllers"
	"altastore/middlewares"
	//"fmt"
	//"reflect"
)

func registerUserRoutes() map[string][]interface{} {
	userRoutesMap := map[string][]interface{}{}

	getUsers := e.GET("/users", handler.GetUsersController, middlewares.AuthenticateUser)
	userRoutesMap["GET"] = append(userRoutesMap["GET"], getUsers.Name)

	postUser := e.POST("/users", handler.AddUserController)
	userRoutesMap["POST"] = append(userRoutesMap["POST"], postUser.Name)

	r := e.Group("/users/:id")

	r.Use(middlewares.AuthenticateUser)

	r.Use(middlewares.CheckId)

	getUserById := r.GET("", handler.GetUserByIdController)
	userRoutesMap["GET"] = append(userRoutesMap["GET"], getUserById.Name)

	editUserById := r.PUT("", handler.EditUserController)
	userRoutesMap["PUT"] = append(userRoutesMap["GET"], editUserById.Name)

	deleteUserById := r.DELETE("", handler.DeleteUserController)
	userRoutesMap["DELETE"] = append(userRoutesMap["GET"], deleteUserById.Name)
	
	return userRoutesMap
}

