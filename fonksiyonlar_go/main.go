package main

import "fmt"

func main() {
	selamKelimesi := "Merhaba "
	selamlamaKisisi := "Bülent"
	merhabaDe(selamKelimesi, selamlamaKisisi) // Veriler gönderiliyor.
	fmt.Println(selamlamaKisisi)

	vedaKelimesi := "Hoşçakal "
	vedaKisisi := "Michael"
	vedaEt(&vedaKelimesi, &vedaKisisi) // Adres verileri gönderiliyor.
	fmt.Println(vedaKisisi)

}

// Pointer tip olmayan argüman alma ( Argümanları kopyalama)
func merhabaDe(selamKelimesi, selamlamaKisisi string) {
	fmt.Println(selamKelimesi, selamlamaKisisi)
	selamlamaKisisi = "Burak" // fonksiyon selamlamaKisisi değişkeni ile main fonksiyonundaki selamlamaKisisi farklı değişkenler
	// selamlamaKisisi main fonksiyondaki selmalamaKisisi değişkeninin farklı bir kopyasıdır.
}

// Pointer tipi argüman alma
func vedaEt(vedaKelimesi, vedaKisisi *string) {
	fmt.Println(*vedaKelimesi, *vedaKisisi)
	*vedaKisisi = "Goeorge" // vedaKisisi main fonksiyondaki vedaKisisi değişkeni ile aynı adresi göstermektedir
	// Bu nedenle burada yapılan değişiklik aynı adresi gösteren ana fonksiyondaki vedaKisisi değişkenini de değiştirir.
}
