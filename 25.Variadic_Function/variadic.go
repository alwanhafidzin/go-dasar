// Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
// Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam Array.

// Apa bedanya dengan parameter biasa dengan tipe data Array?
// Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
// JIka parameter menggunakan varargs, kita bisa langsung mengirim data nya, jika lebih dari satu, cukup gunakan tanda koma

package main

import "fmt"

//variadic argumen hanya bisa dibelakang tidak bisa di awal atau tengah dll
func main() {
	total := getSumAll("Alwan", 10, 90, 30, 50)
	fmt.Println(total)

	//memasukkan slice ke variadic function,dengan ...

	slice := []int{10, 20, 30, 40}
	total = getSumAll2(slice...)
	fmt.Println(total)
}

//[]int slice
func getSumAll(name string, numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}
	return total

}
func getSumAll2(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}
	return total

}
