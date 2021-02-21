package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	//run with 'go run main.go
	//kill with 'ctrl+c'
	fmt.Fprintf(w, "Hello World")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/create", NewUserForm).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{id}", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/update", UpdateUserForm).Methods("GET")
	myRouter.HandleFunc("/user/update", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Go ORM Basics")

	InitialMigration()

	handleRequests()
}
