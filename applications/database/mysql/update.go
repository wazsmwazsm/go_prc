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
	sql := "UPDATE t_services_tree_user_role_domain SET domain = ? WHERE id = ?"
	res, err := db.Exec(sql, 11111, 784)
	if err != nil {
		log.Fatal(err)
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(affectedRows)
}
