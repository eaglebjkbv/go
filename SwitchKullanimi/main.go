package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Cumartesi Ne zaman")
	bugun := time.Now().Weekday()
	switch time.Saturday {
	case bugun + 0:
		fmt.Println("Bugün Cumartesi")
	case bugun + 1:
		fmt.Println("Yarın Cumartesi")
	default:
		fmt.Println("Cumartesiye çok var")
	}

}
