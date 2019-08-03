package main

import (
	"fmt"
)

func deferDeneme() {

	defer fmt.Println("Fonksiyon Defer Merhaba")
	fmt.Println("Fonksiyonun içi")

}

func main() {
	deferDeneme()
	fmt.Println("Main Fonksiyonundan Mesaj")

}
