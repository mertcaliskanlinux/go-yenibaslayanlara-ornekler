package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Println("TCP 8000 PORTUNU DİNLİYORUM")
	http.ListenAndServe(":8000", nil)
}
