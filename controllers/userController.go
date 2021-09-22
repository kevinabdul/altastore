package controllers

import (
	"net/http"
	"strconv"
	//"fmt"

	libdb "altastore/lib/database"
	models "altastore/models"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	// Added for testing purpose
	tableName := c.QueryParam("table")
	users, err := libdb.GetUsers(tableName)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, struct {
		Status 	string
		Message string
		Users 	[]models.UserAPI
	}{Status: "success", Message: "Users are retrieved succesfully!", Users: users})
}

func GetUserByIdController(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))
	
	targetUser, err := libdb.GetUserById(targetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: err.Error()})
	}
	
	return c.JSON(http.StatusOK, struct {
		Status 	string
		Message string
		User 	models.UserAPI
	}{Status: "success", Message: "User retrieval is succesfull!", User: targetUser})
}

func AddUserController(c echo.Context) error {
	newUser := models.User{}
	c.Bind(&newUser)

	res, err := libdb.AddUser(&newUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		User models.UserAPI
	}{Status: "success", Message: "User has been created!", User: res})

}

func EditUserController(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))

	newData:= models.User{}
	c.Bind(&newData)

	edittedUser, err := libdb.EditUser(newData, targetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		User models.UserAPI
	}{Status: "success", Message: "User has been updated!", User: edittedUser})
}

func DeleteUserController(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))
	
	err := libdb.DeleteUser(targetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: err.Error()})
	}
	
	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
	}{Status: "success", Message: "User has been deleted!"})
}

// func LoginUserController(c echo.Context) error {
// 	loggingUser := &models.User{}
// 	c.Bind(loggingUser)

// 	token, err := libdb.LoginUser(loggingUser)

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, struct {
// 			Status string
// 			Message string
// 		}{Status: "failed", Message: err.Error()})
// 	}
	
// 	return c.JSON(http.StatusOK, struct {
// 		Status string
// 		Message string
// 		Token string
// 	}{Status: "success", Message: "You are logged in!", Token: token})
// }