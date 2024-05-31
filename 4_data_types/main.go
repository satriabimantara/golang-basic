package main

import "fmt"

func main() {
	// 1. tipe data numerik non desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// 2. tipe data numerik desimal
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// 3. Tipe data boolean
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	// 4. Tipe data string
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	message = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang".`

	fmt.Println(message)
}
