package main

import (
	"fmt"

	"github.com/CalebEWheeler/go-project-v1/config"
	"github.com/CalebEWheeler/go-project-v1/database"
	"github.com/CalebEWheeler/go-project-v1/person"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func setupApp() *fiber.App {
	app := fiber.New()
	//can use - app.Use(authenticate.New(), attachuser.Handle, authorize.Handle)
	//then - authenticateduser.Setup(app)
	//if a user is needed in the future to modify RestAPI?

	return app
}

//setupRoutes will tie in methods found in person.go to be ran to the corresponding url path
func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/person", person.GetPeople)
	app.Get("/api/v1/person/:id", person.GetPerson)
	app.Post("/api/v1/person", person.CreatePerson)
	app.Put("/api/v1/person/:id", person.UpdatePerson)
	app.Delete("/api/v1/person/:id", person.DeletePerson)
}

//initDatabase will establish the initial database connection when the app is ran with gorm.Open() which will take in two params; database type, database credentials
func initDatabase() {
	var err error

	database.DBConn, err = gorm.Open("mysql", config.MySQLCredentials())
	if err != nil {
		panic("Failed to connect to the database")
	}
	fmt.Println("Database connection successfully opened")

	database.DBConn.AutoMigrate(&person.Person{})
	fmt.Println("Database Migrated")
}

func main() {
	app := setupApp()
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	//will establish the base route is localhost:8080
	// NEED TO CHANGE WHEN DEPLOYED
	app.Listen(8080)
}
