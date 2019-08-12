package main

import "fmt"

func main () {
	var positifNumber uint8 = 89
	var negatifNumber = -1243423644

	fmt.Printf("Bilangan positif: %d\n", positifNumber)
	fmt.Printf("Bilangan negatif: %d\n\n", negatifNumber)

	/*			Tipe data numerik desimal			*/
	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal : %f\n", decimalNumber) // 2.620000
	fmt.Printf("bilangan desimal: %.3f\n\n",decimalNumber) // 2.620
	// fmt.Printf("bilangan desimal: %.2f\n",decimalNumber) // 2.62

	/*			Tipe data bool			*/
	var exist bool = true
	fmt.Printf("exist ? %t \n\n",exist)

	/*			Tipe data string		*/
	// var message string = "Halo"
	
	// with backticks
	var messages = `Nama saya "M.Fadhly NR".
	Salam kenal ya temanteman.
	Mari barengbareng belajar golang (:`

	fmt.Printf("message: %s \n\n",messages)

	/*			Nilai nil & Zero Value			*/
	const firstName ="M.Fadhly"
	fmt.Printf("halo %s",firstName+" !\n")

}
