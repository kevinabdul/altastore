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

func InitProductTest() *echo.Echo{
	config.InitDBTest()
	AddCategories()
	AddProducts()
	e := echo.New()
	return e
}

func AddCategories() {
	categories := []models.Category{{CategoryName: "book"}, {CategoryName: "electronic device"}, {CategoryName: "sport equipment"}}
	config.Db.Create(&categories)
}

func AddProducts() {
	products := []models.Product{{ProductName: "Air Jordan M23", Price: 2300500, CategoryID: 3}, {ProductName: "Manusia Harimau", Price: 90500, CategoryID: 1}}
	config.Db.Create(&products)
}

func Test_GetProductsController(t *testing.T) {
	e := InitProductTest()

	cases := []GetProductCase{
		{
			name : "Get products",
			method: "GET",
			Path: "/products",
			expectedCode: http.StatusOK,
			message: "Products retrieval are succesfull",
			size : 2},
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