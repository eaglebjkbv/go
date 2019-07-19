package main

import "fmt"

func main() {
	dizi := make([]int, 10)
	for i := range dizi {
		dizi[i] = 1 << uint(i)
	}

	for _, deger := range dizi {
		fmt.Printf("Deger =%d\n", deger)
	}

}
