package main

import "fmt"

func main() {
	islemf("carpma", 10, 20, 30, 40, 50)
}

// variadic fonksiyon
func islemf(islem string, a ...int) { // a adında bir slice oluşturuldu

	sonuc := 0
	if islem == "carpma" {
		sonuc = 1
	}
	for _, v := range a {
		if islem == "toplama" {
			sonuc += v
		} else if islem == "carpma" {
			sonuc *= v
		}
	}
	fmt.Println(islem, " ", sonuc)

}
