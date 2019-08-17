package main

import "fmt"

type urun struct {
	id    int
	adi   string
	fiyat float64
}

// Value Receiver Method
func (u urun) FiyatArttir(oran float64) float64 {
	u.fiyat += u.fiyat * oran / 100
	return u.fiyat
}

// FiyatArttir :Ürünün fiyatını arttırır
// Fonksiyon Parametresi olarak struct
// Deneme Deneme
func FiyatArttir(u urun, oran float64) float64 {
	u.fiyat += u.fiyat * oran / 100
	return u.fiyat
}

// Pointer Receiver Method
func (u *urun) pFiyatArttir(oran float64) float64 {
	u.fiyat += u.fiyat * oran / 100
	return u.fiyat
}
func main() {
	urn := urun{
		id:    1,
		adi:   "Patates",
		fiyat: 2.5,
	}
	fmt.Println("Ürünün Fiyatı : ", urn.fiyat)
	fmt.Println("Ürünün Zamlı Fiyatı :", urn.FiyatArttir(5))
	fmt.Println("Ürünün Zamlı Fiyatı Fonksiyondan : ", FiyatArttir(urn, 5))

	fmt.Println("Ürünün Güncel Fiyatı :", urn.fiyat)

	fmt.Println("Ürünün Zamlı fiyatı :", urn.pFiyatArttir(5))
	fmt.Println("Ürünün Güncel Fiyatı", urn.fiyat)

}
