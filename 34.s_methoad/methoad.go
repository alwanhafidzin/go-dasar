//Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
//Namun jika kita ingin menambahkan methoad ke dalam struct, sehingga seakan-akan sebuah struct memiliki function
//method adalah function

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

	//untuk yang pertama
	// sayHi(alwan, "Alwan")

	// yang kedua
	alwan.sayHi("Alwan")
	alwan.sayHuuu()
	//lebih baik menyimpan representasi data dalam struct daripada map atau array,karena lebih bagus structurenya
	fmt.Println(alwan.Name)
	fmt.Println(alwan)
}

//rata2 menggunakan huruf besar di awal saat gunakan struct
type Customer struct {
	Name, Adress string
	Age          int
}

// func sayHi(customer Customer, name string) {
// 	fmt.Println("Hello", name, "My Name is ", customer.Name)

// }
//struct methoad
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is ", customer.Name)

}

func (a Customer) sayHuuu() {
	fmt.Println("Huuu from", a.Name)
}
