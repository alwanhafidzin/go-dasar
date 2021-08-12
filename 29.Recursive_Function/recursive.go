//recursive function adalah fungsi yang memanggil dirinya sendiri
//kadang lebih mudah dengan adanya recursive function
//contoh kasus yang lebih mudah dengan recursive function

package main

import "fmt"

func main() {
	loop := factorialLoop(5)
	fmt.Println(loop)
	fmt.Println(5 * 4 * 3 * 2 * 1)

	recursive := factorialRecursive(5)
	fmt.Println(recursive)
}

func factorialRecursive(value int) int {
	//akan berhenti ketika value 1,harus ada agar tidak infinite loop
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

//dengan looping,tanpa recursive
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}
