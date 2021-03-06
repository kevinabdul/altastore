package middlewares

import (
	// "net/http"
	// "net/http/httptest"
	// "encoding/json"
	"testing"

	"altastore/models"
	"altastore/config"
	"altastore/util/password"
	"altastore/util/jwt"

	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo/v4"

)

func InitProtectTest() *echo.Echo {
	config.InitDBTest("users")
	e := echo.New()
	return e
}

func AddUser(name, email, UserPassword string) uint{
	pass, _ := password.Hash(UserPassword)
	newUser := models.User{Name: name, Email: email, Password: pass}
	config.Db.Create(&newUser)
	return newUser.UserID
}

type AuthResponse struct {
	Status 	string
	Message string
}

type AuthCase struct {
	name			string
	method			string
	Path 		 	string
	expectedCode	int
	requestBody		string
	message			string
}

func Test_checkToken(t *testing.T) {
	userId := 1
	token, _ := implementjwt.CreateToken(userId)

	valid,_,_ := checkToken(token)
	invalid,_,_ := checkToken("asrdf.assa.qwert")
	
	assert.Equal(t, valid, true)
	assert.Equal(t, invalid, false)
}

// func Test_AuthenticateUser(t *testing.T) {
// 	e := InitProtectTest()

// 	testcases := []AuthCase{{
// 		name: "Invalid jwt",
// 		method: "GET",
// 		Path: "/users/1",
// 		expectedCode: http.StatusBadRequest,
// 		requestBody: "",
// 		message: "JWT is invalid!"}}

// 	for _, testcase := range testcases {
// 		req := httptest.NewRequest("GET", "/", nil)
// 		rec := httptest.NewRecorder()
// 		c := e.NewContext(req, rec)
// 		c.SetPath("/users/1")

// 		newUser := AddUser("Kevin", "kevin@gmail.com", "12345")
// 		validToken, _ := implementjwt.CreateToken(int(newUser.UserID))
// 		invalidToken := validToken[0: len(validToken) - 2]
// 		c.Request().Header.Set("Authorization", invalidToken)

// 		auth := AuthenticateUser(echo.HandlerFunc)
// 		if assert.NoError(t, auth(c)) {

// 			var authresponse AuthResponse

// 			json.Unmarshal([]byte(rec.Body.String()), &authresponse)

// 			assert.Equal(t, testcase.message, authresponse.Message)
// 		}
// 	}
// }