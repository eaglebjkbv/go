package model

// Kisi  yapısı kişisel verilerin turulduğu yapıdır.
type Kisi struct {
	Adi      string
	Soyadi   string
	Yasi     int
	Hobileri []string
}

// TamAdi : Bu Fonksiyon Kişinin Adı ve Soyadını birleştirir.
func (k Kisi) TamAdi() string {
	return k.Adi + " " + k.Soyadi
}
