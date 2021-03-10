package registry

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func CreateSchema() error {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "room_status"

	con, err := sql.Open(DBMS, USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME)
	if err != nil {
		return err
	}

	create_table := `
	create table if not exists users (
		card_no     	char(15) not null primary key unique,
		display_name 	varchar(50) default null,
		full_name    	varchar(50) default null,
		email       	varchar(255) default null
	) default CHARSET=utf8 COLLATE=utf8_bin;
	`
	_, err = con.Exec(create_table)
	if err != nil {
		return err
	}
	fmt.Println("create 'users' table")

	create_table = `
	create table if not exists current_users (
		card_no     	char(15) not null primary key unique,
		date_touched 	datetime null default current_timestamp,
		index idx_current(card_no),
		foreign key fk_current(card_no) references users(card_no)
	) default CHARSET=utf8 COLLATE=utf8_bin;
	`
	_, err = con.Exec(create_table)
	if err != nil {
		return err
	}
	fmt.Println("create 'current_users' table")

	return nil
}
