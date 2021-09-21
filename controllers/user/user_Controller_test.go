package controllers

import (
	"altastore/config"
	"altastore/models"
	"altastore/util/password"

	
	"testing"
	"net/http"
	"net/url"
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
	config.InitDBTest()
	e := echo.New()
	return e
}

func AddUser(name, email, userPassword string) *models.User {
	pass, _ := password.Hash(userPassword)
	newUser := models.User{Name: name, Email: email, Password: pass}
	config.Db.Create(&newUser)
	return &newUser
}

func Test_AddUserController(t *testing.T) {
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

	userReqInvalidName := models.User{
		Name: "",
		Email: "empty@gmail.com",
		Password: "1234"}

	marshalledUserInvalidName, _ := json.Marshal(userReqInvalidName)


	cases := []UserCaseWithBody {
		 {
		 	name : "Add user",
		 	method: "POST",
			Path : "/users",
			expectedCode: http.StatusOK,
			requestBody: string(marshalledUserOk),
			message:"User has been created!"},
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
			message: "Error 1062: Duplicate entry 'abdul@gmail.com' for key 'users.email'"}, 
		{
		 	name : "Add user without name",
		 	method: "POST",
			Path : "/users",
			expectedCode: http.StatusBadRequest,
			requestBody: string(marshalledUserInvalidName),
			message:"Name cant be empty"},}


	for _, testcase := range cases {
		req := httptest.NewRequest("POST", "/", strings.NewReader(testcase.requestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testcase.Path)

		if assert.NoError(t, AddUserController(c)) {
			assert.Equal(t, testcase.expectedCode, rec.Code)

			var userResponse UserResponse
			
			if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testcase.message, userResponse.Message)
		}
	}
}

func Test_GetUsersController(t *testing.T) {
	e := initConfigTest()

	AddUser("Fattah","fattah@gmail.com", "1234")

	cases := []GetUserCase{
		{
			name : "Get users",
			method: "GET",
			Path: "/users",
			expectedCode: http.StatusOK,
			message: "Users are retrieved succesfully!",
			size : 1},
		{
			name : "Get users with invalid table query value",
			method: "GET",
			Path: "/users?table=userlist",
			expectedCode: http.StatusBadRequest,
			message: "Table doesnt exist",
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
		q.Set("table", queryValues)

		req := httptest.NewRequest("GET", "/?"+q.Encode(), nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, GetUsersController(c)) {
			assert.Equal(t, testcase.expectedCode, rec.Code)

			var userResponse UsersResponse

			if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testcase.size, len(userResponse.Users))

			assert.Equal(t, testcase.message, userResponse.Message)
		}
	}
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

func Test_GetUserByIdController(t *testing.T) {
	e := initConfigTest()

	AddUser("kevin", "kevin@gmail.com", "1234")

	cases := []GetUserCase{
		{
			name : "Get user with valid id",
			method: "GET",
			Path: "/users/1",
			expectedCode: http.StatusOK,
			message: "User retrieval is succesfull!",
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

		if assert.NoError(t, GetUserByIdController(c)) {
			assert.Equal(t, testcase.expectedCode, rec.Code)

			var userResponse UserResponse

			if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testcase.message, userResponse.Message)
		}
	}
}

func Test_EditUserController(t *testing.T) {
	e := initConfigTest()

	AddUser("kevin", "kevin@gmail.com", "1234")

	validEdit := models.User{
		Name: "Fattah Abdul",
		Email: "fattah.abdul@gmail.com",
		}

	marshalledValidEdit, _ := json.Marshal(validEdit)

	cases := []UserCaseWithBody {
		 {
		 	name : "Valid Edit",
		 	method: "PUT",
			Path : "/users/1",
			expectedCode: http.StatusOK,
			requestBody: string(marshalledValidEdit),
			message:"User has been updated!"},
		{
		 	name : "Invalid Edit due to wrong user Id",
		 	method: "POST",
			Path : "/users/54361",
			expectedCode: http.StatusBadRequest,
			requestBody: string(marshalledValidEdit),
			message: "Wrong User Id"}}

	for _, testcase := range cases {
		req := httptest.NewRequest("PUT", "/", strings.NewReader(testcase.requestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testcase.Path)
		c.SetParamNames("id")
		paramValues := strings.Split(testcase.Path, "/")
		c.SetParamValues(paramValues[2])

		if assert.NoError(t, EditUserController(c)) {
			assert.Equal(t, testcase.expectedCode, rec.Code)

			var userResponse UserResponse

			if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
				assert.Error(t, err, "error")
			}
			
			assert.Equal(t, testcase.message, userResponse.Message)
			if testcase.name == "Valid edit" {
				assert.Equal(t, validEdit.Name, userResponse.User.Name)
				assert.Equal(t, validEdit.Email, userResponse.User.Email)
			}				
		}
	}
}

func Test_DeleteUserController(t *testing.T) {
	e := initConfigTest()

	AddUser("kevin", "kevin@gmail.com", "1234")

	cases := []UserCaseWithBody {
		 {
		 	name : "Valid Delete",
		 	method: "DELETE",
			Path : "/users/1",
			expectedCode: http.StatusOK,
			requestBody: "",
			message:"User has been deleted!"},
		{
		 	name : "Invalid Delete due to wrong user Id",
		 	method: "DELETE",
			Path : "/users/54361",
			expectedCode: http.StatusBadRequest,
			requestBody: "",
			message: "Wrong User Id"}}

	for _, testcase := range cases {
		req := httptest.NewRequest("DELETE", "/", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testcase.Path)
		c.SetParamNames("id")
		paramValues := strings.Split(testcase.Path, "/")
		c.SetParamValues(paramValues[2])

		if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, testcase.expectedCode, rec.Code)

			var userResponse UserResponse

			if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testcase.message, userResponse.Message)
		}
	}
}