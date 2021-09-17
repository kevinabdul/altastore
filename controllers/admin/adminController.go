package controllers

import (
	"net/http"
	"strconv"
	"fmt"

	libdb "altastore/lib/database"
	models "altastore/models"

	"github.com/labstack/echo/v4"
)

func GetAdminByUserIdController(c echo.Context) error {
	targetId, _ := strconv.Atoi(c.Param("id"))
	
	targetUser, err := libdb.GetAdminByUserId(targetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, struct {
		Status string
		User models.UserAPI
	}{Status: "success", User: targetUser})

}

func AddAdminController(c echo.Context) error {
	newUser := models.User{}
	c.Bind(&newUser)

	res, err := libdb.AddAdmin(&newUser)
	fmt.Println(err)

	if err != nil {
		return c.JSON(http.StatusOK, struct {
			Status string
			Message string
		}{Status: "success", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		User models.UserAPI
	}{Status: "success", Message: "Admin has been created!", User: res})

}
