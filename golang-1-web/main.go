package main

import (
	"html/template"
	"net/http"
)

type Söz struct {
	Söyleyen string
	Zaman    int
	Cümle    string
}

func Sözlistele(HTMLcevap http.ResponseWriter, HTMListek *http.Request) {
	var ÖzlüSözler [5]Söz
	ÖzlüSözler[0] = Söz{"Thomas Watson (IBM)", 1943, "Dünyada Belkide En Fazla 5 Bilgisayara İhtiyaç Vardır."}
	ÖzlüSözler[1] = Söz{"Ken Olsan (DEC)", 1977, `Bence Kimse Evine Bilgisayar Sokmak İçin Geçerli Bir Nedeni Olamaz.`}
	ÖzlüSözler[2] = Söz{"Darry Zaylnux (Film Yapımcısı)", 1993, `Televizyonun Modası Geçicek`}
	ÖzlüSözler[3] = Söz{"Murat Özalp (this-Author)", 1993, `Tarayıcıdan Bilgisayara Aktırlan Yazılan Asla İşlenemeyecektir.`}

	şablon, _ := template.ParseFiles("templates/index.html")
	şablon.Execute(HTMLcevap, ÖzlüSözler)
}
func main() {
	http.HandleFunc("/", Sözlistele)
	http.ListenAndServe(":8000", nil)
}
