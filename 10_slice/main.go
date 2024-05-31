package main

import "fmt"

func main() {
	// 1. inisialisasi slice
	// Cara pembuatan slice mirip seperti pembuatan array, bedanya tidak perlu mendefinisikan jumlah elemen ketika awal deklarasi
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0]) // "apple"
	// var fruitsA = []string{"apple", "grape"}     // slice
	// var fruitsB = [2]string{"banana", "melon"}   // array
	// var fruitsC = [...]string{"papaya", "grape"} // array

	// 2. Slice indexing (mirip dengan Python)
	var newFruits = fruits[0:2]
	var newFruits2 = fruits[:] // access all element
	fmt.Println(newFruits)     // ["apple", "grape"]
	fmt.Println(newFruits2)    // ["apple", "grape", "banana", "melon"]
	/*
		Slice merupakan tipe data reference atau referensi. Artinya jika ada slice baru yang terbentuk dari slice lama, maka data elemen slice yang baru akan memiliki alamat memori yang sama dengan elemen slice lama. Setiap perubahan yang terjadi di elemen slice baru, akan berdampak juga pada elemen slice lama yang memiliki referensi yang sama.
	*/

	// 3. Build-in function Go untuk operasi slice
	/*
		a. len()
		digunakan untuk menghitung jumlah elemen slice yang ada

		b. cap()
		digunakan untuk menghitung lebar atau kapasitas maksimum slice.

		c. append()
		digunakan untuk menambahkan elemen pada slice

		d. copy()
		digunakan untuk men-copy elements slice pada src (parameter ke-2), ke dst (parameter pertama).

	*/
	// a. len
	fmt.Println(len(fruits)) // 4
	// b. cap()
	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	var aFruits = fruits[0:3]
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4

	var bFruits = fruits[1:4]
	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3

	// c. append()
	var cFruits = append(fruits, "papaya")

	fmt.Println(fruits)  // ["apple", "grape", "banana"]
	fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]

	// d. copy()
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	dst = []string{"potato", "potato", "potato"}
	src = []string{"watermelon", "pinnaple"}
	n = copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple potato
	fmt.Println(src) // watermelon pinnaple
	fmt.Println(n)   // 2

	// 4. Slicing dengan 3 indeks
	/*
		3 index adalah teknik slicing untuk pengaksesan elemen yang sekaligus menentukan kapasitasnya. Cara penggunaannya yaitu dengan menyisipkan angka kapasitas di belakang, seperti fruits[0:1:1]. Angka kapasitas yang diisikan tidak boleh melebihi kapasitas slice yang akan di slicing.
	*/
	fmt.Println("Akses 3 indeks")
	fruits = []string{"apple", "grape", "banana", "jackfruit", "kingfruit", "watermelon"}
	aFruits = fruits[0:2]
	bFruits = fruits[0:2:3]

	fmt.Println(fruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits)) // len: 6
	fmt.Println(cap(fruits)) // cap: 6

	fmt.Println(aFruits)      // ["apple", "grape"]
	fmt.Println(len(aFruits)) // len: 2
	fmt.Println(cap(aFruits)) // cap: 6

	fmt.Println(bFruits)      // ["apple", "grape"]
	fmt.Println(len(bFruits)) // len: 2
	fmt.Println(cap(bFruits)) // cap: 3

}
