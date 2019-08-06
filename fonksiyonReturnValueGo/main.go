package main

import "fmt"

func main() {
	toplam := topla(10, 20)
	fmt.Println("İşlem Sonucu : ", toplam)

	carpim := carpma(10, 20)
	fmt.Println("Çarpım Sonucu :", *carpim)

	fark := cikarma(9, 8)
	fmt.Println("Çıkarma Sonucu : ", fark)

	bolum := bolme(10, 2)
	fmt.Println("Bölüm Sonucu :", bolum)

	bolme2, err := bolme2(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Bölme Sonucu : ", bolme2)

}

func topla(a, b int) int {
	return a + b
}

func carpma(a, b int) *int {
	carpim := a * b
	return &carpim
}
func cikarma(a, b int) (sonuc int) {
	sonuc = a - b
	return
	// return yanına herhangi bir şey yazmaya gerek yok çünkü sonuc döndürülecek
}

func bolme(a, b float64) float64 {
	return a / b
}
func bolme2(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, fmt.Errorf("Bir Sayı Sıfıra Bölünemez!\n")
	}
	return a / b, nil
}
