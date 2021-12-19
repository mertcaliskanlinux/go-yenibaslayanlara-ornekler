package main

import (
	"fmt"

	"github.com/koyachi/go-nude"
)

func main() {
	resimler := []string{
		"fitness.jpg", "fitness1.jpg", "girl.png", "heykel.jpg", "heykel1.jpg",
	}

	for _, resim := range resimler {
		ciplakmi, _ := nude.IsNude(resim)
		if ciplakmi == true {
			fmt.Printf("Evet Çıplak %s\n", resim)
		} else {
			fmt.Printf("Hayır Çıplak Değil! %s\n", resim)

		}
	}
}
