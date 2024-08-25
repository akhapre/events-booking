package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("count not connect database" + fmt.Sprintf(err.Error()))
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(3)
	DB = db

	createTable()
}

func createTable() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER NOT NULL
		)
	`

	if DB == nil {
		fmt.Println("db init fail")
	}
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("count not connect database" + fmt.Sprintf(err.Error()))
	}
}
