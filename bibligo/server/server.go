package server

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/dheller1/Bibligo/bibligo/core"
	"github.com/dheller1/Bibligo/bibligo/db"
)

var db_connection = db.OpenDb()

var all_templates *template.Template
var dataDir string

var valid_path = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	dataDir = path.Join(home, ".bibligo")
	all_templates = template.Must(template.ParseFiles(
		path.Join(dataDir, "templates", "edit.html"),
		path.Join(dataDir, "templates", "list.html"),
		path.Join(dataDir, "templates", "view.html"),
	))
}

func frontPageHandler(w http.ResponseWriter, r *http.Request) {
	books, err := db.QueryAllBooks(db_connection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = all_templates.ExecuteTemplate(w, "list.html", books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func addBookHandler(w http.ResponseWriter, r *http.Request) {
	// edit template serves dual purpose: if the data is nil, a new entry will be added
	err := all_templates.ExecuteTemplate(w, "edit.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type editBookInfo struct {
	core.Book
	AuthorsText string
}

func makeBookInfo(b core.Book) editBookInfo {
	return editBookInfo{b, strings.Join(b.Authors, ", ")}
}

func editBookHandler(w http.ResponseWriter, r *http.Request) {
	m := valid_path.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}

	bookId, err := strconv.Atoi(m[2]) // id is the second subexpression
	if err != nil {
		http.NotFound(w, r)
		return
	}
	book, err := db.QueryBook(db_connection, bookId)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = all_templates.ExecuteTemplate(w, "edit.html", makeBookInfo(*book))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func viewBookHandler(w http.ResponseWriter, r *http.Request) {
	m := valid_path.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}

	bookId, err := strconv.Atoi(m[2]) // id is the second subexpression
	if err != nil {
		http.NotFound(w, r)
		return
	}
	book, err := db.QueryBook(db_connection, bookId)

	err = all_templates.ExecuteTemplate(w, "view.html", book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Start(addr string) {
	http.HandleFunc("/add", addBookHandler)
	http.HandleFunc("/edit/", editBookHandler)
	http.HandleFunc("/list", frontPageHandler)
	http.HandleFunc("/view/", viewBookHandler)
	log.Fatal(http.ListenAndServe(addr, nil))
}
