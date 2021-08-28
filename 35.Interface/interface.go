//interface merupakan tipe data abstract
//sebuah interface berisikan definisi-definisi method
//biasanya interface digunakan sebagai kontrak
/**
Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
Sehingga kita tidak perlu mengimplementasikan interface secara manual
Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface mana
*/

package main

import "fmt"

type HasName interface {
	//function tipe string
	GetName() string
}
type Person struct {
	Name string
}

//agar tidak error
func (person Person) GetName() string {
	return person.Name
}
func sayHello(hasName HasName) {
	//ngikuti kontrak
	//harus sama function dan parameternya
	fmt.Println("Hello", hasName.GetName())
}
func main() {
	var alwan Person
	alwan.Name = "Alwan"

	sayHello(alwan)

	cat := Animal{
		Name: "Push",
	}
	sayHello(cat)
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}
