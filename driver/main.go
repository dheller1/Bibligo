package main

import (
	"fmt"

	"github.com/dheller1/Bibligo/bibligo"
)

func main() {
	book := bibligo.MakeBook("Mastering Go", "Mihalis Tsoukalos")
	fmt.Println(book.String())

	err := book.Save()
	if err != nil {
		fmt.Println(err.Error())
	}
}
