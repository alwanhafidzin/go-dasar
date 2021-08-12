//Map adalah kumpulan data key-value yang unik dan tidak boleh sama
//Data boleh sebanyak banyaknya
//Mengeset data akan mereplace data yang baru

package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Alwan",
		"address": "Mojokerto",
	}

	//Menambah data
	person["title"] = "Programmer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	//len:mengambil jumlah data di map
	//map[key] = value : mengambil data di dalam map dengan key
	//make(map[TypeKey]TypeValue) : Membuat map baru
	//delete(map, key) :menghapus data di map dengan key

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Alwan"
	book["ups"] = "Salah"

	fmt.Println(len(book))
	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))

	//map len tergantung dari isi data
}
