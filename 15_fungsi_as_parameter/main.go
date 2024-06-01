package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"John", "Ethan", "Wick", "Nataniel"}
	var callback1 = func(datum string) bool {
		return strings.Contains(datum, "o")
	}
	var filterStringContainO = filter(data, callback1)

	fmt.Println("data asli \t\t:", data)
	// data asli : [wick jason ethan]

	fmt.Println("filter ada huruf \"o\"\t:", filterStringContainO)
	// filter ada huruf "o" : [jason]

	var callback2 = func(datum string) bool {
		return len(datum) == 5
	}
	var dataLength5 = filter(data, callback2)
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLength5)
	// filter jumlah huruf "5" : [jason ethan]

	// Alias Closure
	dataLength5 = filter2(data, callback2)
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLength5)
}

func filter(data []string, callback func(string) bool) []string {
	var results []string

	for _, datum := range data {
		filtered := callback(datum)
		if filtered {
			results = append(results, datum)
		}
	}

	return results
}

// alias skema closure
// Untuk fungsi yang skema-nya cukup panjang, akan lebih baik jika menggunakan alias dalam pendefinisiannya, apalagi ketika ada parameter fungsi lain yang juga menggunakan skema yang sama,

type FilterArray func(string) bool

func filter2(data []string, callback FilterArray) []string {
	var results []string

	for _, datum := range data {
		filtered := callback(datum)
		if filtered {
			results = append(results, datum)
		}
	}

	return results

}
