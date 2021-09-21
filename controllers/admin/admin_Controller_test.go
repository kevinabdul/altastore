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

type GetAdminCase struct {
	name 			string
	method  		string
	Path 			string
	expectedCode	int
	message 		string
	size  			int
}

type AdminResponse struct {
	Status string
	Message string
	User  models.UserAPI
}

// add user, edit user, delete user, and login use this
// Whatever cases that need to send a request body can use this struct
type AdminCaseWithBody struct {
	name 			string
	method  		string
	Path 			string
	expectedCode	int
	requestBody  	string
	message 		string
}

func initConfigTest() *echo.Echo{
	config.InitDBTest("users", "admins")
	e := echo.New()
	return e
}

func AddUser(name, email, userPassword string) *models.User {
	pass, _ := password.Hash(userPassword)
	newUser := models.User{Name: name, Email: email, Password: pass}
	config.Db.Create(&newUser)
	return &newUser
}

func AddAdmin(userId uint) *models.Admin {
	newAdmin := models.Admin{UserID: userId}
	config.Db.Create(&newAdmin)
	return &newAdmin
}


func Test_AddAdminController(t *testing.T) {
	e := initConfigTest()

	userReqOK := models.User{
		Name: "abdul",
		Email: "abdul@gmail.com",
		Password: "1234"}

	marshalledUserOk, _ := json.Marshal(userReqOK)

	userReqInvalidEmail := models.User{
		Name: "abdul",
		Email: "",
		Password: "1234"}

	marshalledUserInvalidEmail, _ := json.Marshal(userReqInvalidEmail)

	cases := []AdminCaseWithBody {
		 {
		 	name : "Add user",
		 	method: "POST",
			Path : "/users",
			expectedCode: http.StatusOK,
			requestBody: string(marshalledUserOk),
			message:"Admin has been created!"},
		{
		 	name : "Add user with invalid email",
		 	method: "POST",
			Path : "/users",
			expectedCode: http.StatusBadRequest,
			requestBody: string(marshalledUserInvalidEmail),
			message:"Invalid Email or Password. Make sure its not empty and are of string type"},
		{
		 	name : "Add duplicate user",
		 	method: "POST",
			Path : "/users",
			expectedCode: http.StatusBadRequest,
			requestBody: string(marshalledUserOk),
			message: "Email is already taken"}}


	for _, testcase := range cases {
		req := httptest.NewRequest("POST", "/", strings.NewReader(testcase.requestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testcase.Path)

		if assert.NoError(t, AddAdminController(c)) {
			assert.Equal(t, testcase.expectedCode, rec.Code)

			var AdminResponse AdminResponse
			
			if err := json.Unmarshal([]byte(rec.Body.String()), &AdminResponse); err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testcase.message, AdminResponse.Message)
		}
	}
}

func Test_GetAdminByUserIdController(t *testing.T) {
	e := initConfigTest()

	newUser := AddUser("kevin", "kevin@gmail.com", "1234")
	AddAdmin(newUser.UserID)

	cases := []GetAdminCase{
		{
			name : "Get admin with valid id",
			method: "GET",
			Path: "/users/1",
			expectedCode: http.StatusOK,
			message: "Admin is retrieved succesfully",
			size : 1},
		{
			name : "Get user with invalid id",
			method: "GET",
			Path: "/users/1123456",
			expectedCode: http.StatusBadRequest,
			message: "Wrong User Id",
			size : 0}}

	for _, testcase := range cases {
		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testcase.Path)
		c.SetParamNames("id")
		paramValues := strings.Split(testcase.Path, "/")
		c.SetParamValues(paramValues[2])

		if assert.NoError(t, GetAdminByUserIdController(c)) {
			assert.Equal(t, testcase.expectedCode, rec.Code)

			var AdminResponse AdminResponse

			if err := json.Unmarshal([]byte(rec.Body.String()), &AdminResponse); err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testcase.message, AdminResponse.Message)
		}
	}
}