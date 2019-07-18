package main

import "fmt"

func main() {
	var meyvelerArr [2]string

	meyvelerArr[0] = "Elma"
	meyvelerArr[1] = "Armut"

	fmt.Println(meyvelerArr)
	fmt.Println(meyvelerArr[1])

	isimlerArr := [2]string{"Mehmet", "Ahmet"}

	fmt.Println(isimlerArr)

	hayvanlarSlice := []string{"Kedi", "Köpek", "Eşek", "Kurt"}

	fmt.Println(hayvanlarSlice)
	fmt.Println(len(hayvanlarSlice))
	fmt.Println(hayvanlarSlice[1:3])

}
