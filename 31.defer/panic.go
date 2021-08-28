//panic function adalah function yang bisa kita gunakan untuk menghentikan program
//panic function biasanya dipanggil ketika terjadi error saat program berjalan
//saat panic function dipanggil,program akan terhenti,namun defer function akan tetap dieksekusi

package main

import "fmt"

func main() {
	runApp(true)
}

func endApp() {
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}
