package controllers

import (
	"net/http"
	"strconv"

	libdb "altastore/lib/database"
	models "altastore/models"

	"github.com/labstack/echo/v4"
)

func GetCartByUserIdController(c echo.Context) error {
	userId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	cartTarget, err := libdb.GetCartByUserId(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status 	string
			Message string
		}{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, struct {
		Status 	string
		Cart 	[]models.CartAPI
	}{Status: "success", Cart: cartTarget})
}


// Update is for addition and update of item(s). Subsequent request without previously added item(s) wont discard the already added item(s)
// Attempt to set quantity if an item to zero will be ignored
// You should use delete endpoint to delete item(s)
func UpdateCartByUserIdController(c echo.Context) error {
	userId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	
	var userCart []models.Cart
	c.Bind(&userCart)
	
	err := libdb.UpdateCartByUserId(userCart, userId)	

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status 	string
			Message string
		}{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, struct {
		Status 	string
		Message	string
	}{Status: "success", Message: "Cart is updated!"})
}

func DeleteCartByUserIdController(c echo.Context) error {
	userId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	
	var userCart []int
	c.Bind(&userCart)
	
	rowsAffected, err := libdb.DeleteCartByUserId(userCart, userId)	

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status 	string
			Message string
		}{Status: "failed", Message: err.Error()})
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status 	string
			Message string
		}{Status: "success", Message: "No change in user's cart"})
	}

	return c.JSON(http.StatusOK, struct {
		Status 	string
		Message	string
	}{Status: "success", Message: "Cart is updated!"})
}