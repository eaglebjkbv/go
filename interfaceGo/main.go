package main

import "fmt"


type alanhesap interface{
	alan() int
}

type dikdortgen struct{
	aKenari int
	bKenari int
}
type kare struct{
	kenar int
}

func (d dikdortgen) alan() int{
	return d.aKenari*d.bKenari
}

func (k kare) alan() int{
	return k.kenar*k.kenar
}

func alanBul(a alanhesap){
	fmt.Print(a.alan())
}

func main(){
	dg:=dikdortgen{10,20}
	kr:=kare{10}

	alanBul(dg)
	alanBul(kr)

}