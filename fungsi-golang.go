/* file:///C:/Users/Maslow/Downloads/dasarpemrogramangolang.pdf hal(54) */
package main

import (
	"fmt"
	"math"
	// "math/rand"
	// "time"

	// "strings"

)
/*			Penerapan function			*/

// func main() {
// 	var names = []string{"M.Fadhly","NR"}
// 	printMessage("Halo",names)
// }

// func printMessage(message string, arr []string){
// 	var nameString = strings.Join(arr," ")
// 	fmt.Println(message, nameString)
// }

/*============= Fungsi  dengan return value / Nilai balik(Void) ============*/
// func main() {
// 	rand.Seed(time.Now().Unix()) // untuk memastikan bahwa angka benarbenar diacak
// 	var randomValue int

// 	randomValue = randomWithRange(2,10)
// 	fmt.Println("Random number : ",randomValue)
	
// 	randomValue = randomWithRange(2,10)
// 	fmt.Println("Random number : ",randomValue)

// 	randomValue = randomWithRange(2,10)
// 	fmt.Println("Random number : ",randomValue)
// }

// // Function generate angka acak sesuai dengan range yang ditentukan
// func randomWithRange(min, max int) int {
// 	var value = rand.Int() % (max - min + 1) + min
// 	return value
// }

/*==================== Penggunaan keyword return untuk menghentikan proses ==================*/
func main() {
	dividedNumber(10, 2)
	dividedNumber(4, 0)
	dividedNumber(8, -4)
}

func dividedNumber(m, n int){
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n",m, n)
		return
	}
	var res = m / n
	fmt.Printf("%d / %d = %d\n",m, n, res)
}

/*==================== Multiple return ==================*/
func calculate(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d / 2, 2)
	// hitung keliling
	var circumference = math.Pi * d
	// kembalikan 2 nilai
	return area, circumference
}