/*
	Defer digunakan untuk mengakhirkan eksekusi baris kode. ketika eksekusi blok sudah hampir selesai,
	statement yang didefer yang dijalankan
*/
package main

import "fmt"

func main() {
	orderSomeFood("Pizza")
	orderSomeFood("Burger")
}

func orderSomeFood(menu string) {
	defer fmt.Println("Terima kasih, silahkan tunggu") // ini akan di eksekusi di penerakhiran.

	if menu == "pizza" {
		fmt.Println("Pilihan tepat!"," ")
		fmt.Print("Pizza ditempat kami paling enak!","\n")

		return
	}
	fmt.Println("Pesanan anda", menu)
}
