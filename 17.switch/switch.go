//switch juga adalah percabangan
//switch lebih sederhana dari if
//biasanya switch digunakan untuk pengecekan sederhana

package main

import "fmt"

func main() {
	name := "Alwan"

	switch name {
	case "Alwan":
		fmt.Println("Halo Alwan")
		fmt.Println("Halo Alwan")
	case "Tes":
		fmt.Println("Halo Tes")
		fmt.Println("Halo Tes")
	default:
		fmt.Println("Kenalan Donk")
		fmt.Println("Kenalan Donk")
	}

	//Switch juga mendukung short statemen

	switch length := len(name); length > 3 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	//di golang switch bisa tanpa menggunakan kondisi

	length2 := len(name)
	switch {
	case length2 > 5:
		fmt.Println("Nama Terlalu Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}
