package model
import "fmt"

type Person struct { // exported
    First   string // exported
    last    string // unexported
    age     int    // unexported
}

func (p *Person) AdiEkranaYaz(){ // exported
  fmt.Println("Kişinin Adı :" ,p.First)
  p.First="Riky"
}
