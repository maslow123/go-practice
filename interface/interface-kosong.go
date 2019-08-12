package main 

import "fmt"

func main() {
	var secret interface{} // jika kita tambahkan {} maka kegunaannya menjadi tipedata

	secret = "Budi Utomo"
	fmt.Println(secret)

	secret = []string{"Apple","Banana","Manggo"} // slice
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)
	fmt.Println("================")

	var data map[string]interface{} // interface kosong
	data = map[string]interface{} { // bukti bahwa interface merupakan tipe data
		"name":		"Budi utomo",
		"grade":	2,
		"breakfast":[]string{"Apple","Banana","Manggo"},
	}
	for _,el := range data {
		fmt.Println(el)
	}
}