//Break & continou digunakan dalam perulangan
//Break menghentikan perulangan
//Continou menghentikan perulangan yang berjalan dan melanjutkan ke perulangan selanjutnya,post statement akan terus dilanjutkan

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		// if i == 5 {
		// 	break
		// }
		//Jika genap diskip
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke ", i)
	}
}
