package main

import "fmt"

func main() {
	// 1. contoh deklarasi array dengan tipe data string
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	// 2. inisialisasi nilai awal pada array gaya horisontal
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// gaya vertical
	/*
		Khusus untuk deklarasi array dengan cara vertikal, tanda koma wajib dituliskan setelah setiap elemen (termasuk elemen terakhir), agar tidak memunculkan syntax error.
	*/
	fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}
	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	// 3. Inisialisasi nilai awal array tanpa jumlah elemen
	var numbers = [...]int{2, 3, 2, 4, 3}
	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	// 4. Multidimensional Array
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	// 5. Perulangan Elemen Array
	// a. Keyword for
	fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}
	// b. Keyword for-range
	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}
	// c. Variable _ (mengabaikan index loop) pada keyword for-range
	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	// 6. Alokasi elemen array dengan keyword make
	var buahs = make([]string, 2)
	buahs[0] = "apple"
	buahs[1] = "manggo"

	fmt.Println(buahs) // [apple manggo]

}
