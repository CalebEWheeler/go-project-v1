package person

import (
	"strconv"

	"github.com/CalebEWheeler/go-project-v1/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	Name string
	Age  int
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

	person := new(Person)
	if err := c.BodyParser(person); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&person)
	c.JSON(person)
}

func UpdatePerson(c *fiber.Ctx) {
	id := c.Params("id")
	name := c.Query("name")
	age, _ := strconv.Atoi(c.Query("age"))
	db := database.DBConn

	var person Person
	db.First(&person, id)
	person.Name = name
	person.Age = age

	db.Save(&person)
	c.JSON(person)
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
