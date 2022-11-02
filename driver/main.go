package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/dheller1/Bibligo/bibligo/core"
	"github.com/dheller1/Bibligo/bibligo/db"
)

func main() {
	/*book := bibligo.MakeBook("Mastering Go", "Mihalis Tsoukalos")
	fmt.Println(book.String())

	err := book.Save()
	if err != nil {
		fmt.Println(err.Error())
	}

	pdfPath := `E:/Daten/Ebooks - Categorized/masteringgo.pdf`
	bibligo.MakeBookFromPdf(pdfPath)*/

	//db.CreateDb()
	//db.InsertIntoDb()
	//fmt.Println("OK")

	book := core.MakeBook("Mastering Go", "Mihalis Tsoukalos")
	fmt.Println(book.String())

	bibDb := db.OpenDb()
	books, err := db.QueryAllBooks(bibDb)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of books found: " + fmt.Sprint(len(books)))
	for i, v := range books {
		fmt.Printf("%03d. %s (%s)\n", i+1, v.Title, strings.Join(v.Authors, ", "))
	}
}
