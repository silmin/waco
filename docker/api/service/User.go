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
	if err := db.Find(&users).Error; err != nil {
		return err
	}

	fmt.Println("All Users:", users)

	json_data, _ := json.Marshal(users)
	return context.String(http.StatusOK, string(json_data))
}

func GetUser(context echo.Context) error {
	db := ConnectDB()
	defer db.Close()

	cardNo := context.Param("cardNo")

	user := User{}
	if err := db.Find(&user, "card_no=?", cardNo).Error; err != nil {
		return err
	}

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

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	fmt.Println("Create User:", user)

	json_data, _ := json.Marshal(user)
	return context.String(http.StatusOK, string(json_data))
}

func DeleteUser(context echo.Context) error {
	db := ConnectDB()
	defer db.Close()

	cardNo := context.Param("cardNo")
	if err := db.Where("card_no=?", cardNo).Delete(&User{}).Error; err != nil {
		return err
	}

	fmt.Println("Delete User No:", cardNo)

	return context.String(http.StatusOK, string(""))
}
