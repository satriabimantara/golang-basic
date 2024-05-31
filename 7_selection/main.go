package main

import "fmt"

func main() {
	// 1. seleksi kondisi dengan if, else if, dan else
	var point = 8

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	// 2. variable temporary pada if-else
	/*
		- Variabel temporary adalah variabel yang hanya bisa digunakan pada deretan blok seleksi kondisi di mana ia ditempatkan.
		- Deklarasi variabel temporary hanya bisa dilakukan lewat metode type inference yang menggunakan tanda :=
	*/
	point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// 3. switch case
	/*
		Di Go, ketika sebuah case terpenuhi, tidak akan dilanjutkan ke pengecekan case selanjutnya, meskipun tidak ada keyword break di situ.
	*/
	point = 6

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// 4. switch case dengan banyak kondisi
	point = 6

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// 5. kurung kurawal pada keyword case dan default
	point = 6

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	// 6. switch case dengan gaya if-else
	point = 6

	switch {
	case point == 8:
		fmt.Println("perfect")
	case point < 8 && point > 3:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// 7. keyword fallthrough dalam switch
	/*
		- Keyword fallthrough digunakan untuk memaksa proses pengecekan tetap diteruskan ke case selanjutnya dengan tanpa menghiraukan nilai kondisinya,
		- Pada case dalam sebuah switch, diperbolehkan terdapat lebih dari satu fallthrough.
	*/
	point = 6

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// 8. Nested if
	point = 10

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}

}
