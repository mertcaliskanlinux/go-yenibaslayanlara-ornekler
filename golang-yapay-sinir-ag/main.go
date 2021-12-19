package main

import (
	"fmt"
	"math/rand"

	"github.com/goml/gobrain"
)

func main() {
	//RANDOM SAYILAR ÜRETİNİCİ BESLE
	rand.Seed(0)

	//2 GİRİŞ 1 ÇIKIŞLI XOR sistemi
	Desenler := [][][]float64{
		{{0, 0}, {0}},
		{{0, 1}, {1}},
		{{1, 0}, {1}},
		{{1, 1}, {0}},
	}
	//Yapay Sinir Ağ Oluşturulsun
	//2 giriş,2 gizli katman , 1 çıkış
	YSA := &gobrain.FeedForward{}
	YSA.Init(2, 2, 1)
	fmt.Printf("%+v\n", *YSA)

	//XOR DESENLERİ KULLANARAK AĞ EĞİT 4000 EPOCH ÇALIŞSIN.
	//ÖĞRENME KAT SAYISI 0.6 VE MOMENTUM 0.4 OLSUN.
	YSA.Train(Desenler, 4000, 0.6, 0.4, true)
	fmt.Printf("%+v\n", *YSA)

	fmt.Println("Sistemin Cevapları: ")
	YSA.Test(Desenler)
}
