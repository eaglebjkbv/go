package main

import "fmt"

// takimAdi:="Besiktas" ....Bu tür var ya da const kullnaılmadan yapılan atamalar global tanımlanamıyor.

func main() {

	var isim string = "Bülent"
	var yas int32 = 42

	const havaSicak = true
	takimAdi := "Beşiktaş"

	fmt.Println(isim, yas, takimAdi)
	fmt.Printf("Yas Değişkenin Tip : %T \n", yas)
	fmt.Printf("havaSicak Değişkenin Tip : %T \n", havaSicak)
	fmt.Printf("takimAdi Değişkenin Tip : %T \n", takimAdi)

}
