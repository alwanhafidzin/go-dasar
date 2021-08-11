//Slice adalah potongan dari array
//Ukuran Slice bisa berubah
//Slice dan array terkoneksi,dimana slice mengakses sebagian atau seluruh data di Array
//Tipe : POinter(Penunjuk data pertama di array pada slice),length(panjang slice),capacity(kapasitas dari slice)  length harus kurang dari atau sama dengan capacity

//array[low:high] : dimulai dari index low ke high
//array[low:] : dimulai dari index low sampai index akhir array
//array[:high] : index 0-index sebelum high
//array[:]  : indexs 0-indeks akhir di array

package main

import "fmt"

func main() {
	//gunakan ... jika tidak tau jumlah len nya
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	//panjang sice
	fmt.Println(len(slice1))
	//capasotas slice
	fmt.Println(cap(slice1))

	//array dirubah slice berubah
	months[5] = "Diubah"
	fmt.Println(slice1)
	//Slice dirubah array berubah

	slice1[0] = "Mei Lagi"
	fmt.Println(slice1)

	// var slice2 = months[10:]
	var slice2 = months[2:4]
	fmt.Println(slice2)

	//Jika sudah tidak muat append akan membuat array baru
	var slice3 = append(slice2, "Alwan")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	//membuat slice baru dengan function make

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Alwan"
	newSlice[1] = "Hafidzin"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	//copySlice melakukan copy terhadap slice yang ada
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//hati2 membuat array akan menjadi slice

	//Array diikuti length atau ... ,slice tanpa diikuti length
	intArray := [5]int{1, 2, 3, 4, 5}
	intArray2 := [...]int{1, 2, 3, 4, 5}
	intSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(intArray)
	fmt.Println(intArray2)
	fmt.Println(intSlice)

	//Kebnayakan yang dipake adalah slice daripada array karena lebih dinamis
}
