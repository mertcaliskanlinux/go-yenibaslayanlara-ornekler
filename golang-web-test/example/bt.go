package example

import (
	"fmt"
	"net/http"
)

//birim Testi uygulanack olan fonksiyon
func Selamla(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Merhaba Gopher")
}
