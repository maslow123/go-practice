package main

import "fmt"

func main () {
	for i := 0; i < 5; i++ {
		fmt.Println("Angka ke - ",i)
	}
	fmt.Printf("===============\n")
	/*			LOOPING ONLY CONDITION(WHILE)		*/
	// var i = 0

	// for i < 5 {
	// 	fmt.Println("Angka ke - ",i)
	// 	i++
	// }
	// fmt.Printf("===============\n")

	/*			LOOPING WITHOUT ARGUMENT			*/
	// for {
	// 	fmt.Println("Angka ke - ",i)
	// 	i++
	// 	if i == 1000 {
	// 		break // memberhentikan looping jika pengulangan sudah sampai 5 kali
	// 	}
	// }

	/*			LOOPING WITH CONTINUE AND BREAK			*/
	// for i:= 1; i <= 12; i++ {
	// 	if i % 2 == 1 { // Jika angkanya genap
	// 	// if i % 2 == 0 { // Jika angkanya ganjil
	// 		continue
	// 	}

	// 	if i > 8 {
	// 		break
	// 	}

	// 	fmt.Println("Angka ke ",i)
	// }

	/*			PERULANGAN BERSARANG		*/
	// for i := 0; i < 5; i++ {
		
	// 	for j := i; j < 5; j++ {
	// 		fmt.Print(j, " ")
	// 	}
	// 	fmt.Println()
	// }

	/*			PEMANFAATAN LABEL PADA PERULANGAN			*/
	outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("Matriks [", i, "][", j, "]", "\n")
		}
	}
}
