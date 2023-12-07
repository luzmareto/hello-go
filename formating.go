package main

import "fmt"

func main() {
	// simbol \n untuk mencetak enter

	// string menggunakan %s
	firstName := "adipati"
	lastName := "dolken"
	fmt.Printf("halo %s %s!\n", firstName, lastName)

	// angka pecahan format float menggunakan %f
	var decimalNumber = 2.620000
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)

	// bulangan bulat menggunakan %d
	var bilanganBulat = 262
	fmt.Printf("bilangan desimal: %d\n", bilanganBulat)
}
