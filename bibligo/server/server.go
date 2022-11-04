package server

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/dheller1/Bibligo/bibligo/db"
)

var db_connection = db.OpenDb()

var all_templates *template.Template
var dataDir string

// var valid_path = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	dataDir = path.Join(home, ".bibligo")
	all_templates = template.Must(template.ParseFiles(
		path.Join(dataDir, "templates", "add.html"),
		path.Join(dataDir, "templates", "list.html"),
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
	err := all_templates.ExecuteTemplate(w, "add.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Start(addr string) {
	http.HandleFunc("/list", frontPageHandler)
	http.HandleFunc("/add", addBookHandler)
	log.Fatal(http.ListenAndServe(addr, nil))
}
