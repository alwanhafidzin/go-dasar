//function first class citizen,artinya tidak ada turunan
//function juga merupakan tipe data dan bisa disimpan dalam variabel

//function value :function disimpan dalam value
package main

import "fmt"

func main() {
	goodbye := getGoodBye
	result := goodbye("Alwan")
	fmt.Println(result)
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}
