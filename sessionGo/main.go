package main

import (
	"fmt"
	"os"

	"github.com/gorilla/sessions"

	"html/template"
	"net/http"
)

var kayit *sessions.CookieStore
var tmpl *template.Template

func init() {

	kayit = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	kayit.Options = &sessions.Options{
		MaxAge:   60 * 15,
		HttpOnly: true,
	}
	tmpl = template.Must(template.ParseGlob("templates/*"))
}
func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	session, err := kayit.Get(r, "yeni-oturum")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.Values["giris"] != "tamam" {
		if r.FormValue("kadi") == "kartal" && r.FormValue("sifre") == "1234" {

			session.Values["giris"] = "tamam"
			session.Save(r, w)
			tmpl.ExecuteTemplate(w, "login.html", nil)

		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	} else {
		tmpl.ExecuteTemplate(w, "login.html", nil)
	}
}
func diger(w http.ResponseWriter, r *http.Request) {
	session, err := kayit.Get(r, "yeni-oturum")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Hataaa")
		return
	}
	if session.Values["giris"] == "tamam" {
		tmpl.ExecuteTemplate(w, "diger.html", nil)

	}
}
func logout(w http.ResponseWriter, r *http.Request) {
	session, err := kayit.Get(r, "yeni-oturum")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["giris"] = nil
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/diger", diger)
	http.HandleFunc("/login", login)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}
