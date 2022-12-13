package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug"
)

func main() {

	// Create a new engine
	engine := pug.New("./views", ".pug")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public") 

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/pomodoro")
	})

	app.Get("/layout", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"Title": "Layout!",
		}, "layouts/main")
	})

	log.Fatal(app.Listen(":3000"))
}
