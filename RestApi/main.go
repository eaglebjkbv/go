package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Ogrenci struct {
	ID      string  `json:"id"`
	Num     string  `json:"num"`
	Ad      string  `json:"ad"`
	Soyad   string  `json:"soyad"`
	Notlari *Notlar `json:"notlari"`
}
type Notlar struct {
	DersAdi string `json:"dersadi"`
	Notu    string `json:"notu"`
}

var ogrenciler []Ogrenci

func main() {
	rt := mux.NewRouter()

	ogrenciler = append(ogrenciler, Ogrenci{ID: "1", Num: "1", Ad: "Bülent", Soyad: "Vardal", Notlari: &Notlar{DersAdi: "Matematik", Notu: "100"}})
	ogrenciler = append(ogrenciler, Ogrenci{ID: "2", Num: "2", Ad: "Burak", Soyad: "Vardal", Notlari: &Notlar{DersAdi: "Matematik", Notu: "100"}})
	ogrenciler = append(ogrenciler, Ogrenci{ID: "3", Num: "3", Ad: "Serra", Soyad: "Vardal", Notlari: &Notlar{DersAdi: "Matematik", Notu: "100"}})
	// Tüm öğrencileri getir
	rt.HandleFunc("/ogrenciler", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ogrenciler)
	}).Methods("GET")

	// Sadece Bir öğrenci getir.
	rt.HandleFunc("/ogrenciler/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		for _, item := range ogrenciler {
			if item.ID == params["id"] {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
		// Boş JSON döndürmek için
		json.NewEncoder(w).Encode(&Ogrenci{})
	}).Methods("GET")

	// rt.HandleFunc("ogrenciler", ogrenciOlustur).Methods("POST")
	// rt.HandleFunc("ogrenciler", ogrenciGuncelle).Methods("PUT")
	// rt.HandleFunc("ogrenciler", ogrenciSil).Methods("DELETE")

	http.ListenAndServe(":80", rt)

}
