package person

import (
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
	app := fiber.New()

	app.Get("/api/v1/person/5", func(c *fiber.Ctx) {
		c.SendStatus(400)
	})

	resp, err := app.Test(httptest.NewRequest("GET", "/api/v1/person/5", nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, resp.Request.Method, "GET", "Request Method")
	utils.AssertEqual(t, 400, resp.StatusCode, "Status code")
}

func TestCreatePerson(t *testing.T) {
	app := fiber.New()

	app.Post("/api/v1/person", func(c *fiber.Ctx) {
		c.SendStatus(200)
	})

	resp, err := app.Test(httptest.NewRequest("POST", "/api/v1/person", nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, "POST", resp.Request.Method, "Request Method")
	utils.AssertEqual(t, 200, resp.StatusCode, "OK response is expected")
}

func TestUpdatePerson(t *testing.T) {
	app := fiber.New()

	app.Put("/api/v1/person", func(c *fiber.Ctx) {
		c.SendStatus(200)
	})

	resp, err := app.Test(httptest.NewRequest("PUT", "/api/v1/person", nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, "PUT", resp.Request.Method, "Request Method")
	utils.AssertEqual(t, 200, resp.StatusCode, "OK response is expected")
}

func TestDeletePerson(t *testing.T) {
	app := fiber.New()

	app.Delete("/api/v1/person", func(c *fiber.Ctx) {
		c.SendStatus(200)
	})

	resp, err := app.Test(httptest.NewRequest("DELETE", "/api/v1/person", nil))

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, "DELETE", resp.Request.Method, "Request Method")
	utils.AssertEqual(t, 200, resp.StatusCode, "OK response is expected")
}
