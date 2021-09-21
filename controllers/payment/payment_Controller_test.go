package controllers

import(
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
	"encoding/json"
	"fmt"
	"strconv"
	"bytes"

	"altastore/config"
	"altastore/models"
	"altastore/lib/database"

	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo/v4"
)

type PaymentCase struct {
	name 			string
	method  		string
	Path 			string
	expectedCode 	int
	requestBody   	string
	message			string
	size 			int
}

type GetPendingPaymentResponse struct {
	Status 			string
	Message 		string
	Data			[]models.PendingPaymentAPI
}

type AddPendingPaymentResponse struct {
	Status 			string
	Message 		string
	Detail			models.ReceiptAPI
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

func addCheckOut(payment *models.PaymentMethodAPI, userId int) (models.TransactionAPI, error) {
	transactionAPI, err := libdb.AddCheckoutByUserId(payment, userId) 
	if err != nil {
		return models.TransactionAPI{}, err
	} 
	return transactionAPI, nil
}

func AddPaymentMethods() {
	payment_methods := []models.PaymentMethod{{PaymentMethodName: "alfamart"}, {PaymentMethodName: "gopay"},{PaymentMethodName: "bank transfer"}}
	config.Db.Create(&payment_methods)
}

func InitConfigTest() *echo.Echo {
	config.InitDBTest("users", "categories", "products", "carts", "transactions", "transaction_details", "payment_methods")
	AddCategories()
	AddProducts()
	AddPaymentMethods()
	e := echo.New()
	return e
}

func Test_GetPendingPaymentsController(t *testing.T) {
	e := InitConfigTest()

	testcases := []PaymentCase{
		{
			name: "Get pending payments empty",
			method: "GET",
			Path: "/payments",
			expectedCode: http.StatusBadRequest,
			message: "No pending payments found" ,
			size: 0},
		{
			name:"Get pending payments exist",
			method: "GET",
			Path: "/payments",
			expectedCode: http.StatusOK,
			message: "Pending payments are retrieved succesfully",
			size: 1}}

	testcase0 := testcases[0]
	
	userId , _:= AddUser("Fattah", "fattah@gmail.com")

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(testcase0.Path)

	c.Request().Header.Set("userId", strconv.Itoa(int(userId)))

	if assert.NoError(t, GetPendingPaymentsController(c)) {

		var checkoutResponse GetPendingPaymentResponse

		json.Unmarshal([]byte(rec.Body.String()), &checkoutResponse)

		assert.Equal(t, testcase0.message, checkoutResponse.Message)
		assert.Equal(t, testcase0.size, len(checkoutResponse.Data))
	}

	AddItems(userId)

	addCheckOut(&models.PaymentMethodAPI{PaymentMethodID: 2}, int(userId))

	testcase1 := testcases[1]
	
	req2 := httptest.NewRequest("GET", "/", nil)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)

	c2.SetPath(testcase1.Path)

	c2.Request().Header.Set("userId", strconv.Itoa(int(userId)))

	if assert.NoError(t, GetPendingPaymentsController(c2)) {

		var checkoutResponse GetPendingPaymentResponse

		json.Unmarshal([]byte(rec2.Body.String()), &checkoutResponse)

		assert.Equal(t, testcase1.message, checkoutResponse.Message)
		assert.Equal(t, testcase1.size, len(checkoutResponse.Data))
	}

}

func Test_AddPendingPaymentController(t *testing.T) {
	e := InitConfigTest()

	userId , _:= AddUser("ali", "ali@gmail.com")

	AddItems(userId)

	transactionAPI, _ := addCheckOut(&models.PaymentMethodAPI{PaymentMethodID: 2}, int(userId))

	invalidInvoice := models.UserPaymentAPI{InvoiceID: "1234kss", Amount: transactionAPI.Total, PaymentMethodID: transactionAPI.PaymentMethodID}

	var invalidInvoiceData bytes.Buffer
	json.NewEncoder(&invalidInvoiceData).Encode(invalidInvoice)
	
	invalidPaymentMethod := models.UserPaymentAPI{InvoiceID: transactionAPI.InvoiceID, Amount: transactionAPI.Total, PaymentMethodID: uint(99)}

	var invalidPaymentMethodData bytes.Buffer
	json.NewEncoder(&invalidPaymentMethodData).Encode(invalidPaymentMethod)

	invalidAmount := models.UserPaymentAPI{InvoiceID: transactionAPI.InvoiceID, Amount: transactionAPI.Total + 12, PaymentMethodID: transactionAPI.PaymentMethodID}

	var invalidAmountData bytes.Buffer
	json.NewEncoder(&invalidAmountData).Encode(invalidAmount)

	validInvoice := models.UserPaymentAPI{InvoiceID: transactionAPI.InvoiceID, Amount: transactionAPI.Total, PaymentMethodID: transactionAPI.PaymentMethodID}

	var validInvoiceData bytes.Buffer
	json.NewEncoder(&validInvoiceData).Encode(validInvoice)

	testcases := []PaymentCase{
		{
			name: "Add payment with invalid invoice_id",
			method: "POST",
			Path: "/payments",
			expectedCode: http.StatusBadRequest,
			requestBody: invalidInvoiceData.String(),
			message: "No invoice_id found",
			size: 0},
		{
			name:"Add payment with invalid payment method",
			method: "POST",
			Path: "/payments",
			expectedCode: http.StatusBadRequest,
			requestBody: invalidPaymentMethodData.String(),
			message: fmt.Sprintf("Specified payment method doesnt match. Please pay using payment_method_id: %v", transactionAPI.PaymentMethodID),
			size: 0}, 
		{
			name:"Add payment with incorrect amount",
			method: "POST",
			Path: "/payments",
			expectedCode: http.StatusBadRequest,
			requestBody: invalidAmountData.String(),
			message: fmt.Sprintf("Specified amount doesnt match. Amount to be paid: %v", transactionAPI.Total),
			size: 0}, 
		{
			name:"Add valid payment",
			method: "POST",
			Path: "/payments",
			expectedCode: http.StatusOK,
			requestBody: validInvoiceData.String(),
			message:  "Payment is succesfull",
			size: 0}}

	for _, testcase := range testcases {
		req := httptest.NewRequest("POST", "/", strings.NewReader(testcase.requestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testcase.Path)

		c.Request().Header.Set("userId", strconv.Itoa(int(userId)))

		if assert.NoError(t, AddPendingPaymentController(c)) {

			var checkoutResponse GetPendingPaymentResponse

			json.Unmarshal([]byte(rec.Body.String()), &checkoutResponse)

			assert.Equal(t, testcase.message, checkoutResponse.Message)
		}	
	}

}