package bibligo

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

type Book struct {
	Title    string
	Subtitle string
	Authors  []string
	Edition  string
}

func (b Book) String() string {
	return fmt.Sprintf("%s (%s)", b.Title, strings.Join(b.Authors, ", "))
}

func (b Book) Save() error {
	if !pathExists("records") {
		os.Mkdir("records", os.ModePerm)
	}
	fn := path.Join("records", b.makeFilename())
	json, err := json.Marshal(b)
	if err != nil {
		return err
	}
	return os.WriteFile(fn, json, 0600)
}

func contains(slice []rune, char rune) bool {
	for _, v := range slice {
		if v == char {
			return true
		}
	}
	return false
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (b Book) makeFilename() string {
	ValidSpecialChars := []rune{'-', '_', ',', '(', ')'}
	validTitle := ""
	for _, c := range b.Title {
		if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || contains(ValidSpecialChars, c) {
			validTitle = validTitle + string(c)
		}
	}
	return validTitle + ".bkr"
}

func MakeBook(title string, author string) *Book {
	return &Book{Title: title, Authors: []string{author}}
}
