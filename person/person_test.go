package person

import (
	"net/http/httptest"
	"testing"

	"github.com/CalebEWheeler/go-project-v1/database"
	"github.com/gofiber/fiber"
	"github.com/gofiber/utils"
)

//TOKEN for OAUTH purposes - const bearerToken = "Bearer o8SMWsPx.9wJ7nZ86.XUyQV3yu"
//place the following line under req declaration in methods
//	req.Header.Set("Authorization", bearerToken)

func TestGetPeople(t *testing.T) {
	app := fiber.New()
	db := database.DBConn
	var people []Person
	db.Find(&people)

	app.Get("/api/v1/person", func(c *fiber.Ctx) {
		c.SendStatus(400)
		c.JSON(people)
	})

	resp, err := app.Test(httptest.NewRequest("GET", "/api/v1/person", nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, resp.Request.Method, "GET", "Request Method")
	utils.AssertEqual(t, `[{"ID":3,"CreatedAt":"2021-02-21T14:10:39-06:00","UpdatedAt":"2021-02-21T14:10:39-06:00","DeletedAt":null,"Name":"Caleb","Age":26},{"ID":5,"CreatedAt":"2021-02-21T16:50:30-06:00","UpdatedAt":"2021-02-21T16:50:30-06:00","DeletedAt":null,"Name":"amanda","Age":25},{"ID":6,"CreatedAt":"2021-02-23T14:40:54-06:00","UpdatedAt":"2021-02-23T14:40:54-06:00","DeletedAt":null,"Name":"jason","Age":26},{"ID":7,"CreatedAt":"2021-02-23T14:44:33-06:00","UpdatedAt":"2021-02-23T14:47:05-06:00","DeletedAt":null,"Name":"chase","Age":22}]`, resp.Body, "Response Body Check")
	utils.AssertEqual(t, 400, resp.StatusCode, "Status code")
}

func TestGetPerson(t *testing.T) {

}

func TestCreatePerson(t *testing.T) {

}

func TestUpdatePerson(t *testing.T) {

}

func TestDeletePerson(t *testing.T) {

}
