/*
	Closure jenis ini dieksekusi langsung pada saat deklarasinya. Biasa digunakan untuk membungkus
	proses yang hanya dilakukan sekali, bisa mengembalikan nilai, bisa juga tidak
*/

package main

import "fmt"

func main() {
	var numbers = []int {2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var newNumbers = func(min int) []int {
		var r []int
		for _ , e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3) // memfilter angka yang dimulai dari angka 3

	fmt.Println("Original number:",numbers)
	fmt.Println("Filtered number:",newNumbers)
}