//function bisa mengembalikan data
//untuk memberitahu kita bisa menuliskan tipe data kembaliannya
//jika function di deklarasikan tipe data pengembalian,maka wajib mengembalikan nilai

package main

import "fmt"

func main() {
	result := getHello("Alwan")
	fmt.Println(result)

	//gak berguna karena memanggil data yang return valuenya tidak dipake
	fmt.Println(getHello("tes"))
}

func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
	// return "Hello" + name
	//setelah return code tidak bisa dieksekusi,return selalu di akhir
}
