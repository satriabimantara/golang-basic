package main

import "fmt"

func main() {
	// deklarasi constanta dapat menggunakan cara manifest typing atau type inference
	const firstName string = "john"
	fmt.Print("halo ", firstName, "!\n")

	const lastName = "wick"
	fmt.Print("nice to meet you ", lastName, "!\n")

	// deklarasi multi constanta
	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.2
	)
	/*
		Contoh deklarasi konstanta dengan tipe data dan nilai yang sama:
		Ketika tipe data dan nilai tidak dituliskan dalam deklarasi konstanta, maka tipe data dan nilai yang dipergunakan adalah sama seperti konstanta yang dideklarasikan diatasnya.
		-a dideklarasikan dengan metode type inference dengan tipe data string dan nilainya "konstanta"
		-b dideklarasikan dengan metode type inference dengan tipe data string dan nilainya "konstanta"
	*/
	const (
		a = "konstanta"
		b
	)

	fmt.Print(isToday)
}
