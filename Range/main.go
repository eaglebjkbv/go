package main

import "fmt"

var (
	dilim = []int{10, 20, 30, 40, 50, 60}
	sayi  = "Merhaba"
)

func main() {

	for i, v := range dilim {

		fmt.Printf("indis :%d Değer :%d \n", i, v)
	}
	fmt.Println(sayi)

}
