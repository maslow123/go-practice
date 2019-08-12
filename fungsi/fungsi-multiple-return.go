package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter) // pakai nilai kembalian dari func calculate

	fmt.Printf("Luas lingkaran \t\t : %.2f\n",area)
	fmt.Printf("Keliling lingkaran \t : %.2f\n",circumference)
}
// func calculate (d float64) (float64, float64) {
	/*##### CARA PERTAMA #####*/

	// var area = math.Pi * math.Pow(d / 2, 2) // hitung keliling
	// var circumference = math.Pi * d // hitung luas

	// return area, circumference

	/*##### CARA KEDUA #####*/

func calculate (d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/ 2, 2) // 2 pangkat 2 = 4
	fmt.Printf("%.2f\n",math.Pi)
	circumference = math.Pi * d

	return 
}