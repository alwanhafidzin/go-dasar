//biasanya memberi tahu bahwa function mengembalikan value maka mendeklarasikan tipe data return value di function
//di golang kita bisa membuat variabel secara langsung di tipe data return function nya
//fitur ini hanya di golang
package main

import "fmt"

func main() {
	a, b, c := getFullName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "Alwan"
	middleName = "Hafidzin"
	lastName = "Firdaus"
	//gak perlu seperi di bawah cukup return saja
	// return firstName, middleName, lastName
	return
}
