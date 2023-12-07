package main

import "fmt"

func main() {
	number := 10
	if number%2 == 0 {
		fmt.Println("nilai genap")
	} else {
		fmt.Println("gatau nilainya")
	}

	// short condition
	if waktu := 9; waktu >= 6 && waktu <= 10 {
		fmt.Println("sarapan dulu")
	} else if waktu > 10 && waktu < 14 {
		fmt.Println("sana makan siang dulu")
	} else {
		fmt.Println("makan sore atau malam")
	}

	//yang paling sering dipake
	admin := 1
	if admin == 1 {
		fmt.Println("kamu admin")
	}

}
