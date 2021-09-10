package controllers

import (
	"altastore/config"

	"github.com/labstack/echo/v4"
)

func InitEchoTestAPI() *echo.Echo {

	config.InitDBTest()
	e := echo.New()
	return e
}
