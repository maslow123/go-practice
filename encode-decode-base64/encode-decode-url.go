package main 

import "fmt"
import "encoding/base64"

func main() {
	var data ="https://kalipare.com/"

	var encodedString = base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedString)

	var decodedByte, _ = base64.URLEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println(decodedString)
}