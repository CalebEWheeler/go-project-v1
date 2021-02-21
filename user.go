package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/CalebEWheeler/go-project-v1/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type User struct {
	// gorm.Model
	ID        uint `gorm:"primaryKey"`
	Name      string
	Age       string
	CreatedAt time.Time
	UpdatedAt time.Time
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

	parsedTemplate, _ := template.ParseFiles("static/users.html")
	err := parsedTemplate.Execute(w, users)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}

func NewUserForm(w http.ResponseWriter, r *http.Request) {
	var user User

	parsedTemplate, _ := template.ParseFiles("static/createUser.html")
	err := parsedTemplate.Execute(w, user)
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

	// vars := mux.Vars(r)
	// name := vars["name"]
	// email := vars["email"]

	name := r.FormValue("name")
	age := r.FormValue("age")

	fmt.Printf("Name: %s | Age: %s", name, age)

	db.Create(&User{Name: name, Age: age})

	http.Redirect(w, r, "/users", 302)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", config.MySQLCredentials())
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	id := r.URL.Query().Get("id")
	fmt.Fprintf(w, "ID value from URL: %s", id)

	// var user User
	// db.Where("id = ?", id).Find(&user)
	// db.Delete(&user)

	// db.Delete(&user, id)

	// http.Redirect(w, r, "/users", 200)
}

func UpdateUserForm(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", config.MySQLCredentials())
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	// name := r.URL.Query().Get("name")
	id := r.URL.Query().Get("id")

	var user User
	db.Where("id = ?", id).Find(&user)

	parsedTemplate, _ := template.ParseFiles("static/updateUser.html")
	err := parsedTemplate.Execute(w, user)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	name := r.FormValue("name")
	age := r.FormValue("age")

	var user User
	db.Where("id = ?", id).Find(&user)

	user.Name = name
	user.Age = age

	db.Save(&user)
	// fmt.Fprintf(w, "Successfully Updated User")
}
