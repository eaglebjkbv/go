package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("template/index.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}
func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}
