package person

import (
	"github.com/CalebEWheeler/go-project-v1/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	Name string
	Age  string
}

func GetPeople(c *fiber.Ctx) {
	db := database.DBConn
	var people []Person
	db.Find(&people)
	c.JSON(people)
}

func GetPerson(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var person Person
	db.Find(&person, id)
	c.JSON(person)
}

func CreatePerson(c *fiber.Ctx) {
	db := database.DBConn
	var person Person
	person.Name = "Caleb"
	person.Age = "26"

	db.Create(&person)
	c.JSON(person)
}

func UpdatePerson(c *fiber.Ctx) {
	c.Send("Updates a Person")
}

func DeletePerson(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var person Person
	db.First(&person, id)
	if person.Name == "" {
		c.Status(500).Send("No person found with given ID")
	}
	db.Delete(&person)
	c.Send("Person successfully deleted")
}
