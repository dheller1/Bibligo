package db

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDb() *sql.DB {
	db, err := sql.Open("sqlite3", "bibligo.sqlite")
	if err != nil {
		log.Panic(err)
	}
	return db
}

func InsertIntoDb() {
	db := OpenDb()

	tx, _ := db.Begin()
	statement, _ := tx.Prepare("INSERT INTO data(id, content) values(?, ?)")
	defer statement.Close()

	for i := 0; i < 22; i++ {
		_, err := statement.Exec(i+22, strings.Repeat("b", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
}

func CreateDb() {
	db := OpenDb()

	statement := "CREATE TABLE data (id TEXT not null primary key, content TEXT);"
	_, err := db.Exec(statement)
	if err != nil {
		log.Panic(err)
	}
}
