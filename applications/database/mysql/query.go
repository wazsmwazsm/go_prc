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
	rows, err := db.Query("SELECT user, role, domain FROM t_services_tree_user_role_domain WHERE user=?", 752)
	if err != nil {
		log.Fatal("Query faild!")
	}
	defer rows.Close()
	for rows.Next() {
		var user, role, domain int
		if err := rows.Scan(&user, &role, &domain); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("User: %v, Role: %v, Domain: %v\n", user, role, domain)
	}

	row := db.QueryRow("SELECT user, role, domain FROM t_services_tree_user_role_domain WHERE id=?", 597)
	var user, role, domain int
	if err := row.Scan(&user, &role, &domain); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User: %v, Role: %v, Domain: %v\n", user, role, domain)
}
