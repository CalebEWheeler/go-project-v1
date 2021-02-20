package main

import (
	"database/sql"
	"fmt"

	"github.com/CalebEWheeler/go-project-v1/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Integration")

	db, err := sql.Open("mysql", config.MySQLCredentials())

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Successfully Connected to MySQL database with ignored config file")
}
