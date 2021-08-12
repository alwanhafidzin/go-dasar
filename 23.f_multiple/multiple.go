//di golang bisa multiple return
//untuk memberi tahu jika function menggembalikan multiple value,harus menulikan semua tipe return value
//jika ingin menghiraukan return value tersebut bisa menggunakna underscore(_)
package main

import "fmt"

func main() {
	firstName, _ := getFullName()
	//gak peduli semua, ;= butuh satu variabel
	// getFullName()
	fmt.Println(firstName)
	// fmt.Println(lastName)
}

func getFullName() (string, string) {
	return "Alwan", "Hafidzin"
}
