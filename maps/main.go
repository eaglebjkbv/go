package main

import "fmt"

func main() {
	// Map oluşturma Yöntem 1
	ogrenciler := map[string]int{
		"Ahmet":  100,
		"Mehmet": 90,
		"Salih":  88,
	}
	ogrenciler["Deniz"] = 90
	delete(ogrenciler, "Ahmet")
	fmt.Println(ogrenciler)
	//Map Oluşturma Yöntem 2
	illereGorePlakalar := make(map[string]string)
	illereGorePlakalar = map[string]string{"izmir": "35", "Aydın": "09"}
	fmt.Println(illereGorePlakalar)
	// map i key value şeklinde listeletme
	for k, v := range ogrenciler {
		fmt.Println(k, v)
	}
	i := 0
	for _, v := range illereGorePlakalar {
		i++
		fmt.Println(i, "ilin plakası", v)
	}
	// Başka Bir şekilde Map oluşturma Yöntem 3
	// Nerelerde Kullanılır ??
	mp := map[[3]int]int{}
	mp[[3]int{1, 2, 3}] = 1
	fmt.Println(mp)
}
