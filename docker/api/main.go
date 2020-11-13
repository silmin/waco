package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"eleuth/waco/service"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", service.TopPage)

	e.Start(":80")
}
