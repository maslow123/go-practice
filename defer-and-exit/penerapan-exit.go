/*
	Fungsi os.Exit() berada dalam package os. Fungsi ini memiliki sebuah parameter bertipe numerik yang wajib diisi.
	Angka yang dimasukkan akan muncul sebagai exit status ketika program berhenti
*/

package main 

import "fmt"
import "os"

func main() {
	defer fmt.Println("Halo")
	os.Exit(1)
	fmt.Println("selamat datang")
}