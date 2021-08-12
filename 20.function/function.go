//function main wajib dibuat
//function sebuah blok kode agar bisa digunakan berulang
//golang tidak berorientasi object
//car menggunakan function dengan cara funct() {}
//memannggil function nama()

package main

import "fmt"

func main() {
	//case sensitif
	for i := 0; i < 10; i++ {
		sayHello()
	}

}

func sayHello() {
	fmt.Println("Hallo")
}
