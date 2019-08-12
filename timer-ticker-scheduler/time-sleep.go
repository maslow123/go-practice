package main

import "fmt"
import "time"

func main() {
	fmt.Println("start")
	time.Sleep(time.Second * 2) // jeda waktu 4 detik untuk mengeksekusi statement selanjutnya
	fmt.Println("after 2 seconds")
	time.Sleep(time.Second * 2)
	fmt.Println("after 2 seconds")

	// Scheduler menggunakan time.Sleep()
	for i:= 0; i <= 10; i++ {
		fmt.Println("Anak ayam netes",i)
		time.Sleep(1 * time.Second)
	}
}