package routes

import (
	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func New() *echo.Echo {
	e = echo.New()

	registerRootMiddlewares()

	registerUserRoutes()

	registerProductRoutes()

	registerCheckoutRoutes()

	registerLoginRoutes()

	registerCartRoutes()
	
	return e
}