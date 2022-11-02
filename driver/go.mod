module github.com/dheller1/bibligo_test

go 1.19

replace github.com/dheller1/Bibligo/bibligo => ../bibligo

require github.com/dheller1/Bibligo/bibligo v0.0.0-00010101000000-000000000000

require (
	github.com/ledongthuc/pdf v0.0.0-20220302134840-0c2507a12d80 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
)
