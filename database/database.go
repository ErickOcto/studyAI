package database

import (
	"database/sql"
	"fmt"
)

func InitDatabase() *sql.DB {
	fmt.Println("Initializing database...")

	dsn := ""
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	return db
}
