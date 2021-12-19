package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	fmt.Println("Giriş")

	r.GET("/selamla/:degisken", Anasayfa)
	http.ListenAndServe(":8000", r)
}

func Anasayfa(HTMLcevap http.ResponseWriter, HTMListek *http.Request, Arguman httprouter.Params) {
	isim := Arguman.ByName("degisken")
	fmt.Fprintf(HTMLcevap, "Welcom : %s!\n", isim)
}
