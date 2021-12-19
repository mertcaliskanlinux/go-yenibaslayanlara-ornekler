package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

//SORGU SONUCUNU AKTARACAĞIMIZ YAPI
type Kitap struct {
	isim  string
	yazar string
}

//vt değişkenini global olarak tanımladığımız için argüman göndermeye gerek yok.
var VT *sql.DB

func main() {

	VT, _ = sql.Open("sqlite3", "kitaplar.sqlite")
	http.HandleFunc("/", Listele)

	http.ListenAndServe(":8000", nil)

}

func Listele(HTMLcevap http.ResponseWriter, HTMListek *http.Request) {
	//SORGUYU ÇALIŞTIR.
	Satirlar, _ := VT.Query("SELECT * FROM kitap")
	defer Satirlar.Close()
	//YENİ BİR KESİT OLUŞTUR
	kitaplar := make([]*Kitap, 0)
	//SQL CEVABINI HER SATIR İÇİN DÖN.
	for Satirlar.Next() {
		satir := new(Kitap)
		//SATIR YAPISININ ALANLARINI DOLDUR
		Satirlar.Scan(&satir.isim, &satir.yazar)
		//YENİ SATIR ALANLARINI KİTAPLAR KESİTİNE EKLE
		kitaplar = append(kitaplar, satir)
	}
	//HTML CEVABINI OLUŞTUR.
	for i, kitap := range kitaplar {
		fmt.Fprintf(HTMLcevap, "%d),İsim: %s-Yazar:%s\n", i, kitap.isim, kitap.yazar)
	}
}
