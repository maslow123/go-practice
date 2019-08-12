package main 

import "fmt"
import "crypto/sha1"

func main() {
	var text = "MyPassword"
	var sha = sha1.New() // variable dari sha1.New() adalah objek bertipe hash.Hash, dan memiliki 2 method Write() dan Sum()
	sha.Write([]byte(text)) // set data yang akan di hash, dan data harus berbentuk *[]byte*

	var encrypted = sha.Sum(nil) // eksekusi proses hash, dan menghasilkan data yang sudah di hash dalam bentuk *[]byte*
	var encryptedString = fmt.Sprintf("%x", encrypted) // ambil bentuk heksadesimal string dari data yang sudah di hash dengan memanfaatkan fungsi Sprintf() dengan layout format "%x"

	fmt.Println(encryptedString)
}