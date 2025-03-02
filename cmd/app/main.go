package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=220021 dbname=testdb post=5432 sslmode=disable")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
