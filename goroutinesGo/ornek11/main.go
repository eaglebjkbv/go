package main

import (
	"fmt"
	"time"
)

func main() {
	go selamVer()
	time.Sleep(100 * time.Millisecond)
}
func selamVer() {
	fmt.Println("Merhaba")
}
