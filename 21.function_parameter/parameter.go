//saat membuat function memerlukan data dari luar function atau parameter
//parameter bisa lebih dari satu
//jika menambahkan parameter,wajib memasukkan parameter

package main

import "fmt"

func main() {
	firstName := "Alwan"
	for i := 0; i < 5; i++ {
		sayHello(firstName, "Hafidzin")
	}
}

func sayHello(firstName string, lastName string) {
	fmt.Println("Hallo", firstName, lastName)
}
