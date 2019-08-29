package main

import (
	"fmt"

	"github.com/eaglebjkbv/goytb/exportedStructGo/model"
)

func main() {
	p := model.Kisi{
		Adi:      "Suzuki",
		Soyadi:   "Yamanguchi",
		Yasi:     54,
		Hobileri: []string{"Koşmak", "Okumak", "Müzik Dinlemek"},
	}

	fmt.Println(p)
	fmt.Println("Kişinin Tam Adı :", p.TamAdi())

	d := model.Dene{Adi: "Merhaba"}

	fmt.Println(d)

}
