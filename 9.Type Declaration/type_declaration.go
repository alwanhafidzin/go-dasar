//Kemampuan membuat ulang tipe data

//untuk membuat alias tipe data yang sudah ada,agar lebih mudah dimengerti

package main
import "fmt"

func main() {
	//melakukan deklarasi dengan type
	type Nktp string
	type Man bool

	//noKtpsaya akan menjadi string sesuai type
	var noKtpsaya Nktp = "13414311432431"
	//IamMan jadi boolean juga tanpa perlu menyebutkan type data
	var IamMan = true
	fmt.Println(noKtpsaya)
	fmt.Println(IamMan)
}