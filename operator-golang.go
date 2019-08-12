package main

import "fmt"

func main () {
	var value=(( (2+6) % 3) * 4 - 2) / 3
	var isEqual = ( value == 2)

	fmt.Printf("nilai %d (%t) \n",value, isEqual) // %d = desimal, %t = bool
	
	/*			Operator Logika		*/
	var left = false
	var right = true

	var leftAndRight = left && right// False dan true == false
	fmt.Printf("Left && Right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right // False atau true == true
	fmt.Printf("Left || Right \t(%t) \n", leftOrRight)

	var leftReverse = !left // Lawan dari false == True
	fmt.Printf("!Left \t\t(%t) \n", leftReverse)

}
