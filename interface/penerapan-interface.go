/*
	Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), yang dibungkus dengan
	nama tertentu. Interface merupakan tipedata. Nilai objek bertipe interface zero value-nya adalah nil. Dan interface
	dapat digunakan jika sudah ada isinya
*/
package main 

import(
	"fmt"
	"math"
)

type Hitung interface {
	luas() float64
	keliling() float64
}

// struct bangun datar lingkaran
type lingkaran struct {
	diameter float64
}

/********** METHOD LINGKARAN **********/
func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	// Math powel = untuk perpangkatan
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}
/********** END METHOD **********/


// struct bangun datar persegi
type persegi struct {
	sisi float64
}
/********** METHOD PERSEGI **********/
func (p persegi) luas() float64 {
	return math.Pow(p.sisi,2)
}

func (p persegi) keliling() float64 {
	return p.sisi *4
}
/********** END METHOD **********/

func main() {
	var bangunDatar Hitung // bertipe interface hitung

	bangunDatar = persegi{10.0}
	fmt.Println("===== persegi")
	fmt.Println("luas		:", bangunDatar.luas())
	fmt.Println("keliling	:", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas		:", bangunDatar.luas())
	fmt.Println("keliling	:", bangunDatar.keliling())
	fmt.Println("jari-jari	:", bangunDatar.(lingkaran).jariJari())
}