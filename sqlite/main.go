package main

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func seedDatabase(db *sql.DB) {
	sqlStatement := `
		create table if not exists users (id integer not null primary key, name text);
		delete from users;
	`

	_, err := db.Exec(sqlStatement)

	if err != nil {
		log.Printf("%q: %s\n", err, sqlStatement)

		return
	}
}

func createUser(db *sql.DB) {
	transaction, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	stmt, err := transaction.Prepare("insert into users(id, name) values(1, 'gugu')")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		log.Fatal(err)
	}

	err = transaction.Commit()

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	os.Remove("./database.db")

	db, err := sql.Open("sqlite", "./database.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	seedDatabase(db)

	createUser(db)
}
