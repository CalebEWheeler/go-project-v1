package person

import (
	"io/ioutil"
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
		c.Body()
	})

	req := httptest.NewRequest("GET", "/api/v1/person/3", nil)
	//name = Caleb
	//age = 26

	resp, _ := app.Test(req)

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		// utils.AssertEqual(t, nil, err, "app.Test")
		utils.AssertEqual(t, "27", body, "test")
		utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
	}

	// utils.AssertEqual(t, 400, resp.StatusCode, "OK response is expected.")
}

func TestCreatePerson(t *testing.T) {
	// app := fiber.New()

	// person := &Person{
	// 	Name: "bob",
	// 	Age:  22,
	// }

	// jsonPerson, _ := json.Marshal(person)

	// app.Post("/api/v1/person", func(c *fiber.Ctx) {
	// 	// c.SendStatus(200)
	// 	person := new(Person)
	// 	person.Name = "bob"
	// 	person.Age = 22

	// 	if err := c.BodyParser(person); err != nil {
	// 		fmt.Println("error = ", err)
	// 		c.SendStatus(200)
	// 	}

	// 	fmt.Println("person = ", person)
	// 	fmt.Println("person.Name = ", person.Name)
	// 	fmt.Println("person.Age = ", person.Age)

	// 	c.SendString(person)
	// })

	// resp, err := app.Test(httptest.NewRequest("POST", "/api/v1/person", bytes.NewBuffer(jsonPerson)))

	// utils.AssertEqual(t, nil, err, "app.Test")
	// utils.AssertEqual(t, person.Name, jsonPerson, "Expected and Actual name values are equal.")
	// utils.AssertEqual(t, 22, person.Age, "Expected and Actual Age values are equal.")
	// utils.AssertEqual(t, 200, resp.StatusCode, "OK response is expected")
}

func TestUpdatePerson(t *testing.T) {

}

func TestDeletePerson(t *testing.T) {

}
