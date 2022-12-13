package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug"
)

type Pomodoro struct {
	Id          int
	Inserted_at string
	Updated_at  string
	Date        string `json:"date"`
	Counter     int    `json:"counter"`
}

func main() {

	// Create a new engine
	engine := pug.New("./views", ".pug")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {

		pomodoros := GetAllPomodoroActivity()

		// Render index
		return c.Render("index", fiber.Map{
			"Title":     "Hello, World!",
			"Pomodoros": pomodoros,
		}, "layouts/pomodoro")
	})

	app.Post("/increment", func(c *fiber.Ctx) error {
		p := new(Pomodoro)

		if err := c.BodyParser(p); err != nil {
			log.Fatal(err)
		}

		res := AddPomodoro(p)

		return c.Send([]byte(res))
	})

	app.Put("/increment", func(c *fiber.Ctx) error {
		p := new(Pomodoro)

		if err := c.BodyParser(p); err != nil {
			log.Println(err)
		}

		res, err := UpdatePomodoro(p)

		if err != nil {
			return c.Send([]byte(err.Error()))
		}

		// return c.Send([]byte(res))
		return c.Send([]byte(res))
	})

	app.Get("/layout", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"Title": "Layout!",
		}, "layouts/main")
	})

	log.Fatal(app.Listen(":3000"))
}
