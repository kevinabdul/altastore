package controllers

import (
	"altastore/config"
	"altastore/models"
	"altastore/util/password"

	//"fmt"
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
	path 			string
	expectedCode	int
	message 		string
	size  			int
}

type GetUsersResponse struct {
	Status 	string
	Message string
	Users  []models.UserAPI			
}

type GetUserByIdResponse struct {
	Status 	string
	Message string
	User  models.UserAPI			
}

type AddUserResponse struct {
	Status string
	Message string
}

type AddUserCase struct {
	name 			string
	method  		string
	path 			string
	expectedCode	int
	requestBody  	string
	message 		string
}
func InsertUser(name, email, givenPassword string) error {
	pass, _ := password.Hash(givenPassword)

	newUser := models.User{
		Name: name,
		Email: email,
		Password: pass}

	if err :=config.Db.Model(models.User{}).Create(&newUser).Error; err != nil {
		return err
	}

	return nil
}

func initConfigTest() *echo.Echo{
	config.InitDBTest()
	InsertUser("kevin", "kevin@gmail.com", "1234")
	InsertUser("fattah", "fattah@gmail.com", "1234")
	e := echo.New()
	return e
}

var e = initConfigTest()

func Test_GetUsersController(t *testing.T) {
	cases := []GetUserCase{
		{
			name : "Get users",
			method: "GET",
			path: "/users",
			expectedCode: http.StatusOK,
			message: "Users are retrieved succesfully!",
			size : 2}}

	for _, testcase := range cases {
		req := httptest.NewRequest(testcase.method, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
	
		c.SetPath(testcase.path)

		if assert.NoError(t, GetUsersController(c)) {
			assert.Equal(t, testcase.expectedCode, rec.Code)

			var userResponse GetUsersResponse

			if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testcase.size, len(userResponse.Users))
		}
	}
}

func Test_AddUserController(t *testing.T) {
	userReqOK := models.User{
		Name: "abdul",
		Email: "abdul@gmail.com",
		Password: "1234"}

	marshalledUserOk, _ := json.Marshal(userReqOK)
	
	cases := []AddUserCase{
		 {
		 	name : "Add user",
		 	method: "POST",
			path : "/users",
			expectedCode: http.StatusOK,
			requestBody: string(marshalledUserOk),
			message:"User has been created!"}}

	//e := initConfigTest()

	for _, testcase := range cases {
		req := httptest.NewRequest(testcase.method, "/", strings.NewReader(testcase.requestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
	
		c.SetPath(testcase.path)

		if assert.NoError(t, AddUserController(c)) {
			assert.Equal(t, testcase.expectedCode, rec.Code)


			var userResponse AddUserResponse
			
			if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testcase.message, userResponse.Message)
		}
	}
}