package main

import (
	"fmt"
	"strings"
)

/*
Go mengadopsi konsep variadic function atau pembuatan fungsi dengan parameter bisa menampung nilai sejenis yang tidak terbatas jumlahnya.

Parameter variadic memiliki sifat yang mirip dengan slice, yaitu nilai dari parameter-parameter yang disisipkan bertipe data sama, dan kesemuanya cukup ditampung oleh satu variabel saja. Cara pengaksesan tiap nilai juga mirip, yaitu dengan menggunakan index.
*/
func main() {
	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)

	// pengisian parameter pada fungsi variadic dengan slice
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	avg = calculate(numbers...)
	msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)

	// pemanggilan fungsi dengan parameter kombinasi
	yourHobbies("wick", "sleeping", "eating", "reading")
	var hobbies = []string{"sleeping", "eating"}
	yourHobbies("John Wick", hobbies...)

}

// 1. contoh penerapan fungsi variadic
func calculate(numbers ...int) (avg float64) {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	avg = float64(total) / float64(len(numbers))
	return
}

// 2. Kombinasi parameter biasa dan variadic
// parameter variadic harus diletakkan di akhir
func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}
