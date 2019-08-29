package main

import "fmt"

type Kisi struct {
	id     int
	adi    string
	soyadi string
}

type Mysql struct {
	kisi Kisi
}

func (m Mysql) Kaydet() bool {
	fmt.Println(m.kisi.adi, " ye ait veriler Mysql Veritabanına kaydedildi.")
	return true
}

type Oracle struct {
	kisi Kisi
}

func (o Oracle) Kaydet() bool {
	fmt.Println(o.kisi.adi, " ye ait veriler Oracle Veritabanına kaydedildi.")
	return true
}

type Kaydedici interface {
	Kaydet() bool
}

func Kaydet(k Kaydedici) bool {
	return k.Kaydet()
}

func main() {
	m := Mysql{
		kisi: Kisi{
			id:     1,
			adi:    "Yoshi",
			soyadi: "Takamatsu",
		},
	}
	m2 := Oracle{
		kisi: Kisi{
			id:     2,
			adi:    "Ryu",
			soyadi: "Takamatsu",
		},
	}

	if Kaydet(m) {
		fmt.Println("Kayıt Başarılı")
	}
	if Kaydet(m2) {
		fmt.Println("Kayıt Başarılı")
	}

}
