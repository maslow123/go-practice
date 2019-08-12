package main
import "fmt"

func main() {
	var numberA int = 4
	// ambil nilai referency dari numberA	
	var numberB *int = &numberA

	fmt.Println("numberA (value)    :", numberA)
	fmt.Println("numberB (address)  :", &numberA)

	fmt.Println("numberB (value)    :", *numberB) // deference
	fmt.Println("numberB (address)  :", numberB) // reference

	numberA = 5
	/* #####	MAKA AKAN BERUBAH JUGA ##### */
	
	fmt.Println("numberA (value)    :", numberA)
	fmt.Println("numberB (address)  :", &numberA)

	fmt.Println("numberB (value)    :", *numberB) // deference
	fmt.Println("numberB (address)  :", numberB) // reference

}