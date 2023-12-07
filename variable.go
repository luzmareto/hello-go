package main

import (
	"fmt"
	"strings"
)

func main() {
	// instan value
	fmt.Println("hello world")

	fmt.Println(strings.Repeat("=", 20), "STRING", strings.Repeat("=", 20))

	// long declaration string
	var name string = "supermandi"
	fmt.Println(name)

	// zero declaration string
	var name1 string
	name1 = "Onno"
	fmt.Println(name1)

	// declaration var string
	var name2 = "Widodo"
	fmt.Println(name2)

	// short declaration string
	name3 := "Purbo"
	fmt.Println(name3)

	// print all
	fmt.Println(name1, name2, name3)

	// mengganti value
	gantiNama := "ahmad nathan" //value pertama
	gantiNama = "luzmareto"     // mengubah value gantiNama
	fmt.Println(gantiNama)

	//variable const(value tidak bisa diubah)
	const varConst = "asus"
	fmt.Println(varConst)

	var kampus, alamat string = "ITTS", "TANGSEL"
	fmt.Println(kampus, alamat)

	fmt.Println(strings.Repeat("=", 20), "INT", strings.Repeat("=", 20))

	var angka int = 1
	fmt.Println(angka)

	// penjumlahan
	angka1 := 10
	angka2 := 10
	fmt.Println(angka1 + angka2)
}
