package model

// Kisi structu içerisinde kişisel verilerin bulunduğu bir yapıdır.
type Kisi struct {
	Adi      string
	Soyadi   string
	Yasi     int
	Hobileri []string
}

// TamAdi Kişinin adı ve soyadı birleştirilir...
func (k Kisi) TamAdi() string {
	return k.Adi + " " + k.Soyadi
}
