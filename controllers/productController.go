package controllers

import (
	"net/http"

	libdb "altastore/lib/database"
	models "altastore/models"

	"github.com/labstack/echo/v4"
)

func GetProductsController(c echo.Context) error {
	categoryName := c.QueryParam("category")

	productsTarget, err := libdb.GetProducts(categoryName)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status 	string
			Message string
		}{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, struct {
		Status 		string
		Message 	string	
		Products 	[]models.ProductAPI
	}{Status: "success", Message: "Products retrieval are succesfull", Products: productsTarget})
}
