package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MultiplyRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func CreateApp(handler AppHandler) *fiber.App {
	app := fiber.New()

	app.Get("/sum", func(c *fiber.Ctx) error {
		// Example URL: "/sum?x=5&y=4"
		x, _ := strconv.Atoi(c.Query("x"))
		y, _ := strconv.Atoi(c.Query("y"))

		r := handler.Sum(x, y)
		return c.Status(200).JSON(r)
	})

	app.Post("/multiply", func(c *fiber.Ctx) error {
		/*
			Example body payload:
			{
				"x": 5,
				"y": 4
			}
		*/
		req := new(MultiplyRequest)
		if err := c.BodyParser(req); err != nil {
			log.Println(err)
			return c.Status(400).SendString("Invalid payload")
		}

		r := handler.Multiply(req.X, req.Y)
		return c.Status(200).JSON(r)
	})

	return app
}

func main() {

	app := CreateApp(&AppHandlerStruct{})

	app.Listen(":3000")

}
