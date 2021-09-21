package controllers

import (
	"altastore/config"
	"altastore/models"
	"altastore/util/password"

	
	"testing"
	"net/http"
	"strings"
	"net/http/httptest"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo/v4"

)

type GetUserCase struct {
	name 			string
	method  		string
	Path 			string
	expectedCode	int
	message 		string
	size  			int
}

type UsersResponse struct {
	Status 	string
	Message string
	Users  []models.UserAPI			
}

type UserResponse struct {
	Status string
	Message string
	User  models.UserAPI
}

// add user, edit user, delete user, and login use this
// Whatever cases that need to send a request body can use this struct
type UserCaseWithBody struct {
	name 			string
	method  		string
	Path 			string
	expectedCode	int
	requestBody  	string
	message 		string
}

func initConfigTest() *echo.Echo{
	config.InitDBTest("users")
	e := echo.New()
	return e
}

func AddUser(name, email, userPassword string) *models.User {
	pass, _ := password.Hash(userPassword)
	newUser := models.User{Name: name, Email: email, Password: pass}
	config.Db.Create(&newUser)
	return &newUser
}

func Test_LoginUserController(t *testing.T) {
	e := initConfigTest()

	AddUser("abdul", "abdul@gmail.com", "1234")

	userReqOK := models.User{
		Name: "abdul",
		Email: "abdul@gmail.com",
		Password: "1234"}

	marshalledUserOk, _ := json.Marshal(userReqOK)

	userReqInvalidEmail := models.User{
		Name: "fattah",
		Email: "",
		Password: "1234"}

	marshalledUserInvalidEmail, _ := json.Marshal(userReqInvalidEmail)

	userReqInvalidPassword := models.User{
		Name: "abdul",
		Email: "abdul@gmail.com",
		Password: "123"}

	marshalledUserInvalidPassword, _ := json.Marshal(userReqInvalidPassword)

	cases := []UserCaseWithBody {
		 {
		 	name : "Valid login",
		 	method: "POST",
			Path : "/login",
			expectedCode: http.StatusOK,
			requestBody: string(marshalledUserOk),
			message:"You are logged in!"},
		{
		 	name : "Invalid login with invalid email",
		 	method: "POST",
			Path : "/login",
			expectedCode: http.StatusBadRequest,
			requestBody: string(marshalledUserInvalidEmail),
			message: "No user with corresponding email"},
		{
		 	name : "Invalid login with invalid password",
		 	method: "POST",
			Path : "/login",
			expectedCode: http.StatusBadRequest,
			requestBody: string(marshalledUserInvalidPassword),
			message: "Given password is incorrect"}}

	

	for _, testcase := range cases {
		req := httptest.NewRequest("POST", "/", strings.NewReader(testcase.requestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		
		c.SetPath(testcase.Path)

		if assert.NoError(t, LoginUserController(c)) {
			assert.Equal(t, testcase.expectedCode, rec.Code)


			var userResponse UserResponse
			
			if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testcase.message, userResponse.Message)
		}
	}
}

