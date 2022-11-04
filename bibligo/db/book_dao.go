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
	stmt := fmt.Sprintf("SELECT id, title, authors, subtitle, edition FROM %s", tableName)
	rows, err := db.Query(stmt)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var authorsText string
		var subtitle string
		var edition string
		err = rows.Scan(&id, &title, &authorsText, &subtitle, &edition)
		if err != nil {
			return result, err
		}
		authors := strings.Split(authorsText, ", ")

		book := core.Book{Id: id, Title: title, Authors: authors, Subtitle: subtitle, Edition: edition}
		result = append(result, book)
	}
	return result, nil
}

func QueryBook(db *sql.DB, id int) (*core.Book, error) {
	stmt := fmt.Sprintf("SELECT title, authors, subtitle, edition FROM %s WHERE id=%d", tableName, id)
	row := db.QueryRow(stmt)

	var title string
	var authorsText string
	var subtitle string
	var edition string
	err := row.Scan(&title, &authorsText, &subtitle, &edition)
	if err != nil {
		return nil, err
	}
	authors := strings.Split(authorsText, ", ")

	book := core.Book{Id: id, Title: title, Authors: authors, Subtitle: subtitle, Edition: edition}
	return &book, nil
}
