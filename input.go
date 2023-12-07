package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// //membuat tampilan di terminal
	// fmt.Printf("masukan nama produk : ")

	// //membuat scan yang berfungsi menerima input nilai dari terminal
	// scan := bufio.NewScanner(os.Stdin)
	// scan.Scan()

	// //variable penampung yang berfungsi menyimpan nilai dari terminal
	// produk := scan.Text()

	// //menampilkan nilai yang diinput dari terminal
	// fmt.Println("nama kamu adalah : " + produk)

	fmt.Println(strings.Repeat("=", 20), "simulasi biodata", strings.Repeat("=", 20))

	// simulasi input biodata
	fmt.Print("Masukkan nama kamu: ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	nama := scan.Text()

	fmt.Print("Masukkan umur kamu: ")
	scan.Scan()
	umur := scan.Text()

	fmt.Print("Masukkan alamat kamu: ")
	scan.Scan()
	alamat := scan.Text()

	fmt.Println()
	fmt.Println(strings.Repeat("=", 20), "biodata kamu adalah ", strings.Repeat("=", 20))
	fmt.Println()

	fmt.Println("Nama kamu adalah:", nama)
	fmt.Println("Umur kamu adalah:", umur)
	fmt.Println("Alamat kamu adalah:", alamat)
}
