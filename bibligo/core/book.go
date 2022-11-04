package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/ledongthuc/pdf"
)

type Book struct {
	Id       int
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

func MakeBookFromPdf(filename string) error {
	pdf.DebugOn = true
	f, reader, err := pdf.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	outline := reader.Outline()
	fmt.Println("Outline:")
	fmt.Println(outline.Title)
	for k, v := range outline.Child {
		fmt.Printf("%03d. Outline('%s', %d children)\n", k, v.Title, len(v.Child))
	}

	//npages := reader.NumPage()

	for ipage := 1; ipage <= 3; ipage++ {
		page := reader.Page(ipage)
		if page.V.IsNull() {
			continue
		}
		fmt.Printf("Page %d:\n", ipage)
		var pageText string
		textElements := page.Content().Text
		for _, v := range textElements {
			pageText = pageText + v.S
		}
		fmt.Println(pageText)
		fmt.Println()
	}
	return nil
}
