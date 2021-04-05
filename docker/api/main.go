package main

import (
	"log"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"eleuth/waco/service"
	"eleuth/waco/service/registry"
	"eleuth/waco/service/webhook"
)

func main() {
	var err error
	for {
		if err = registry.CreateSchema(); err != nil {
			log.Println(err)
			log.Println("mysql is unavailable - sleeping")
			time.Sleep(time.Second)
			log.Println("reconnecting to mysql ...")
		} else {
			break
		}
	}

	webhook.WebhookRules, err = webhook.ImportWebhookRules()
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, rule := range webhook.WebhookRules {
		log.Println("WebhookRule: ", rule)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", service.GetAllUsers)
	e.GET("/users/:cardNo", service.GetUser)
	e.POST("/users/:cardNo", service.RegisterUser)
	e.PUT("/users/:cardNo", service.UpdateUser)
	e.DELETE("/users/:cardNo", service.DeleteUser)

	e.GET("/currents", service.GetCurrentUsers)
	e.PUT("/currents/:cardNo", service.PushCurrentUser)
	e.DELETE("/currents/:cardNo", service.PopCurrentUser)

	e.Start(":80")
}
