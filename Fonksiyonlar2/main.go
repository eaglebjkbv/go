package main

import "fmt"

func deneme() {
	fmt.Println("Deneme Fonksyonu çağırıldı")
}
func bar(x string) string {
	return fmt.Sprintln(x, " Burada")
}

func main() {
	fmt.Println(fmt.Sprintln("Merhaba Dünya"))
	
	a := bar("Bülent")
	fmt.Print(a)

	deneme()

}
