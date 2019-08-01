package main

import "fmt"

type alanhesap interface {
	alan() (int, string)
}

type dikdortgen struct {
	aKenari int
	bKenari int
}
type kare struct {
	kenar int
}

func (d dikdortgen) alan() (int, string) {
	return d.aKenari * d.bKenari, "Dikdorgen"
}

func (k kare) alan() (int, string) {
	return k.kenar * k.kenar, "Kare"
}

func alanHsp(a alanhesap) {
	fmt.Println(a.alan())
}

func alanBul(a ...alanhesap) {
	for _, b := range a {
		al, tip := b.alan()
		fmt.Println(tip, " alanı =", al)
	}

}

func main() {

	alanHsp(dikdortgen{10, 20})
	alanHsp(kare{10})

	alanBul(kare{30}, dikdortgen{3, 4})
}
