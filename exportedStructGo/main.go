package main

import "github.com/eaglebjkbv/go_Examples/exportedStructGo/model"
import "fmt"

func main() {
	k := model.Kisi{
		Adi:      "Suzuki",
		Soyadi:   "Yamanichi",
		Yasi:     65,
		Hobileri: []string{"Okumak", "Koşmak"},
	}
	fmt.Println(k)

	fmt.Println(k.TamAdi())

}
