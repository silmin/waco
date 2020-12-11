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

	e.GET("/users", service.GetAllUsers)
	e.GET("/users/:cardNo", service.GetUser)
	e.POST("/users/:cardNo", service.RegisterUser)
	e.DELETE("/users/:cardNo", service.DeleteUser)

	e.GET("/currents", service.GetCurrentUsers)
	e.PUT("/currents/:cardNo", service.PushCurrentUser)
	e.DELETE("/currents/:cardNo", service.DeleteCurrentUser)

	e.Start(":80")
}
