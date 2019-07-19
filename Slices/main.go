package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60}

	s = s[:3]
	fmt.Println(fmt.Sprintf("Uzunluk=%d Kapasite=%d", len(s), cap(s)))
	fmt.Println(s)

	s = s[:2]
	fmt.Printf("Dizi Uzunluğu =%d Dizi kapasitesi =%d \n", len(s), cap(s))
	fmt.Println(s)

	var a []int
	fmt.Println(a, len(a), cap(a))
	if a == nil {
		fmt.Println("sıfır")
	}

}
