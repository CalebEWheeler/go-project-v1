package person

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/gofiber/utils"
)

//TOKEN for OAUTH purposes - const bearerToken = "Bearer o8SMWsPx.9wJ7nZ86.XUyQV3yu"
//place the following line under req declaration in methods
//	req.Header.Set("Authorization", bearerToken)

func TestGetPeople(t *testing.T) {
	app := fiber.New()

	app.Get("/api/v1/person", func(c *fiber.Ctx) {
		c.SendStatus(400)
	})

	resp, err := app.Test(httptest.NewRequest("GET", "/api/v1/person", nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, resp.Request.Method, "GET", "Request Method")
	utils.AssertEqual(t, 400, resp.StatusCode, "Status code")
}

func TestGetPerson(t *testing.T) {

}

func TestCreatePerson(t *testing.T) {
	app := fiber.New()

	person := &Person{
		Name: "bob",
		Age:  22,
	}

	jsonPerson, _ := json.Marshal(person)

	app.Post("/api/v1/person", func(c *fiber.Ctx) {
		c.SendStatus(200)
	})

	resp, err := app.Test(httptest.NewRequest("POST", "/api/v1/person", bytes.NewBuffer(jsonPerson)))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, "bob", person.Name, "Expected and Actual name values are equal.")
	utils.AssertEqual(t, 22, person.Age, "Expected and Actual Age values are equal.")
	utils.AssertEqual(t, 200, resp.StatusCode, "OK response is expected")
}

func TestUpdatePerson(t *testing.T) {

}

func TestDeletePerson(t *testing.T) {

}
