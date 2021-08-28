/**
Biasanya di dalam bahasa pemrograman lain, object yang belum diinisialisasi maka secara otomatis nilainya adalah null atau nil
Berbeda dengan Go-Lang, di Go-Lang saat kita buat variable dengan tipe data tertentu, maka secara otomatis akan dibuatkan default value nya
Namun di Go-Lang ada data nil, yaitu data kosong
Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice, pointer dan channel
*/

package main

import "fmt"

func main() {
	//nil kosong
	// var person map[string]string = nil
	person := NewMap("Alwan")
	fmt.Println(person)

	//dibawah ini akibat jika pengecekan map tanpa nil
	// var person2 map[string]string

	// if person2["name"] == "" {
	// 	fmt.Println("Data Kosong")
	// } else {
	// 	fmt.Println(person2)
	// }

	//tips simple
	var person2 map[string]string = NewMap("Alwan")
	if person2 == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person2)
	}
	//hanya beberapa tipe data seperti map,slice dll.lihat di atas
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
