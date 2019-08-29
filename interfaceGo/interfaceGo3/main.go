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

func (m Mysql) Kaydet() {
	fmt.Println(m.kisi.adi, " ya ait veriler Mysql e kaydedildi.")
}

type Oracle struct {
	kisi Kisi
}

func (o Oracle) Kaydet() {
	fmt.Println(o.kisi.adi, " ya ait veriler Oracle e kaydedildi.")
}

type Kaydedici interface {
	Kaydet()
}

func vtKaydet(k Kaydedici) {
	k.Kaydet()
}

func main() {
	k := Kisi{
		id:     1,
		adi:    "Yoshi",
		soyadi: "Takamatsu",
	}

	m := Oracle{k}
	m2 := Mysql{k}
	m.Kaydet()
	vtKaydet(m)
	vtKaydet(m2)

}
