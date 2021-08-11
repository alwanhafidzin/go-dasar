//var
//unique
package main
import "fmt"
func main(){
	var name string
	name = "Alwan Hafidzin"
	fmt.Println(name)

	//variabel bisa diubah
	name = "Alwan Firdaus"
	fmt.Println(name)

	// name =true error karena tipe data berbeda

	//golang bisa mendeteksi tipe data otomatis
	var friendName = "Hafidzin"
	fmt.Println(friendName)

	var age = 21
	fmt.Println(age)

	//bisa menggunkan var atau :=

	halo := "Alwan"
	fmt.Println(halo)


	//golang bisa multi variabel
	var (
		firstName = "Alwan"
		lastName = "Hafidzin"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}