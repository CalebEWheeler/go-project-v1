package person

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	Name string
	Age  string
}

func GetPeople(c *fiber.Ctx) {
	c.Send("All People")
}

func GetPerson(c *fiber.Ctx) {
	c.Send("A Person")
}

func CreatePerson(c *fiber.Ctx) {
	c.Send("Adds a Person")
}

func UpdatePerson(c *fiber.Ctx) {
	c.Send("Updates a Person")
}

func DeletePerson(c *fiber.Ctx) {
	c.Send("Delete a Person")
}
