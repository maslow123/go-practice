/*
	Dipersiapkan 2 buah goroutine, satu untuk pencarian rata-rata, dan satu untuk nilai tertinggi.
*/
package main

import "fmt"
import "runtime"

// make a goroutine for get average
func getAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, e := range numbers {
		sum += e
	}
	ch <- float64(sum) / float64(len(numbers))
}

// make a goroutine for get maxValue
func getMax(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}
	}
	ch <- max
}

func main() {
	runtime.GOMAXPROCS(2)

	var numbers = []int{3,4,3,5,6,3,2,2,6,3,4,6,3}
	fmt.Println("Numbers :",numbers)

	var ch1 = make(chan float64)
	go getAverage(numbers, ch1)

	var ch2 = make(chan int)
	go getMax(numbers, ch2)

	for i:=0; i < 2; i++ { // karena ada 2 channel, maka dilakukan perulang sebanyak 2 kali
		select { // pengiriman data dikontrol dengan menggunakan select
		case avg := <- ch1: // case ini akan terpenuhi ketika ada penerimaan data dari channel ch1, yang kemudian ditampung oleh variable avg
			fmt.Printf("Avg \t: %.2f \n", avg)
		case max := <- ch2:
			fmt.Printf("Max \t: %d \n", max)
		}
	}
}