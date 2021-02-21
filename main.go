package main

import (
	"fmt"

	"github.com/CalebEWheeler/go-project-v1/config"
	"github.com/CalebEWheeler/go-project-v1/database"
	"github.com/CalebEWheeler/go-project-v1/person"

	// _ "github.com/go-sql-driver/mysql"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/person", person.GetPeople)
	app.Get("/api/v1/person/:id", person.GetPerson)
	app.Post("/api/v1/person", person.CreatePerson)
	app.Put("/api/v1/:id", person.UpdatePerson)
	app.Delete("/api/v1/person/:id", person.DeletePerson)
}

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
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(8080)
}
