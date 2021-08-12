// Sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama pada function tersebut
// Namun kadang ada kalanya lebih mudah membuat function secara langsung di variable atau parameter tanpa harus membuat function terlebih dahulu
// Hal tersebut dinamakan anonymous function, atau function tanpa nama

package main

import "fmt"

type Blacklist func(string) bool

func main() {
	//di bawah adalah anonymus function
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", blacklist)
	registerUser("Alwan", blacklist)

	//bisa seperti dibawah ini cuman harus satu2
	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("anjing", func(name string) bool {
		return name == "jing"
	})
}
func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

//dibawah terlalu lama
// func blacklistAdmin(name string) bool {
// 	return name == "admin"
// }
// func blacklistRoot(name string) bool {
// 	return name == "root"
// }
