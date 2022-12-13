package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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
	// fmt.Println("Successfully connected to the database!")
	return db
}

func GetAllPomodoroActivity() []Pomodoro {
	db := connectDB()
	rows, err := db.Query(`SELECT * FROM pomodoro`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	defer db.Close()

	var pomodoros []Pomodoro

	for rows.Next() {
		var pomodoro Pomodoro

		err = rows.Scan(&pomodoro.Id, &pomodoro.Inserted_at, &pomodoro.Updated_at, &pomodoro.Date, &pomodoro.Counter)
		if err != nil {
			log.Fatal(err)
		}
		pomodoros = append(pomodoros, pomodoro)
	}
	return pomodoros
}

func AddPomodoro(pomodoro *Pomodoro) string {
	res := "Add pomodoro succesfull"
	db := connectDB()
	defer db.Close()

	insertStatement := `INSERT INTO pomodoro(date, counter) VALUES ($1, $2)`
	_, err := db.Exec(insertStatement, pomodoro.Date, 1)
	if err != nil {
		// log.Fatal(err)
		log.Println(err)
		res = fmt.Sprintf("Add pomodoro failed\n %s", err)
	}
	return res
}

func UpdatePomodoro(pomodoro *Pomodoro) (string, error) {
	res := "Update pomodoro counter successfull"
	db := connectDB()
	defer db.Close()
	updateErr := error(nil)

	selectStatement := `SELECT counter FROM pomodoro where date = $1`
	rows, err := db.Query(selectStatement, pomodoro.Date)
	if err != nil {
		log.Println(err)
	}

	var isPomodoroExist = false
	for rows.Next() {
		isPomodoroExist = true

		err = rows.Scan(&pomodoro.Counter)
		if err != nil {
			log.Println(err)
		}
	}

	if isPomodoroExist {
		pomodoro.Counter++

		updateStatement := `UPDATE pomodoro SET counter = $1 WHERE date = $2`
		_, err := db.Exec(updateStatement, pomodoro.Counter, pomodoro.Date)
		if err != nil {
			log.Println(err)
		}

	} else {
		// updateErr = fmt.Errorf("update pomodoro counter failed\n Pomodoro with date %s does not exist", pomodoro.Date)
		AddPomodoro(pomodoro)
	}

	return res, updateErr
}
