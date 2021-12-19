package main

import (
	"io"
	"net/http"
)

func Anasayfa(cevap http.ResponseWriter, istek *http.Request) {
	cevap.Header().Set("Content-type", "text/html")
	io.WriteString(cevap, `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
			<h1>hello</h1>
	</body>
	</html>`)
}

func ikinciSayfa(cevap http.ResponseWriter, istek *http.Request) {
	cevap.Header().Set("Content-Type", "text/html")
	io.WriteString(cevap, "Go Programlama Dili WEBMASTER")
}
func main() {
	http.HandleFunc("/", Anasayfa)
	http.HandleFunc("/ikincisayfa/", ikinciSayfa)
	http.ListenAndServe(":1983", nil)

}
