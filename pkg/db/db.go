package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", name)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS orders (id INTEGER PRIMARY KEY, name VARCHAR, status VARCHAR)")
	if err != nil {
		fmt.Println("Error creating table:", err)
		return nil, err
	}

	return db, nil
}

func InsertValue(db *sql.DB, name string, status string) (int64, error) {
	res, err := db.Exec("INSERT INTO orders (name, status) VALUES (?, ?)", name, status)
	if err != nil {
		fmt.Println("Error inserting data:", err)
		return 0, err
	}

	return res.LastInsertId()
}
