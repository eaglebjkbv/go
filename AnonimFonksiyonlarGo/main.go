package main

import "fmt"

func main() {

	// Anonim Fonksiyon basit tanımlama
	func() {
		fmt.Println("Selamlar")
	}() // "()"" Fonksiyon kendi kendini çağırır.

	for i := 0; i < 3; i++ {
		// Anonim Fonksiyon
		func() {
			fmt.Println("i Sayısı :", i)
		}()

	}
	for i := 0; i < 3; i++ {
		// Esas Yazım Şekli
		func(i int) {
			fmt.Println("-i Sayısı :", i)
		}(i) // fonksiyona aktarılacak değişken

	}
	// Bir tip olarak fonksiyon
	f := func(s string) {
		fmt.Println("Deneme", s)
	}
	f("merhaba")
	fmt.Printf("Tip : %T \n", f)
	// var ifadesi kullanarak fonksiyon tanımlama
	var fon func(s string)
	fon = func(s string) {
		fmt.Println("Var ile tanımlanmış fonksiyondan ", s)
	}
	fon("merhaba")

}
