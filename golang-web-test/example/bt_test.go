package example

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Selamla(t *testing.T) {
	//HERHANGİ BİR REQUEST OLUŞTURALIM
	//AŞAĞIDA fonksiyonu çağırırken boş çağırAmayacagımız için lazım olacak.
	HTMListek, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	//TEST EDİLECEK FONKSİYONU ÇAĞIR:
	HTMLCevap := httptest.NewRecorder()
	Selamla(HTMLCevap, HTMListek)
	//KARŞILAŞTIRMAYI YAP
	Beklenen := "Merhaba Gopher"
	Gelen := HTMLCevap.Body.String()
	if Beklenen != Gelen {
		t.Fatalf("\n\n***Selamla Başarısız:\n '%s' Beklerken '%s' Geldi...", Beklenen, Gelen)
	}
}
