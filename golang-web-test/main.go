package main

import (
	"golang-web-test/example"
	"net/http"
)

//birim testlerinde bu kısma bakılmayacak
//web sunucu çalıştırılmasada sıkıntı olmaz
//TCP 3000 PORTDAN DİNLEYECEK
func main() {
	http.HandleFunc("/", example.Selamla)
	http.ListenAndServe(":3000", nil)
}
