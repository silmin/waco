package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllUsers(context echo.Context) error {
	db := ConnectDB()
	defer db.Close()

	users := []User{}
	db.Find(&users)

	fmt.Println("All Users:", users)

	json_data, _ := json.Marshal(users)
	return context.String(http.StatusOK, string(json_data))
}

func GetUser(context echo.Context) error {
	db := ConnectDB()
	defer db.Close()

	cardNo := context.Param("cardNo")

	user := User{}
	db.Find(&user, "card_no=?", cardNo)

	fmt.Println("Get User:", user)

	json_data, _ := json.Marshal(user)
	return context.String(http.StatusOK, string(json_data))
}

func RegisterUser(context echo.Context) error {
	db := ConnectDB()
	defer db.Close()

	user := new(User)
	if err := context.Bind(user); err != nil {
		return err
	}

	user.CardNo = context.Param("cardNo")

	db.Create(&user)

	fmt.Println("Create User:", user)

	json_data, _ := json.Marshal(user)
	return context.String(http.StatusOK, string(json_data))
}

func DeleteUser(context echo.Context) error {
	db := ConnectDB()
	defer db.Close()

	cardNo := context.Param("cardNo")
	db.Where("card_no=?", cardNo).Delete(&User{})

	fmt.Println("Delete User No:", cardNo)

	return context.String(http.StatusOK, string(""))
}
