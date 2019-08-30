package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("Merhaba")
		ch <- 10 // goroutine bu noktada bloklanıyor.
		time.Sleep(time.Millisecond)
		fmt.Println("Deneme")
	}()

	<-ch // groutine bloğu kaldırılıyor.
	//time.Sleep(100 * time.Millisecond)

}
