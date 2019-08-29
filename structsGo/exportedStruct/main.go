package main

import "github.com/eaglebjkbv/goytb/exportedStruct/model"
import "fmt"

func main() {
	p := model.Kisi{
		Adi:      "Suzuki",
		Soyadi:   "Yamaguchi",
		Yasi:     54,
		Hobileri: []string{"Kitap Okumak", "Müzik Dinlemek", "Koşmak"},
	}

	fmt.Println(p)
	fmt.Println(p.TamAdi())

}
