package registry

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDB() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "room_status"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
	return db
}
