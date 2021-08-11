//
//&& jika 1 false mengembalikan false
// &&(dan)->true+true=true,true+false=false,false+true=false,false+false=false
// ||(Atau)->true+true=true,false+true=true,true+false=true,false+false=false
// !(kebalikan)->true=false,false=true

package main
import "fmt"

func main(){
	var nilaiAkhir = 98
	var absensi = 88

	var lulusNilaiAkhir = nilaiAkhir >80
	var lulusAbsensi = absensi > 80

	var lulus = lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)

	fmt.Println(ujian >=80 && absensi >=80)
}