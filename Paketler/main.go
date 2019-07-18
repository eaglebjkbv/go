package main

import "fmt"
import "github.com/eaglebjkbv/go_Examples/Paketler/islem"

func main() {
	fmt.Println(islem.Topla(10, 30))
	a := islem.Topla(10, 20)
	fmt.Printf("İşlem sonucunun tipi %T\n", a)
}
