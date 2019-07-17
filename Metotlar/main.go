package main

import (
	"fmt"
)

type Hayvan struct {
	name string
}

func main() {
	denem := Hayvan{"Fino"}
	fmt.Println(denem.kopek())

}
func (h Hayvan) kopek() string {

	return h.name + " hav hav der"
}
