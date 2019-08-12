package main 

import "fmt"
import "encoding/base64"

func main() {
	var data = "M Fadhly NR"

	var encodeString = base64.StdEncoding.EncodeToString([]byte(data)) // casting kedalam bentuk []byte, sebelum di encode menggunakan fungsi *base64.StdEncoding.EncodeToString()*
	fmt.Println("encoded:", encodeString) // hasilnya bertipe string

	var decodeByte, _ = base64.StdEncoding.DecodeString(encodeString) // di decode kembali ke tipe string aslinya, tapi bertipe *byte[]*
	var decodedString = string(decodeByte) // mengcasting dari *byte[]* ke tipedata *string()*
	fmt.Println("decoded:", decodedString)
}