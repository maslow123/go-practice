package main 

import(
	"fmt"
	"math"
)

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}
/* Menghitung luas, keliling, dan volume */
type hitung interface {
	hitung2d
	hitung3d
}

// siapkan struct KUBUS yang memiliki method luas(), keliling(), dan volume()
type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3) // k.sisi^3
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2)* 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

// MAIN FUNCTION
func main() {
	var bangunRuang hitung = &kubus{4} // ambil referensi terlebih dahulu setelah itu baru tampung ke variable interface

	fmt.Println("===== kubus")
	fmt.Println("luas		:",bangunRuang.luas())
	fmt.Println("keliling	:",bangunRuang.keliling())
	fmt.Println("volume 		:",bangunRuang.volume())
}