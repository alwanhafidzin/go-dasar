// array kumpulan data dengan tipe sama
//dalam array perlu menentukan jumlah data yang bisa ditampung
//daya tampung tidak bisa ditambah setelah Array dibuat
//Indeks dimulai dari 0
package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Alwan"
	names[1] = "Hafidzin"
	names[2] = "firdaus"

	fmt.Println(names[0])
	//error index of bones
	// fmt.Println(names[3])

	//Membuat array secara langsung,ingat selalu diakhiri koma
	var values = [3]int{
		98,
		80,
		95,
	}
	fmt.Println(values)
	fmt.Println(values[1])

	// function Array
	// len(Array) -> Untuk mendpatkan panjang Array
	// array[index] -> mendapatkan data di posisi index
	// array[index] = value -> Mengubah data di posisi index

	fmt.Println(len(values))
	values[1] = 100
	fmt.Println(values[1])

	//cek len tanpa data akan mengembalikan sesuai data yang ditentukan
	var test [10]string
	fmt.Println(len(test))
}
