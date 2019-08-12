/*
	Fungsi closure adalah sebuah fungsi yang dapat disimpan dalam variable, dengan menggunakan fungsi ini
	kita dapat membuat FUNGSI DIDALAM FUNGSI serta kita dapat membuat fungsi yang mengembalikan fungsi
*/

package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (int, int) { // pembuatan fungsi closure
		var min, max int
		// lakukan perulangan untuk mendapatkan nilai max dan min
		for i, e := range n {
			fmt.Println(max)
			switch {
			case i == 0 :
				max, min = e, e
			case e > max :
				max = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3} // slice
	var min, max = getMinMax(numbers) // pemanggilan fungsi closure
	fmt.Printf("data : %v\nmin : %v\nmax: %v\n", numbers, min, max) // %v ini untuk menampilkan segala jenis tipe data
}
