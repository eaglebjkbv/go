package main

import "fmt"

type kisi struct {
	adi    string
	soyadi string
}

func veriTipiGoster(a interface{}) {

	switch v := a.(type) {
	case string:
		fmt.Println("Bu string tipinde bri veridir...", v)
	case int:
		fmt.Println("Bu integer tipinde bir veridir.", v)
	case kisi:
		fmt.Println("Bu kisi tipinde bir veridir", v.adi)
	default:
		fmt.Println("Bu verinin tipi anlaşılamadı...")
	}

}
func main() {
	veriTipiGoster(1.2)
}
