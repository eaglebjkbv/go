package main

import "fmt"

type Kisi struct {
	Adi    string
	Soyadi string
	yasi   int
}

func main() {

	ps := new(Kisi)

	ps.Adi = "Yamamoto"
	ps.Soyadi = "Sorimachi"
	fmt.Println(ps)
	fmt.Println(ps.Adi)

}
