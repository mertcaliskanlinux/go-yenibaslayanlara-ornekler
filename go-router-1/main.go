package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", Anasayfa)
	r.GET("/gizli", GizemliSeylerYap)
	log.Println("Gizli Giriş Tıklandı")
	http.ListenAndServe(":8000", r)

}

func Anasayfa(HTMLcevap http.ResponseWriter, HTMListek *http.Request, HTMLRouter httprouter.Params) {
	fmt.Fprint(HTMLcevap,
		`<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Document</title>
		</head>
		<body>
			<h2>Anasayfa</h2>
			<a href="/gizli" src="">Gizli Giriş...</a>
		</body>
		</html>`,
	)

}
func GizemliSeylerYap(HTMLcevap http.ResponseWriter, HTMListek *http.Request, router httprouter.Params) {
	fmt.Fprint(HTMLcevap, `Kimseye Söyleme`)
}
