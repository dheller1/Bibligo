package core

import (
	"errors"
	"fmt"
	"regexp"
)

// taken from https://www.oreilly.com/library/view/regular-expressions-cookbook/9781449327453/ch04s13.html
// but without lookahead, which is not supported by Go regexp...
var isbn_regex = regexp.MustCompile(`^(?:ISBN(?:-1[03])?:?\ )?(?:97[89][-\ ]?)?[0-9]{1,5}[-\ ]?[0-9]+[-\ ]?[0-9]+[-\ ]?[0-9X]$`)

// Represents an ISBN, which is either invalid (empty), or contains either 10 or 13 digits.
// Create ISBNs using Make() instead of constructing them manually.
type ISBN string

func (i ISBN) isIsbn13() bool {
	return len(i) == 13
}

// Returns the ISBN as a condensed string comprising 10 or 13 digits
// (e.g. 9781801070 or 9781801079310)
func (i ISBN) Short() string {
	return string(i)
}

// Returns the ISBN as a formatted long string with dash characters as separators
// (e.g. 978-1-80107-0 or 978-1-80107-931-0)
func (i ISBN) Long() string {
	s := string(i)
	if i.isIsbn13() {
		return fmt.Sprintf("%s-%s-%s-%s-%s", s[:3], string(s[3]), s[4:9], s[9:12], string(s[12]))
	} else {
		return fmt.Sprintf("%s-%s-%s-%s", s[:3], string(s[3]), s[4:9], string(s[9]))
	}
}

// Creates an ISBN from a string which may either be the full ISBN format
// with dashes, or the short format containing only digits.
func MakeISBN(s string) (ISBN, error) {
	if isbn_regex.MatchString(s) {
		// find non-digits (ISBN, dashes, and spaces) in the string and remove them
		// to only retain the pure numeric ISBN.
		nonIsbn := regexp.MustCompile("[- ]|^ISBN(?:-1[03])?:?")
		s = nonIsbn.ReplaceAllString(s, "")
		return ISBN(s), nil
	}
	return ISBN(""), errors.New("invalid format for an ISBN")
}
