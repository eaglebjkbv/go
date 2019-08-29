package model

// Kisi yapsısı kişisel bilgilerin kaydedilebileceği bir yapıdır...
type Kisi struct { //ulaşılabiliir.
	Adi      string
	Soyadi   string
	Yasi     int
	Hobileri []string
}

// TamAdi Kişinin adını ve soyadını birleştirir...
func (k Kisi) TamAdi() string {
	return k.Adi + " " + k.Soyadi
}
