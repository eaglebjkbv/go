package main

import "fmt"

type kisi struct {
	ad    string
	soyad string
}

// kural : receiver ile fonksiyon aynı program kodu içerisinde olmalıdır.
// go dili receiver ile method faklı kodlarda olmasına izin vermez.
// func (receiver) methodadı(parametlereler)(dönüştipleri) şeklinde kullanıldığında fonksiyona method denir.
func (k kisi) tamAdi() string {
	return k.ad + " " + k.soyad
}

// pointer işaretini kaldırarak denendiğinde ikinci fonksiyondaki soyad değişmiyor
// pointer kullanırsak k.soyad="siyah olarak diğer fonksiyona etki ediyor."
func (k *kisi) soyadAd() string {
	k.soyad = "Siyah"
	return k.soyad + " " + k.ad
}
func (k *kisi) adSoyad() string {
	return k.ad + k.soyad
}

type deneme int

func (d deneme) denemefonsiyonu() {
	fmt.Println("Merhaba", d)
}

func main() {
	k := kisi{
		ad:    "Kartal",
		soyad: "Kara",
	}
	fmt.Println(k.tamAdi())

	k2 := kisi{}
	k2.ad = "Şahin"
	k2.soyad = "Kara"
	tamAd := k2.tamAdi()
	fmt.Println(tamAd)
	// Pointer Kullnaıldığında
	k3 := kisi{
		ad:    "Atmaca",
		soyad: "Atar",
	}
	fmt.Println(k3.soyadAd())
	fmt.Println(k3.adSoyad())

	// integer type method (non-struct type)
	den := deneme(1)
	den.denemefonsiyonu()
}
