//Recover adalah function yang bisa kita gunakan untuk menangkap data panic
//Dengan recover proses panic akan terhenti,sehingga program akan tetap berjalan

package main

import "fmt"

func main() {
	runApp(true)
	//ketika manggil run aplication dan panic akan berhenti semuanya
	fmt.Println("Alwan")
}

func endApp() {
	//dipindah ke defer akhir biar bisa,jangan di function satunya
	message := recover()
	if message == nil {
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}
