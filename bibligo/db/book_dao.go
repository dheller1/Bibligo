package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/dheller1/Bibligo/bibligo/core"
)

const tableName string = "books"

func getColumns() []string {
	return []string{
		"id       INTEGER not null primary key",
		"title    TEXT    not null",
		"authors  TEXT    not null",
		"subtitle TEXT",
		"edition  TEXT",
	}
}

func CreateTable(db *sql.DB) error {
	columnsExpr := strings.Join(getColumns(), ", ")
	statement := fmt.Sprintf("CREATE TABLE %s (%s);", tableName, columnsExpr)
	_, err := db.Exec(statement)
	return err
}

func QueryAllBooks(db *sql.DB) ([]core.Book, error) {
	result := make([]core.Book, 0)
	stmt := fmt.Sprintf("SELECT id, title, authors FROM %s", tableName)
	rows, err := db.Query(stmt)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var authorsText string
		err = rows.Scan(&id, &title, &authorsText)
		if err != nil {
			return result, err
		}
		authors := strings.Split(authorsText, ", ")

		book := core.Book{Title: title, Authors: authors}
		result = append(result, book)
	}
	return result, nil
}
