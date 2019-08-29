package main

import "fmt"

type sekilHasaplayici interface {
	alan() float64
	cevre() float64
}

type dikdortgen struct {
	e, b float64
}

func (d dikdortgen) alan() float64 {
	return d.e * d.b

}
func (d dikdortgen) cevre() float64 {
	return (d.e + d.b) * 2
}

const pi = 3.14

type cember struct {
	r float64
}

func (c cember) alan() float64 {
	return pi * c.r * c.r
}
func (c cember) cevre() float64 {
	return 2 * pi * c.r
}

func main() {

	var d sekilHasaplayici = dikdortgen{10, 20}
	fmt.Println("Dikdörtgenin alanı : ", d.alan())

	// var d [1]sekilHasaplayici
	// d[0] = dikdortgen{10, 5}
	// fmt.Println("Alan : ", d[0].alan())

}
