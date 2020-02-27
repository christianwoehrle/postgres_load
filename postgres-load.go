package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	user := flag.String("user", "foo", "database user")
	password := flag.String("password", "foo", "database password")
	url := flag.String("url", "foo", "database url")
	database := flag.String("database", "postgres", "database name")
	param := flag.String("param", "sslmode=disable", "db-parameter like sslmode=disable")
	flag.Parse()

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?%s", *user, *password, *url, *database, *param)
	// + "?" + "sslmode=verify-full"
	for {
		insert(connStr)
	}
}

func insert(connStr string) {
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		log.Println("error getting connection: " + err.Error())
		return
	}
	for {
		t := time.Now()
		id := 0

		err = db.QueryRow("INSERT INTO id(created_at, updated_at) VALUES($1, $2 ) RETURNING id", t, t).Scan(&id)
		if err != nil {
			log.Println("Error inserting row")
			return
		} else {
			log.Println("Row %d inserted: %s", id, t.String())
		}
	}
}
