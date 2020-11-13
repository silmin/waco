package service

import (
	"net/http"

	"github.com/labstack/echo"
)

func TopPage(context echo.Context) error {
	return context.String(http.StatusOK, "Hello World")
}
