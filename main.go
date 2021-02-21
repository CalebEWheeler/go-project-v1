package main

import (
	"github.com/CalebEWheeler/go-project-v1/person"
	"github.com/gofiber/fiber"
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

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(8080)
}
