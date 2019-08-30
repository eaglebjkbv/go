package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(2)
		go selamVer()
		go hoscakalDe()
	}
	wg.Wait()
}

func selamVer() {
	fmt.Println("Selamlar")
	wg.Done()
}

func hoscakalDe() {
	fmt.Println("Hoşçakal")
	wg.Done()
}
