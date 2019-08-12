package main 

import "fmt"
import "encoding/base64"

func main() {
	var data ="M Fadhly NR"

	var encode = make([]byte, base64.StdEncoding.EncodedLen(len(data))) // menghasilkan informasi lebar data ketika sudah di encode, nilai ini kemudian ditentukan sebagai lebar alokasi tipe *[]byte* pada variable *encode* nantinya dipakai untuk menampung hasil encoding
	base64.StdEncoding.Encode(encode, []byte(data))
	var encodeString = string(encode)
	fmt.Println(encodeString)

	var decode = make([]byte, base64.StdEncoding.DecodedLen(len(encode)))
	var _, err = base64.StdEncoding.Decode(decode, encode)
	if err != nil {
		fmt.Println(err.Error())
	}
	var decodedString = string(decode)
	fmt.Println(decodedString)
}