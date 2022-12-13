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
	var pomodoros Pomodoro
	for rows.Next() {

		err = rows.Scan(&pomodoros.Id, &pomodoros.Inserted_at, &pomodoros.Updated_at, &pomodoros.Date, &pomodoros.Counter)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(pomodoros)
	}
}

func PostIncrementPomodoro(pomodoro *Pomodoro) string {
	res := "Increment pomodoro succesfull"
	db := connectDB()

	insertStatement := `INSERT INTO pomodoro(date, counter) VALUES ($1, $2)`
	_, err := db.Exec(insertStatement, pomodoro.Date, pomodoro.Counter)
	if err != nil {
		// log.Fatal(err)
		log.Println(err)
		res = fmt.Sprintf("Increment pomodoro failed\n %s", err)
	}
	return res
}

func UpdatePomodoro(pomodoro *Pomodoro) (string, error) {
	res := "Update pomodoro counter successfull"
	db := connectDB()
	updateErr := error(nil)

	updateStatement := `UPDATE pomodoro SET counter = $1 WHERE date = $2`
	rows, err := db.Exec(updateStatement, pomodoro.Counter, pomodoro.Date)
	if err != nil {
		log.Println(err)
	}

	if zeroRowAffected, _ := rows.RowsAffected(); zeroRowAffected == 0 {
		updateErr = fmt.Errorf("Update pomodoro counter failed\n Pomodoro with date %s does not exist", pomodoro.Date)
	}

	return res, updateErr
}

type Pomodoro struct {
	Id          int
	Inserted_at string
	Updated_at  string
	Date        string `json:"date"`
	Counter     int    `json:"counter"`
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

	app.Post("/increment", func(c *fiber.Ctx) error {
		p := new(Pomodoro)

		if err := c.BodyParser(p); err != nil {
			log.Fatal(err)
		}

		res := PostIncrementPomodoro(p)

		return c.Send([]byte(res))
	})

	app.Put("/increment", func(c *fiber.Ctx) error {
		p := new(Pomodoro)

		if err := c.BodyParser(p); err != nil {
			log.Println(err)
		}

		res, err := UpdatePomodoro(p)

		fmt.Println(err)
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
