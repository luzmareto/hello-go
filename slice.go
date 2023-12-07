package main

import (
	"fmt"
	"strings"
)

func main() {
	//menghitung slice menggunakan index dimana angka pertama adalah 0

	//value slice bulan
	var bulan = [...]string{
		"januari",   //0
		"februari",  //1
		"maret",     //2
		"april",     //3
		"mei",       //4
		"juni",      //5
		"juli",      //6
		"agustus",   //7
		"september", //8
		"oktober",   //9
		"november",  //10
		"desember",  //11
	}

	//variable untuk memproses slice bulan
	var slice1 = bulan[4:7]
	fmt.Println(cap(slice1)) //menghitung 8 data
	fmt.Println(len(slice1)) //ada 3 data yang diinput
	fmt.Println("value slice ori adalah : ")
	fmt.Println(slice1) //[mei juni juli]

	fmt.Println(strings.Repeat("=", 80))

	//mengubah value bulan
	fmt.Println("value bulan yang telah diubah adalah")
	bulan[5] = "loncat"
	fmt.Println(slice1)
}
