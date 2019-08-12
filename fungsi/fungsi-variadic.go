package main

import "fmt"
func main() {
	var avg = calculating(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata-rata: %.2f",avg) // sprintf untuk mengembalikan nilai agar dapat diletakkan di variable

	fmt.Println(msg)
}

func calculating(numbers ...int) float64 { // fungsi dari (... int) yaitu untuk menampung banyaknya nilai yang diisikan pada parameter numbers
	var total int = 0
	for _, number := range numbers {
		total += number // total = total + number
		fmt.Println(total)
	}
	fmt.Println("Panjang :",len(numbers))

	var avg = float64(total) / float64(len(numbers))
	return avg
}