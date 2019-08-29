package main

import "fmt"

func main() {

	var hersey interface{} = "Merhaba"

	birString, bulundu := hersey.(float64) // type assertions (tip idiası)

	fmt.Println("", birString, bulundu)

}
