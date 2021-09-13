package controllers

import (
	"net/http"
	"strconv"

	libdb "altastore/lib/database"
	models "altastore/models"

	"github.com/labstack/echo/v4"
)

func GetPendingPaymentsController(c echo.Context) error {
	userId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	pendingPayments, err := libdb.GetPendingPaymentsByUserId(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status 	string
			Message string
		}{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, struct {
		Status 	string
		Data 	[]models.PendingPaymentAPI
	}{Status: "success", Data: pendingPayments})
}

func AddPendingPaymentController(c echo.Context) error {
	userId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	payment := models.UserPaymentAPI{}
	c.Bind(&payment)
	receiptAPI, err := libdb.AddPendingPaymentByUserId(payment, userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status 	string
			Message string
		}{Status: "failed", Message: err.Error()})
	}

	// if rowsAffected == 0 {
	// 	return c.JSON(http.StatusBadRequest, struct {
	// 		Status 	string
	// 		Message string
	// 	}{Status: "failed", Message: "fail to make a payment!"})
	// }

	return c.JSON(http.StatusOK, struct {
		Status 		string
		Detail		models.ReceiptAPI
	}{Status: "success", Detail: receiptAPI})
}