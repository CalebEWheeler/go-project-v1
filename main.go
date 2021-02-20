package main

import (
	"database/sql"
	"fmt"

	"github.com/CalebEWheeler/go-project-v1/config"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("Go MySQL Integration")

	db, err := sql.Open("mysql", config.MySQLCredentials())

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	results, err := db.Query("SELECT name FROM users")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User

		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.Name)
	}

}
