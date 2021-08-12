//di golang hanya for loops

package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	//dalam for loop ada statement,dimana terdapat 2 statement yang bisa ditambahan
	//init statement:pertama kali diekseskusi
	//post statement:dieksekusi tiap perulangan
	//angka := 0; init statement
	//angka++ : post statement
	for angka := 0; angka <= 10; angka++ {
		fmt.Println("Perulangan Ke", angka)
	}

	//for range bisa digunakan untuk iterasi terhadap semua data collection(Array,Slice,Map)
	slice := []string{"Alwan", "Hafidzin", "Firdaus"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	names := []string{"Alwan", "Hafidzin", "Firdaus"}
	//di golang variabel yang tidak digunakan akan error,bisa diganti dengan underscore
	for i, value := range names {
		fmt.Println("Index", i, "#", value)
	}
	for _, value := range names {
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Alwan"
	person["title"] = "Programmer"
	//array dan slice berupa index,map berupa key
	for key, value := range person {
		fmt.Println(key, "#", value)
	}
}
