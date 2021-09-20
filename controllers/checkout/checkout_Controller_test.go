package controllers

// import(
// 	"testing"
// 	"net/http"
// 	"net/http/httptest"
// 	//"strings"
// 	"encoding/json"
// 	"fmt"

// 	"altastore/config"
// 	"altastore/models"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/labstack/echo/v4"
// )

// type CheckoutCase struct {
// 	name 			string
// 	method  		string
// 	Path 			string
// 	expectedCode 	int
// 	message			string
// 	size 			int
// }

// type GetCheckoutResponse struct {
// 	Status 			string
// 	Message 		string
// 	Data 			[]models.CheckoutAPI
// }

// type AddCheckoutResponse struct {
// 	Status 			string
// 	Message 		string
// 	Detail 			models.TransactionAPI
// }

// func InitConfigTest() *echo.Echo {
// 	config.InitDBTest()
// 	e := echo.New()
// 	return e
// }

// var e = InitConfigTest()

// func AddDummyUserAndItem() error{
// 	dummyUser := models.User{Name: "Fattah", Email: "fattah@gmail.com", Password: "1234"}
// 	config.Db.Create(&dummyUser)
// 	fmt.Println(dummyUser)
// 	dummyCart := []models.Cart{{UserID: dummyUser.UserID, ProductID: 1, Quantity: 3}, {UserID: dummyUser.UserID, ProductID: 3, Quantity: 1}}
// 	err := config.Db.Create(&dummyCart).Error; 
	
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func Test_GetCheckoutByUserIdController(t *testing.T) {

// 	testcases := []CheckoutCase{
// 		{
// 			name: "Get user checkout empty",
// 			method: "GET",
// 			Path: "/checkout",
// 			expectedCode: http.StatusBadRequest,
// 			message: "No product found in the cart" ,
// 			size: 0},
// 		{
// 			name:"Get user checkout item exist",
// 			method: "GET",
// 			Path: "/checkout",
// 			expectedCode: http.StatusOK,
// 			message: "Cart is retrieved succesfully",
// 			size: 1}}

// 	testcase0 := testcases[0]
	
// 	req := httptest.NewRequest("GET", "/", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	c.SetPath(testcase0.Path)

// 	if assert.NoError(t, GetCheckoutByUserIdController(c)) {

// 		var checkoutResponse GetCheckoutResponse

// 		json.Unmarshal([]byte(rec.Body.String()), &checkoutResponse)

// 		assert.Equal(t, testcase0.message, checkoutResponse.Message)
// 		assert.Equal(t, testcase0.size, len(checkoutResponse.Data))
// 	}

// 	AddDummyUserAndItem()

// 	testcase1 := testcases[1]
	
// 	req2 := httptest.NewRequest("GET", "/", nil)
// 	rec2 := httptest.NewRecorder()
// 	c2 := e.NewContext(req2, rec2)

// 	c2.SetPath(testcase1.Path)
// 	c2.Request().Header.Set("userId", "1")

// 	if assert.NoError(t, GetCheckoutByUserIdController(c2)) {

// 		var checkoutResponse GetCheckoutResponse

// 		json.Unmarshal([]byte(rec.Body.String()), &checkoutResponse)

// 		assert.Equal(t, testcase1.message, checkoutResponse.Message)
// 		assert.Equal(t, testcase1.size, len(checkoutResponse.Data))
// 	}

	

// }