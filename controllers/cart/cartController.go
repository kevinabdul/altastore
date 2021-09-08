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

func UpdateCartByUserIdController(c echo.Context) error {
	userId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	
	var userCart []models.Cart
	c.Bind(&userCart)
	
	rowsAffected, err := libdb.UpdateCartByUserId(userCart, userId)	

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status 	string
			Message string
		}{Status: "failed", Message: err.Error()})
	}

	if rowsAffected == int64(0) {
		return c.JSON(http.StatusBadRequest, struct {
			Status 	string
			Message string
		}{Status: "failed", Message: "No change done to cart"})
	}

	return c.JSON(http.StatusOK, struct {
		Status 	string
		Message	string
	}{Status: "success", Message: "Cart is updated!"})
}