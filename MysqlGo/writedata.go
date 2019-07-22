package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Merhaba")
	db, err := sql.Open("mysql", "bv2:12345@tcp(localhost:3306)/restproje")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO plcdeger (tarih, saat, deger) VALUES (?,?,?)")

	rows, err := stmt.Exec("2019.07.22", "16:40", "132")

	if err != nil {
		panic(err)
	}
	fmt.Println(rows)
}
