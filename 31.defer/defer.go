//Defer : function yang dijadwalkan ketika sebuah function selesai dieksekusi
//Defer function akan selalu dieksekusi walapun terjadi error di function yang dieksekusi

package main

import "fmt"

func main() {
	//Tanpa defer
	// runApplication()
	//coba 0 maka error
	runApplication(0)
}

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	//deffer di atas
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result", result)
	// logging()
}
