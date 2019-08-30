package main

import (
	"fmt"
	"time"
)

func main() {
	mesaj := "Selamlar"

	go func(msj string) {
		fmt.Println(msj)
	}(mesaj)

	mesaj = "Hoşçakal"
	time.Sleep(100 * time.Millisecond)
}
