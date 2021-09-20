package controllers

import (
	"altastore/config"
	"altastore/models"

	"testing"
	"net/http"
	"net/url"
	"strings"
	"net/http/httptest"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo/v4"

)

type GetProductCase struct {
	name 			string
	method  		string
	Path 			string
	expectedCode	int
	message 		string
	size  			int
}

type ProductsResponse struct {
	Status 		string
	Message 	string
	Products  	[]models.ProductAPI
}

func InsertDummyCategories() error{
	data := []models.Category{{CategoryName: "book"}, {CategoryName: "electronic device"}, {CategoryName: "sport equipment"}}

	if err := config.Db.Table("categories").Create(&data).Error; err != nil {
		return err
	}

	return nil
}


func InsertDummyProducts() error{
	data := []models.Product{{ProductName: "Air Jordan M23", CategoryID: 3, Price: 2400000}, 
	{ProductName: "Air Jordan M24", CategoryID: 3, Price: 2600000}, 
	{ProductName: "Iphone 13", CategoryID: 2, Price: 21500000}}
	
	if err := config.Db.Table("products").Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func initConfigTest() *echo.Echo{
	config.InitDBTest()
	InsertDummyCategories()
	InsertDummyProducts()
	e := echo.New()
	return e
}

var e = initConfigTest()

func Test_GetProductsController(t *testing.T) {
	cases := []GetProductCase{
		{
			name : "Get products",
			method: "GET",
			Path: "/products",
			expectedCode: http.StatusOK,
			message: "Products retrieval are succesfull",
			size : 3},
		{
			name : "Get products with invalid category query value",
			method: "GET",
			Path: "/products?category=kitchen utility",
			expectedCode: http.StatusBadRequest,
			message: "No product found for the given cateogory",
			size : 0}}

	

	for _, testcase := range cases {
		var queryValues string
		temp := strings.Split(testcase.Path, "?")
		if len(temp) == 1 {
			queryValues = ""
		} else {
			queryValues = strings.Split(temp[1], "=")[1]
		}

		q := make(url.Values)
		q.Set("category", queryValues)

		req := httptest.NewRequest("GET", "/?"+q.Encode(), nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, GetProductsController(c)) {
			assert.Equal(t, testcase.expectedCode, rec.Code)

			var productResponse ProductsResponse

			if err := json.Unmarshal([]byte(rec.Body.String()), &productResponse); err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testcase.size, len(productResponse.Products))

			assert.Equal(t, testcase.message, productResponse.Message)
		}
	}
}