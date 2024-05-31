package main

import "fmt"

/*
- Mirip analoginya dengan dictionary pada Python atau array assosiatif pada PHP
- Map adalah tipe data asosiatif yang ada di Go yang berbentuk key-value pair.
- Kalau dilihat, map mirip seperti slice, hanya saja identifier yang digunakan untuk pengaksesan bukanlah index numerik, melainkan bisa dalam tipe data apapun sesuai dengan yang diinginkan.
*/
func main() {
	// 1. contoh deklarasi dan penggunaan map
	// var chicken = map[string]int{} // inisialisasi value map di awal agar tidak nil
	chicken := map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	// 2. Inisialisasi nilai map
	// cara horizontal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// cara vertical
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}
	fmt.Println(chicken1, chicken2)

	// 3. cara lain deklarasi map di golang
	// var chicken3 = map[string]int{}
	// var chicken4 = make(map[string]int)
	// var chicken5 = *new(map[string]int)

	// 4. Iterasi Map dengan For-Range
	chicken = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
		"agustus":  900,
	}

	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}
	// 5. Menghapus key tertentu dengan delete()
	delete(chicken, "januari")

	fmt.Println(len(chicken)) // 1
	fmt.Println(chicken)

	// 6. check items in map if exist with key
	var value_yang_dicari, isExist = chicken["agustus"]

	if isExist {
		fmt.Printf("Value yang dicari: %d\n", value_yang_dicari)
	} else {
		fmt.Println("item is not exists")
	}

	// 7. Kombinasi slice dan map
	var chickenSlices = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}
	for _, chicken := range chickenSlices {
		fmt.Println(chicken["gender"], chicken["name"])
	}
	// Dalam []map[string]string, tiap elemen bisa saja memiliki key yang berbeda-beda, contohnya seperti kode berikut.
	var data = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}
	fmt.Println(data)
	fmt.Println(data[0]["color"])

}
