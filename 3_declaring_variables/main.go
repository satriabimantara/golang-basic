package main

import "fmt"

func main() {
	// deklarasi variabel dengan manifest typing
	var firstName string = "John"
	var lastName string
	lastName = "Wick"
	fmt.Printf("Halo %s %s!\n", firstName, lastName)

	// deklarasi variable tanpa tipe data (type inference)
	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	middleName := "Ethan"
	fmt.Printf(middleName + "\n")

	// deklarasi multivariable
	var first, second, third string
	first, second, third = "John", "Ethan", "Wick"
	// lebih ringkas
	// fourth, fifth, sixth := "4", "5", "6"
	fmt.Printf("%s %s %s\n", first, second, third)

	// variable (_) untuk menampung nilai yang tidak digunakan. Variable ini tidak akan ditampilkan dalam running program
	_ = "Studi"

	// deklarasi variabel menggunakan keyword new untuk mmebuat variable pointer
	name := new(string)
	fmt.Println(name)  // print alamat memori
	fmt.Println(*name) //print isi variable setelah deference

}
