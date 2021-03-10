package registry

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDB() *gorm.DB {
	dbHost := os.Getenv("MYSQL_HOST")

	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := fmt.Sprintf("tcp(%s:3306)", dbHost)
	DBNAME := "room_status"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
	return db
}
