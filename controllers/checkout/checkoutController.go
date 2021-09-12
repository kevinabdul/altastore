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
		Data 	models.CheckoutAPI
	}{Status: "success", Data: checkout})
}

func AddCheckoutByUserIdController(c echo.Context) error {
	userId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	payment := models.PaymentMethodAPI{}
	c.Bind(&payment)
	transactionAPI, rowsAffected, err := libdb.AddCheckoutByUserId(&payment, userId)

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
		}{Status: "failed", Message: "fail to create invoice!"})
	}

	return c.JSON(http.StatusOK, struct {
		Status 		string
		Detail		models.TransactionAPI
	}{Status: "success", Detail: transactionAPI})
}