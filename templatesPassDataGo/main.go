package main

import (
	"html/template"
	"net/http"
)

type Ogrenci struct {
	Id      int
	Ad      string
	Soyad   string
	Notlari []Notlar
	Telno   []string
}
type Notlar struct {
	DersAdi  string
	DersNotu int
}

var ogr Ogrenci
var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}
func index(w http.ResponseWriter, r *http.Request) {

	ogr = Ogrenci{
		Id:    10,
		Ad:    "Michael",
		Soyad: "Jackson",
		Notlari: []Notlar{
			Notlar{
				DersAdi:  "Matematik",
				DersNotu: 50,
			},
			Notlar{
				DersAdi:  "Türkçe",
				DersNotu: 80,
			},
		},
		Telno: []string{"123123", "34535345", "475675765"},
	}
	tmpl.ExecuteTemplate(w, "index.html", ogr)
}

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

// Template ile ilgil  web Sayfası: https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet
