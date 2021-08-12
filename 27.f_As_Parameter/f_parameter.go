//function tidak hanya disimpan di dalam variabel sebagai value saja
//bisa digunakan sebagai parameter untuk fungsi lain

package main

import "fmt"

type Filter func(string) string

func main() {
	sayHelloWithFilter("Alwan", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}

func sayHelloWithFilter(name string, filter Filter) {
	namedFiltered := filter(name)
	fmt.Println("Hello ", namedFiltered)
}

//mempermudah function sebagai parameter,daripada masuk di sayHelloWithFilter
func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

//kadang function agak panjang agak ribet,bisa menggunakan type declaration untuk membuat alias
