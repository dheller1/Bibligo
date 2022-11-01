package main

import (
	"encoding/json"
	"fmt"

	"github.com/dheller1/Bibligo/bibligo"
)

func main() {
	book := bibligo.MakeBook("Mastering Go", "Mihalis Tsoukalos")
	fmt.Println(book.String())

	jsonCode, _ := json.Marshal(book)
	fmt.Println(string(jsonCode))
}
