package controllers

import (
	"net/http"
	"strconv"

	libdb "altastore/lib/database"
	models "altastore/models"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	users, err := libdb.GetUsers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func GetUserByIdController(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))
	
	targetUser, rowsAffected, err := libdb.GetUserById(targetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Wrong User Id"})
	}
	return c.JSON(http.StatusOK, targetUser)
}

func AddUserController(c echo.Context) error {
	newUser := models.User{}
	c.Bind(&newUser)

	if newUser.Email == "" || newUser.Password == "" {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Invalid Email or Password. Make sure its not empty and are of string type"})
	}

	res, err := libdb.AddUser(&newUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
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

	edittedUser, rowsAffected, err := libdb.EditUser(newData, targetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Wrong User Id"})
	}

	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		User models.UserAPI
	}{Status: "success", Message: "User has been updated!", User: edittedUser})
}

func DeleteUserController(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))
	
	rowsAffected, err := libdb.DeleteUser(targetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Wrong User Id"})
	}
	

	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
	}{Status: "success", Message: "User has been deleted!"})

}
