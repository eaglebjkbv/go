package main

import "fmt"

type Kisi struct {
	ad string
}

func deneme(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Println("Gönderilen Stringin içeriği : ", v)
	case int:
		fmt.Println("Gönderilen integer verisi ", v)
	case Kisi:
		fmt.Println("Kişi Adı :", v.ad)
	default:
		fmt.Println("Tanımlanamayan tip")
	}
}

func main() {
	k := Kisi{"Yoko"}
	deneme(k)
	deneme("Selamlar")
}
