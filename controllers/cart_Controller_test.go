package controllers

import(
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
	"encoding/json"
	//"fmt"
	"strconv"
	"bytes"

	"altastore/config"
	"altastore/models"

	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo/v4"
)

type CartCase struct {
	name 			string
	method  		string
	Path 			string
	expectedCode 	int
	requestBody   	string
	message			string
	size 			int
}

type GetCartResponse struct {
	Status 			string
	Message 		string
	Carts			[]models.CartAPI
}

// type AddcartResponse struct {
// 	Status 			string
// 	Message 		string
// 	Detail 			models.TransactionAPI
// }

func AddItems(userId uint) (int64,error){
	cart := []models.Cart{{UserID: userId, ProductID: 1, Quantity: 3}, {UserID: userId, ProductID: 2, Quantity: 1}}
	res := config.Db.Create(&cart); 
	
	if res != nil || res.RowsAffected == 0 {
		return res.RowsAffected, res.Error
	}
	return res.RowsAffected, nil
}

func InitCartTest() *echo.Echo {
	config.InitDBTest("users", "categories", "products", "carts")
	AddCategories()
	AddProducts()
	e := echo.New()
	return e
}

func Test_GetCartByUserIdController(t *testing.T) {

	e := InitCartTest()

	testcases := []CartCase{
		{
			name: "Get user cart empty",
			method: "GET",
			Path: "/carts",
			expectedCode: http.StatusBadRequest,
			message: "No product found in the cart" ,
			size: 0},
		{
			name:"Get user cart item exist",
			method: "GET",
			Path: "/carts",
			expectedCode: http.StatusOK,
			message: "Cart is retrieved succesfully",
			size: 2}}

	testcase0 := testcases[0]
	
	userId , _:= AddUser("Fattah", "fattah@gmail.com", "1234")

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(testcase0.Path)

	c.Request().Header.Set("userId", strconv.Itoa(int(userId)))

	if assert.NoError(t, GetCartByUserIdController(c)) {

		var cartResponse GetCartResponse

		json.Unmarshal([]byte(rec.Body.String()), &cartResponse)

		assert.Equal(t, testcase0.message, cartResponse.Message)
		assert.Equal(t, testcase0.size, len(cartResponse.Carts))
	}

	AddItems(userId)

	testcase1 := testcases[1]
	
	req2 := httptest.NewRequest("GET", "/", nil)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)

	c2.SetPath(testcase1.Path)

	c2.Request().Header.Set("userId", strconv.Itoa(int(userId)))

	if assert.NoError(t, GetCartByUserIdController(c2)) {

		var cartResponse GetCartResponse

		json.Unmarshal([]byte(rec2.Body.String()), &cartResponse)

		assert.Equal(t, testcase1.message, cartResponse.Message)
		assert.Equal(t, testcase1.size, len(cartResponse.Carts))
	}

}

func Test_UpdateCartByUserIdController(t *testing.T) {
	e := InitCartTest()

	validUpdate := []models.Cart{
		{
			ProductID: 1, Quantity: 1}, 
		{
			ProductID: 2, Quantity: 1}}

	var validData bytes.Buffer
	json.NewEncoder(&validData).Encode(validUpdate)

	invalidUpdate := []models.Cart{{ProductID: 9891211, Quantity: 1}}

	var invalidData bytes.Buffer
	json.NewEncoder(&invalidData).Encode(invalidUpdate)


	testcases := []CartCase{
		{
			name: "Update user cart",
			method: "PUT",
			Path: "/carts",
			expectedCode: http.StatusOK,
			requestBody: validData.String(),
			message: "Cart is updated!",
			size: 0},
		{
			name:"Update user cart with invalid item",
			method: "PUT",
			Path: "/carts",
			expectedCode: http.StatusBadRequest,
			requestBody: invalidData.String(),
			message: "No product id 9891211 found in the product table",
			size: 0}}

	userId , _:= AddUser("Ali", "ali@gmail.com", "1234")

	for _, testcase := range testcases {	
		req := httptest.NewRequest("PUT", "/", strings.NewReader(testcase.requestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testcase.Path)

		c.Request().Header.Set("userId", strconv.Itoa(int(userId)))

		if assert.NoError(t, UpdateCartByUserIdController(c)) {

			var cartResponse GetCartResponse

			json.Unmarshal([]byte(rec.Body.String()), &cartResponse)

			assert.Equal(t, testcase.message, cartResponse.Message)
			assert.Equal(t, testcase.size, len(cartResponse.Carts))
		}
	}
}

func Test_DeleteCartByUserIdController(t *testing.T) {
	e := InitCartTest()

	userId , _:= AddUser("Ankara", "ankara@gmail.com", "1234")

	AddItems(userId)

	validDelete := []int{2} 
	var validData bytes.Buffer
	json.NewEncoder(&validData).Encode(validDelete)

	invalidDelete := []int{9891211}
	var invalidData bytes.Buffer
	json.NewEncoder(&invalidData).Encode(invalidDelete)

	invalidDeleteEmpty := []int{} 
	var invalidDataEmpty bytes.Buffer
	json.NewEncoder(&invalidDataEmpty).Encode(invalidDeleteEmpty)


	testcases := []CartCase{
		{
			name: "Delete user cart",
			method: "DELETE",
			Path: "/carts",
			expectedCode: http.StatusOK,
			requestBody: validData.String(),
			message: "Cart is updated!",
			size: 0},
		{
			name:"Delete user cart with invalid item id",
			method: "DELETE",
			Path: "/carts",
			expectedCode: http.StatusBadRequest,
			requestBody: invalidData.String(),
			message: "No product with id 9891211 is found in user's cart.",
			size: 0}, 
		{
			name: "Delete user cart empty",
			method: "DELETE",
			Path: "/carts",
			expectedCode: http.StatusBadRequest,
			requestBody: invalidDataEmpty.String(),
			message: "No item found in delete list. Please specify before deleting",
			size: 0},}

	for _, testcase := range testcases {	
		req := httptest.NewRequest("DELETE", "/", strings.NewReader(testcase.requestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testcase.Path)

		c.Request().Header.Set("userId", strconv.Itoa(int(userId)))

		if assert.NoError(t, DeleteCartByUserIdController(c)) {

			var cartResponse GetCartResponse

			json.Unmarshal([]byte(rec.Body.String()), &cartResponse)

			assert.Equal(t, testcase.message, cartResponse.Message)
			assert.Equal(t, testcase.size, len(cartResponse.Carts))
		}
	}
}