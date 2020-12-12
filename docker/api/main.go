package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"eleuth/waco/service"
	"eleuth/waco/service/webhook"
)

func main() {
	var err error
	webhook.WebhookRules, err = webhook.ImportWebhookRules()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("WebhookRule: ", webhook.WebhookRules)

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
