package middlewares

import (
	"net/http"
	"fmt"
	"strings"
	"strconv"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt"
)

func AuthenticateUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer") {
			return c.JSON(http.StatusBadRequest, struct {
				Status  string
				Message string
			}{Status: "Failed", Message: "Invalid Authorization type"})
		}
		
		tokenArr := strings.Split(authHeader, " ")

		if len(tokenArr) != 2 {
			return c.JSON(http.StatusBadRequest, struct {
				Status  string
				Message string
			}{Status: "Failed", Message: "Invalid JWT Format!"})
		}

		tokenString := strings.Split(tokenArr[1],".")

		if len(tokenString) != 3 {
			return c.JSON(http.StatusBadRequest, struct {
				Status  string
				Message string
			}{Status: "Failed", Message: "Invalid JWT Format!"})
		}

		valid, id, _ := checkToken(tokenArr[1])
		// id can be either float64 or int. In any case, its numeric type so its save to 
		// ignore if the assertion is failed and just convert it to int when we set it to header
		userId, _ := id.(float64)

		if !valid {
			return c.JSON(http.StatusBadRequest, struct {
				Status  string
				Message string
			}{Status: "Failed", Message: "JWT is invalid!"})
		}

		c.Request().Header.Set("userId", strconv.Itoa(int(userId)))

		return next(c)
		}
}

func CheckId(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		paramId, _ := strconv.Atoi(c.Param("id"))
		userId,_ := strconv.Atoi(c.Request().Header.Get("userId"))
		if  userId != paramId {
			return c.JSON(http.StatusUnauthorized, struct {
				Status  string
				Message string
			}{Status: "Failed", Message: "You do not have permission to access this resource!"})
		}
		return next(c)
	}
}

func checkToken(tokenString string) (bool, interface{}, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, claims["userId"], nil
	}

	return false, -1, err
}
