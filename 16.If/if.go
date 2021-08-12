//if adalah percabangan
//Terdapat if else,else if

package main

import "fmt"

func main() {

	name := "Al"

	if name == "Alwan" {
		fmt.Println("Hallo Alwan")
	} else if name == "Al" {
		fmt.Println("Halo Al")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	//If dengan short statement,titak koma sebagai pemisah short statement

	var length = len(name)
	if length > 5 {
		fmt.Println("terlalu Panjang")
	} else {
		fmt.Println("Nama SUdah benar")
	}

	if length := len(name); length > 5 {
		fmt.Println("terlalu Panjang")
	} else {
		fmt.Println("Nama SUdah benar")
	}

}
