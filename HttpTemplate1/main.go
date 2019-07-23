package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type mesaj struct {
	MesajNo     int
	MesajIcerik string
}

func index(w http.ResponseWriter, r *http.Request) {
	msj := mesaj{MesajNo: 1,
		MesajIcerik: "Naber Nasılsın"}
	fmt.Println(msj.MesajIcerik)
	tmpl := template.Must(template.ParseFiles("template/layout.gohtml"))
	tmpl.Execute(w, msj)

}

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}
