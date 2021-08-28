//struct adalah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
//struct biasanya representasi data dalam program aplikasi yang kita buat
//data di struct disimpan dalam field
//sederhananya struct adalah kumpulan dari field
//di golang tidak ada oop atau class
//struct boleh tipe data berbeda
//struct tidak bisa digunakan langsung
//namun kita bisa membuat data/object dari struct yang kita buat

package main

import "fmt"

func main() {
	// cara 1
	// var alwan Customer
	// alwan.Name = "Alwan"
	// alwan.Adress = "Mojokerto"
	// alwan.Age = 21

	//cara 2
	// alwan := Customer{
	// 	Name:   "Alwan",
	// 	Adress: "Mojokerto",
	// 	Age:    21,
	// }

	//yang ini harus sama jumlah data dan urutan dari struct,yang cara 1 dan 2 gak perlu sama
	alwan := Customer{"Alwan", "Mojokerto", 21}
	//lebih baik menyimpan representasi data dalam struct daripada map atau array,karena lebih bagus structurenya
	fmt.Println(alwan.Name)
	fmt.Println(alwan)
}

//rata2 menggunakan huruf besar di awal saat gunakan struct
type Customer struct {
	Name, Adress string
	Age          int
}
