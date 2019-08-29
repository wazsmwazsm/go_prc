package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "mrq:mafengwo@tcp(127.0.0.1:12345)/test?charset=utf8")
	if err != nil {
		log.Fatal("Db connect faild!")
	}
	defer db.Close()
	sql := "INSERT INTO t_services_tree_user_role_domain VALUES (null, ?, ?, ?, \"2019-08-7 09:13:12\")"
	res, err := db.Exec(sql, 777, 0, 170)
	if err != nil {
		log.Fatal(err)
	}
	lastIsID, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(lastIsID)
}
