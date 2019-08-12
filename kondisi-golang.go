package main

import "fmt"

func main () {
	var point = 5

	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir lulus")
	} else {
		fmt.Printf("Tidak lulus, nilai anda %d\n", point)
	}

	/*			Variable Temporary pada if-else			*/
	var points = 8840.0

	if percent := points / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%") // 100.0% perfect !
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%") // 88.4% good
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%") // 60.5% not bad
	}

	/*			Switch Case */

	var pointSwitch = 5

	// switch pointSwitch {
	// case 8, 9, 10 :
	// 	fmt.Println("Perfect")
	// case 7, 6, 5, 4 :
	// 	fmt.Println("Awesome !")
	// default :
	// 	{
	// 		fmt.Println("Not Bad !")	
	// 		fmt.Println("You can be better !")
	// 	}
	// }

	/*			Switch dengan gaya if-else			*/
	
	// switch {
	// case pointSwitch == 8 :
	// 	fmt.Println("Perfect !")
	// case (pointSwitch < 8) && (pointSwitch > 3):
	// 	fmt.Println("Awesome !")
	// default :
	// 	{
	// 		fmt.Println("Not bad")
	// 		fmt.Println("You can be better !")
	// 	}
	// }

	/*			Seleksi kondisi bersarang			*/
	
	fmt.Println("Your point :",pointSwitch)
	if pointSwitch > 7 {
		switch pointSwitch {
		case 10:
			fmt.Println("Perfect !")
		default:
			fmt.Println("Nice!")
		}
	} else {
		if pointSwitch == 5 {
			fmt.Println("Not bad..")
		} else if pointSwitch == 3 {
			fmt.Println("Keep Trying !")
		} else {
			fmt.Println("You can do it !")
			if pointSwitch == 0 {
				fmt.Println("Try harder !")
			}
		}
	}
}
