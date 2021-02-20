package main

import (
	"database/sql"
	"fmt"

	"github.com/CalebEWheeler/go-project-v1/config"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"string"`
}

func main() {
	fmt.Println("Go MySQL Integration")

	db, err := sql.Open("mysql", config.MySQLCredentials())

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO users VALUES('ELLIOT')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	// results, err := db.Query("SELECT name FROM users")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// for results.Next() {
	// 	var user user

	// 	err = results.Scan(&user.Name) {
	// 		if err != nil {
	// 			panic(err.Error())
	// 		}

	// 		fmt.Println(user.Name)
	// 	}
	// }

	fmt.Println("Successfully Connected to MySQL database with ignored config file")
}
