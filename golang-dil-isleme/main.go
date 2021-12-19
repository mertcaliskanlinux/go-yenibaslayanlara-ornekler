package main

import (
	"fmt"

	DilTahmin "github.com/endeveit/guesslanguage"
)

func main() {
	fmt.Println(DilTahmin.Guess("Benim Sadık Yarim Kara Topraktır."))
	fmt.Println(DilTahmin.Guess("Love All,Trust a few,do wrong to none"))
	fmt.Println(DilTahmin.Guess("Si quieres hacer reir a dios."))
	fmt.Println(DilTahmin.Guess("Steter Tropfen höhlt den Stein."))
	fmt.Println(DilTahmin.Guess("Merhaba Dünya MERT ÇALIŞKAN"))

}
