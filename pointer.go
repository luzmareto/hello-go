package main

import "fmt"

func main() {
	var pertama int = 12
	var keuda *int = &pertama //var kedua mengambil alamat pertama

	fmt.Println("pertama (value)   :", pertama)  // 4
	fmt.Println("pertama (address) :", &pertama) // 0xc20800a220

	fmt.Println("keuda (value)   :", *keuda) // 4
	fmt.Println("keuda (address) :", keuda)  // 0xc20800a220
}
