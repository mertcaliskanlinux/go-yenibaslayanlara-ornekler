package main

import (
	"fmt"

	"github.com/go-opencv/go-opencv"
	"github.com/go-opencv/go-opencv/opencv"
)

func main() {
	//Resim Dosyasını yükle

	ResimDosyası := "Gopher.png"
	resim := opencv.LoadImage(ResimDosyası)

	defer resim.Release()

	//Pencereyi ayarla

	pencere := opencv.NewWindow("Go-OpenCV Uygulaması")
	defer pencere.Destroy()

	//Ayar Çubuğu Yerleştir ve Değerini alıp işle
	pencere.CreateTrackbar("Ayar Çubuğu", 1, 100, func(deger int) {
		fmt.Printf("Değer = %d", deger)
	})
	pencere.ShowImage(resim)
	//Bir tuşa Basıncaya kadar Bekle
	opencv.WaitKey(0)

}
