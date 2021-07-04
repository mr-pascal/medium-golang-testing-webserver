package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Request body payload of the 'POST /multiply' endpoint
type MultiplyRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func CreateApp(handler AppHandler) *fiber.App {
	// Create new instance
	app := fiber.New()

	// Add new 'GET /sum' endpoint
	app.Get("/sum", func(c *fiber.Ctx) error {
		// Example URL: "/sum?x=5&y=4"

		// Get values from query parameters
		x, _ := strconv.Atoi(c.Query("x"))
		y, _ := strconv.Atoi(c.Query("y"))

		// Use "Sum" handler to calculate the sum of both values
		r := handler.Sum(x, y)

		// Return '200 OK' with a 'Result' object containing the calculated value
		return c.Status(200).JSON(r)
	})

	app.Post("/multiply", func(c *fiber.Ctx) error {
		/*
			Example request body payload:
			{
				"x": 5,
				"y": 4
			}
		*/

		// Try to parse request body to 'req' object
		req := new(MultiplyRequest)
		if err := c.BodyParser(req); err != nil {
			log.Println(err)

			// Return '400 Bad Request' with a text message
			return c.Status(400).SendString("Invalid payload")
		}

		// Use "Mutliply" handler to calculate the multiplication of both values
		r := handler.Multiply(req.X, req.Y)

		// Return '200 OK' with a 'Result' object containing the calculated value
		return c.Status(200).JSON(r)
	})

	return app
}

func main() {

	// Create the App
	app := CreateApp(&AppHandlerStruct{})

	// Listne to Port 3000
	app.Listen(":3000")

}
