package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	kadi := r.FormValue("kadi")
	sifre := r.FormValue("sifre")

	if kadi == "deneme" && sifre == "1234" {
		tmpl.ExecuteTemplate(w, "login.html", kadi)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func main() {

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}
