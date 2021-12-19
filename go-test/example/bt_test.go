package example

import "testing"

func TestSelamla(t *testing.T) {
	BeklenenCevap := "Merhaba Dünya!"
	GelenCevap := Selam()
	if GelenCevap != BeklenenCevap {
		t.Errorf("\nSelamla testi başarısız!\nBeklenen Cevap:\t'%s'\nGelen Cevap:\t'%s'", BeklenenCevap, GelenCevap)
	}
}

func TestVedalas(t *testing.T) {
	BeklenenCevap := "Görüşmek Üzere Ayşe..."
	GelenCevap := Vedalas("Ayşe")
	if GelenCevap != BeklenenCevap {
		t.Logf("\n Vedalaşma Testi Başarısız %s\nBeklenenCevap GelenCevap:%s\n", BeklenenCevap, GelenCevap)
	}
}

func TestTopla(t *testing.T) {
	BeklenenCevap := 20
	GelenCevap := Topla(10, 10)
	if GelenCevap != BeklenenCevap {
		t.Logf("\n Vedalaşma Testi Başarısız \nBeklenenCevap:%d\n GelenCevap:%d\n", BeklenenCevap, GelenCevap)
	}
}
