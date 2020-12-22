package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"eleuth/waco/service/room_user"
	"eleuth/waco/service/webhook"
)

func ListUsers(context echo.Context) error {
	db := ConnectDB()
	defer db.Close()

	users := []room_user.User{}
	db.Find(&users)

	fmt.Println("Users:", users)

	json_data, _ := json.Marshal(users)
	return context.String(http.StatusOK, string(json_data))
}

func GetAllUsers(context echo.Context) error {
	db := ConnectDB()
	defer db.Close()

	users := []room_user.User{}
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

	user := room_user.User{}
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

	user := room_user.User{}
	if err := context.Bind(user); err != nil {
		return err
	}

	user.CardNo = context.Param("cardNo")

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	fmt.Println("Create User:", user)

	if errs := webhook.CallWebhook(webhook.RegisterUserEvent, user); errs != nil {
		fmt.Println(errs)
		return errs[0]
	}

	json_data, _ := json.Marshal(user)
	return context.String(http.StatusOK, string(json_data))
}

func DeleteUser(context echo.Context) error {
	db := ConnectDB()
	defer db.Close()

	cardNo := context.Param("cardNo")
	user := room_user.User{}
	if err := db.Find(&user, "card_no=?", cardNo).Error; err != nil {
		return err
	}

	if err := db.Where("card_no=?", cardNo).Delete(&room_user.User{}).Error; err != nil {
		return err
	}

	fmt.Println("Delete User No:", cardNo)

	if errs := webhook.CallWebhook(webhook.DeleteUserEvent, user); errs != nil {
		fmt.Println(errs)
		return errs[0]
	}

	return context.String(http.StatusOK, string(""))
}

func GetCurrentUsers(context echo.Context) error {
	db := ConnectDB()
	defer db.Close()

	currentUser := []room_user.CurrentUser{}
	if err := db.Find(&currentUser).Error; err != nil {
		return err
	}

	fmt.Println("CurrentUsers: ", currentUser)

	json_data, _ := json.Marshal(currentUser)
	return context.String(http.StatusOK, string(json_data))
}

func PushCurrentUser(context echo.Context) error {
	db := ConnectDB()
	defer db.Close()

	cardNo := context.Param("cardNo")
	currentUser := room_user.CurrentUser{CardNo: cardNo}
	if err := db.Create(&currentUser).Error; err != nil {
		return err
	}

	user := room_user.User{}
	if err := db.Find(&user, "card_no=?", cardNo).Error; err != nil {
		return err
	}

	fmt.Println("Push User:", currentUser)

	if errs := webhook.CallWebhook(webhook.PushCurrentUserEvent, user); errs != nil {
		fmt.Println(errs)
		return errs[0]
	}

	json_data, _ := json.Marshal(currentUser)
	return context.String(http.StatusOK, string(json_data))
}

func PopCurrentUser(context echo.Context) error {
	db := ConnectDB()
	defer db.Close()

	user := room_user.User{}
	cardNo := context.Param("cardNo")
	if err := db.Find(&user, "card_no=?", cardNo).Error; err != nil {
		return err
	}

	if err := db.Where("card_no=?", cardNo).Delete(&room_user.CurrentUser{}).Error; err != nil {
		return err
	}

	fmt.Println("Pop CurrentUser No:", cardNo)

	if errs := webhook.CallWebhook(webhook.PopCurrentUserEvent, user); errs != nil {
		fmt.Println(errs)
		return errs[0]
	}

	return context.String(http.StatusOK, string(""))
}
