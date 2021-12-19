package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Sözler struct {
	Söyleyen string
	Zaman    int
	Cümle    string
}

func main() {
	//websayfasını indir ve işle

	url := "http://127.0.0.1:8000/"
	http_cevap, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s BAGLANAMADI", url)
	}

	json_cevap, _ := ioutil.ReadAll(http_cevap.Body)

	defer http_cevap.Body.Close()

	//JSON ELEMANLARINI TUTACAK OLAN YAPI
	var sözler []Sözler
	//json ayrıştır ve sözler yapsınının değişkenine aktar
	json.Unmarshal(json_cevap, &sözler)
	//sözleri eleman sayısını yazdır
	fmt.Printf("JSONDAN Gelen Eleman Sayısı %d\n", len(sözler))
	//sözler dizisini ekrana düzenleyerek yazdır
	for i, s := range sözler {
		fmt.Printf("%d) %d sene önce   %s demişki: %s\n", i, time.Now().Year()-s.Zaman, s.Söyleyen, s.Cümle)
	}

}
