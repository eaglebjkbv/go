package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type mesaj struct {
	mesajNo     string
	mesajIcerik string
}

func index(w http.ResponseWriter, r *http.Request) {
	msj := mesaj{mesajNo: "Deneme",
		mesajIcerik: "Naber Nasılsın"}
	fmt.Println(msj.mesajIcerik)
	tmpl := template.Must(template.ParseFiles("template/layout.gohtml"))
	tmpl.Execute(w, msj)

}

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}
