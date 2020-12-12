package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func GetCurrentUsers(context echo.Context) error {
	db := ConnectDB()
	defer db.Close()

	currentUser := []CurrentUser{}
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

	currentUser := CurrentUser{CardNo: context.Param("cardNo")}
	if err := db.Create(&currentUser).Error; err != nil {
		return err
	}

	fmt.Println("Push User:", currentUser)

	json_data, _ := json.Marshal(currentUser)
	return context.String(http.StatusOK, string(json_data))
}

func DeleteCurrentUser(context echo.Context) error {
	db := ConnectDB()
	defer db.Close()

	cardNo := context.Param("cardNo")
	if err := db.Where("card_no=?", cardNo).Delete(&CurrentUser{}).Error; err != nil {
		return err
	}

	fmt.Println("Delete CurrentUser No:", cardNo)

	return context.String(http.StatusOK, string(""))
}
