package bibligo

import (
	"fmt"
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

func contains(slice []rune, char rune) bool {
	for _, v := range slice {
		if v == char {
			return true
		}
	}
	return false
}

func (b Book) makeFilename() string {
	ValidSpecialChars := []rune{'-', '_', ',', '(', ')'}
	validTitle := ""
	for _, c := range b.Title {
		if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || contains(ValidSpecialChars, c) {
			validTitle = validTitle + string(c)
		}
	}
	return validTitle
}

func MakeBook(title string, author string) *Book {
	return &Book{Title: title, Authors: []string{author}}
}
