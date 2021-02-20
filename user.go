package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/CalebEWheeler/go-project-v1/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() {
	db, err = gorm.Open("mysql", config.MySQLCredentials())
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", config.MySQLCredentials())
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	//returning in json format to the browser the extracted values from database
	//store it in a var and pass to html
	// json.NewEncoder(w).Encode(users)
	// var retrievedUsers = users

	parsedTemplate, _ := template.ParseFiles("static/users.html")
	err := parsedTemplate.Execute(w, users)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", config.MySQLCredentials())
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})

	fmt.Fprintf(w, "New User Successfully Created")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", config.MySQLCredentials())
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "User Successfully Deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", config.MySQLCredentials())
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email

	db.Save(&user)
	fmt.Fprintf(w, "Successfully Updated User")
}
