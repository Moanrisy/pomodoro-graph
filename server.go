package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func connectDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	password := os.Getenv("PASSWORD")

	// Open a connection to the database
	dbinfo := fmt.Sprintf("user=postgres password=%s host=db.fiakzcmolcxandnyppnp.supabase.co port=5432 dbname=postgres", password)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	// Try to ping the database
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database!")
	return db
}

func GetAllPomodoroActivity() {
	db := connectDB()
	rows, err := db.Query(`SELECT * FROM pomodoro`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	defer db.Close()
	var pomodoros pomodoro
	for rows.Next() {

		err = rows.Scan(&pomodoros.id, &pomodoros.inserted_at, &pomodoros.updated_at, &pomodoros.date, &pomodoros.id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(pomodoros)
	}
}

type pomodoro struct {
	id          int
	inserted_at string
	updated_at  string
	date        string
	counter     int
}

func main() {

	GetAllPomodoroActivity()

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
