package main

import "fmt"
import "strconv"

func main() {
	var input string
	fmt.Print("Type some number: ")
	fmt.Scanln(&input) // untuk mengcapture inputan yang di ketik oleh user sebelum menekan enter

	var number int
	var err error
	number, err = strconv.Atoi(input) // konversi ke tipe numerik, fungsi ini mengembalikan 2 data yang pertama yaitu number(berisi hasil konversi) dan yang kedua yaitu err(berisi informasi error)
	fmt.Println(number)

	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not a number")
		fmt.Println(err.Error())
	}
}