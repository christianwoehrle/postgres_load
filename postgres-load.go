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
	// for param see https://godoc.org/github.com/lib/pq
	param := flag.String("param", "sslmode=disable", "db-parameter like sslmode=disable")
	flag.Parse()

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?%s", *user, *password, *url, *database, *param)
	for {
		insert(connStr)
	}
}

func insert(connStr string) {

	log.Println("opening connection ...")
	db, err := sql.Open("postgres", connStr)
	log.Println("opened")
	defer func() {
		log.Println("closing connection ...")
		db.Close()
		log.Println("closed")
	}()
	if err != nil {
		log.Println("error getting connection: " + err.Error())
		return
	}
	for {
		t := time.Now()
		id := 0

		log.Print("about to insert row ... ")
		err = db.QueryRow("INSERT INTO id(created_at, updated_at) VALUES($1, $2 ) RETURNING id", t, t).Scan(&id)
		if err != nil {
			log.Println("  --> error inserting row")
			return
		} else {
			log.Printf(" --> row %d inserted: %s \n", id, t.String())
		}
	}
}
