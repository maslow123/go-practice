package main

import (
	"fmt"
	"strings"
)
func main(){
	var data = []string{"Wick","Jason","Ethan"} // slice
	var dataContains0 = filter(data, func(each string) bool {
		return strings.Contains(each,"o")
	})
	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})
	
	fmt.Println("Data asli \t\t:",data)
	//data asli : [Wick Jason Ethan]

	fmt.Println("Filtering ada huruf \"o\"\t:", dataContains0)
	//filtering ada huruf "o" : [jason]

	fmt.Println("Filter jumlah huruf \"5\"\t:", dataLength5)
	//filter jumlah huruf "5": [jason ethan]
}

// fungsi filter untuk filtering data array(yang datanya dapat dari parameter pertama), dengan kondisi filter bisa ditentukan sendiri
func filter(data[]string, callback func(string)bool)[]string{ // function callback akan dipanggil disetiap perulangan dalam fungsi filter()
	var result []string
	// looping data dari parameter
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}