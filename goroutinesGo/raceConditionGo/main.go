package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var mt = sync.Mutex{}

func main() {
	say := 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mt.Lock() // diğer goroutine lerin yazma yapmasını engeller
			say++
			mt.Unlock() // Engeli Kaldırır
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("Sayı ifadesi : ", say)

}
