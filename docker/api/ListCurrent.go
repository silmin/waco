package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func CurrentUsers(context echo.Context) error {
	db := ConnectDB()
	defer db.Close()

	users := []User{}
	db.Find(&users)

	fmt.Println("Users:", users)

	json_data, _ := json.Marshal(users)
	return context.String(http.StatusOK, string(json_data))
}
