package main

import "fmt"

func main() {
	// gak bisa var data int = ups(1)
	//interface kosong harus menggunakan tipe interface kosong
	var data interface{} = Ups(3)
	fmt.Println(data)
}

// inteface{} kosong ,tanpa function,semua tipe data di golang mengikuti kontrak interface kosong
//interface kosong gak peduli parameter apapun
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}

}
