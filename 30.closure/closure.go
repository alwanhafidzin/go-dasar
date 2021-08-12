//closure kemampuan function berinteraksi dengan data2 disekitarnya dalam scope yang sama
//harap gunakan closure dengan bijak,agar tidak merubah data yang berubah

package main

import "fmt"

func main() {
	counter := 0
	name := "Alwan"

	increment := func() {
		//hati2 nama akan berubah dari scope di atas
		//scope diatas bisa diakses closure,tapi tidak bisa diakses dibawah
		// name = "Hafidzin"
		fmt.Println("Increment")
		//mencegah nya dengan deklarasi ulang
		name := "Alw"
		fmt.Println(name)
		counter++
	}
	//akan menjadi dua karena tidak disengaja
	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
