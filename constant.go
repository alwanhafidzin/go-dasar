//Kata kunci tidak bisa dirubah

package main

import "fmt"

func main (){
	//const bila tidak digunakan tidak akan ada pesan error
	//const tidak ada :=
	const firstName string = "Alwan"
	const lastName = "Hafidzin"

	//Error
	firstName = "Tes"
	lastName = "Yoi"

	//constant multi 
	const {
		firstName = "Alwan"
		lastName = "Hafidzin"
	}
}