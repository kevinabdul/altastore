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

type CheckoutCase struct {
	name 			string
	method  		string
	Path 			string
	expectedCode 	int
	requestBody   	string
	message			string
	size 			int
}

type GetCheckoutResponse struct {
	Status 			string
	Message 		string
	Data			models.CheckoutAPI
}

func AddUser(name, email string) (uint,error) {
	user := models.User{Name: name, Email: email, Password: "1234"}
	res := config.Db.FirstOrCreate(&user)
	
	if res.Error != nil {
		return uint(0), res.Error
	}
	return user.UserID, nil
}

func AddCategories() {
	categories := []models.Category{{CategoryName: "book"}, {CategoryName: "electronic device"}, {CategoryName: "sport equipment"}}
	config.Db.Create(&categories)
}

func AddProducts() {
	products := []models.Product{{ProductName: "Air Jordan M23", Price: 2300500, CategoryID: 3}, {ProductName: "Manusia Harimau", Price: 90500, CategoryID: 1}}
	config.Db.Create(&products)
}

func AddItems(userId uint) (int64,error){
	cart := []models.Cart{{UserID: userId, ProductID: 1, Quantity: 3}, {UserID: userId, ProductID: 2, Quantity: 1}}
	res := config.Db.Create(&cart); 
	
	if res != nil || res.RowsAffected == 0 {
		return res.RowsAffected, res.Error
	}
	return res.RowsAffected, nil
}

func AddPaymentMethods() {
	payment_methods := []models.PaymentMethod{{PaymentMethodName: "alfamart"}, {PaymentMethodName: "gopay"},{PaymentMethodName: "bank transfer"}}
	config.Db.Create(&payment_methods)
}

func InitConfigTest() *echo.Echo {
	config.InitDBTest("users", "categories", "products", "carts", "payment_methods")
	AddCategories()
	AddProducts()
	AddPaymentMethods()
	e := echo.New()
	return e
}

func Test_GetCheckoutByUserIdController(t *testing.T) {
	e := InitConfigTest()

	testcases := []CheckoutCase{
		{
			name: "Get user checkout empty",
			method: "GET",
			Path: "/checkout",
			expectedCode: http.StatusBadRequest,
			message: "No product found in the cart" ,
			size: 0},
		{
			name:"Get user checkout item exist",
			method: "GET",
			Path: "/checkout",
			expectedCode: http.StatusOK,
			message: "Cart is retrieved succesfully",
			size: 2}}

	testcase0 := testcases[0]
	
	userId , _:= AddUser("Fattah", "fattah@gmail.com")

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(testcase0.Path)

	c.Request().Header.Set("userId", strconv.Itoa(int(userId)))

	if assert.NoError(t, GetCheckoutByUserIdController(c)) {

		var checkoutResponse GetCheckoutResponse

		json.Unmarshal([]byte(rec.Body.String()), &checkoutResponse)

		assert.Equal(t, testcase0.message, checkoutResponse.Message)
		assert.Equal(t, testcase0.size, len(checkoutResponse.Data.Products))
	}

	AddItems(userId)

	testcase1 := testcases[1]
	
	req2 := httptest.NewRequest("GET", "/", nil)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)

	c2.SetPath(testcase1.Path)

	c2.Request().Header.Set("userId", strconv.Itoa(int(userId)))

	if assert.NoError(t, GetCheckoutByUserIdController(c2)) {

		var checkoutResponse GetCheckoutResponse

		json.Unmarshal([]byte(rec2.Body.String()), &checkoutResponse)

		assert.Equal(t, testcase1.message, checkoutResponse.Message)
		assert.Equal(t, testcase1.size, len(checkoutResponse.Data.Products))
	}

}

func Test_AddCheckoutByUserIdController(t *testing.T) {
	e := InitConfigTest()

	invalidCheckout := models.PaymentMethodAPI{PaymentMethodID: 1}

	var validData bytes.Buffer
	json.NewEncoder(&validData).Encode(invalidCheckout)
	
	invalidCheckout2 := models.PaymentMethodAPI{PaymentMethodID: 112}

	var invalidData2 bytes.Buffer
	json.NewEncoder(&invalidData2).Encode(invalidCheckout2)

	testcases := []CheckoutCase{
		{
			name: "Add checkout with empty cart",
			method: "POST",
			Path: "/checkout",
			expectedCode: http.StatusBadRequest,
			requestBody: validData.String(),
			message: "No products found in the cart. Add products first before checking out",
			size: 0},
		{
			name:"Add checkout with invalid payment method",
			method: "POST",
			Path: "/checkout",
			expectedCode: http.StatusBadRequest,
			requestBody: invalidData2.String(),
			message: "No payment_method_id '112' found in the payment method table",
			size: 0}, 
		{
			name:"Add checkout wit valid payment method",
			method: "POST",
			Path: "/checkout",
			expectedCode: http.StatusBadRequest,
			requestBody: validData.String(),
			message: "Checkout is succesfull",
			size: 0}}

	userId , _:= AddUser("ali", "ali@gmail.com")


	req := httptest.NewRequest("POST", "/", strings.NewReader(testcases[0].requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(testcases[0].Path)

	c.Request().Header.Set("userId", strconv.Itoa(int(userId)))

	if assert.NoError(t, AddCheckoutByUserIdController(c)) {

		var checkoutResponse GetCheckoutResponse

		json.Unmarshal([]byte(rec.Body.String()), &checkoutResponse)

		assert.Equal(t, testcases[0].message, checkoutResponse.Message)
	}

	AddItems(userId)

	req = httptest.NewRequest("POST", "/", strings.NewReader(testcases[1].requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	c.SetPath(testcases[1].Path)

	c.Request().Header.Set("userId", strconv.Itoa(int(userId)))

	if assert.NoError(t, AddCheckoutByUserIdController(c)) {

		var checkoutResponse GetCheckoutResponse

		json.Unmarshal([]byte(rec.Body.String()), &checkoutResponse)

		assert.Equal(t, testcases[1].message, checkoutResponse.Message)
	}


	req = httptest.NewRequest("POST", "/", strings.NewReader(testcases[2].requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	c.SetPath(testcases[2].Path)

	c.Request().Header.Set("userId", strconv.Itoa(int(userId)))

	if assert.NoError(t, AddCheckoutByUserIdController(c)) {

		var checkoutResponse GetCheckoutResponse

		json.Unmarshal([]byte(rec.Body.String()), &checkoutResponse)

		assert.Equal(t, testcases[2].message, checkoutResponse.Message)
	}

}