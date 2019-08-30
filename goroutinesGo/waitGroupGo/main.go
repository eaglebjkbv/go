package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} // Waitgroup bir tür sayaçtır.

func main() {
	mesaj := "Selamlar"
	wg.Add(1) // sayaca bir goroutine ekliyoruz.
	go func(msj string) {
		fmt.Println(msj)
		wg.Done() // sayactan eksiltiyoruz.
	}(mesaj)

	wg.Wait()

}
