package controllers

import (
	"net/http"
	"strconv"

	libdb "altastore/lib/database"
	models "altastore/models"

	"github.com/labstack/echo/v4"
)

func GetCheckoutByUserIdController(c echo.Context) error {
	userId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	checkout, err := libdb.GetCheckoutByUserId(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status 	string
			Message string
		}{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, struct {
		Status 	string
		Message string
		Data 	models.CheckoutAPI
	}{Status: "success", Message: "Cart is retrieved succesfully", Data: checkout})
}

func AddCheckoutByUserIdController(c echo.Context) error {
	userId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	payment := models.PaymentMethodAPI{}
	c.Bind(&payment)
	transactionAPI, err := libdb.AddCheckoutByUserId(&payment, userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status 	string
			Message string
		}{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, struct {
		Status 		string
		Message 	string
		Detail		models.TransactionAPI
	}{Status: "success",Message: "Checkout is succesfull", Detail: transactionAPI})
}