package main

import "fmt"

/*
Closure->suatu anonymous function (atau fungsi tanpa nama) yang disimpan dalam variabel. Closure biasa dimanfaatkan untuk membungkus suatu proses yang hanya dijalankan sekali saja atau hanya dipakai pada blok tertentu saja.
*/

func main() {
	// 1. closure disimpan dalam variabel
	var getMinMax = func(numbers ...int) (int, int) {
		var min, max int
		for i, e := range numbers {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3, 15}
	var min, max = getMinMax(numbers...)
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)

	// 2. IIFE
	iife()

	// 3. closure sebagai nilai kembalian
	max = 3
	numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, getNumbers = findMax(numbers, max)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t:", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany)    // 9
	fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]

}

func iife() {
	// Immediately-Invoked Function Expression (IIFE)
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	fmt.Println("original number :", numbers)
	fmt.Println("filtered number :", newNumbers)
}

// closure sebagai nilai kembalian
// closure bisa dijadikan sebagai nilai balik dari suatu fungsi

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	// tampung closure untuk dijadikan sebagai return value
	var getNumbers = func() []int {
		return res
	}
	return len(res), getNumbers
}
