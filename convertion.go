package main


//hati hati konversi nilainya tidak cukup
import "fmt"

func main () {
	// var nilai32 int32 = 100000
	// var nilai32 int32 = 128
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	//melebihi batas akan ke bawah terus naik 128 ke -128, 129 ke -127
	fmt.Println(nilai8)


	var name = "Alwan"
	//Mengambil byte
	var e = name[0]
	var eSting string = string(e)

	fmt.Println(name)
	fmt.Println(eSting)
}