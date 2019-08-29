package main


import "github.com/eaglebjkbv/go_Examples/exportedStructs/model"
import "fmt"

func main() {

var p model.Person
fmt.Println(p.First)
p.First="Pike"
p.AdiEkranaYaz()
p.AdiEkranaYaz()

}
