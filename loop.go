package main

import (
	"fmt"
	"strings"
)

func main() {
	// long loop
	i := 10       //var i nilainya 10
	for i <= 25 { //mencetak nilai dari sepuluh sampai 20
		fmt.Println(i)
		i = i + 5 //mencetak dengan menambah angka 5
	} //nilai berenti di angka 20

	fmt.Println(strings.Repeat("=", 40))

	// short loop
	for angka := 1; angka < 3; angka++ {
		fmt.Println(angka)
	} //angka 3 tidak dicetak
}
