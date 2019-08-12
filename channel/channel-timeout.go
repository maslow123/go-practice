/*
	Sebuah goroutine baru dijalankan dengan tugas mengirimkan data setiap interval tertentu, dengan durasi intervalnya sendiri adalah acak/random
*/

package main

import(
	"fmt"
	"runtime"
	"math/rand"
	"time"
)

func sendData(ch chan <- int) { // parameter ch hanya dapat mengirim data
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int() % 10 + 1) * time.Second)
	}
}

// lakukan perulangan tanpa henti, yang disetiap perulangan ada seleksi kondisi channel menggunakan select
func retreiveData(ch <- chan int) { // parameter ch hanya dapat menerima data
	loop :
	for {
		select {
		case data := <- ch : // akan terpenuhi ketika ada serah terima data pada channel
			fmt.Print(`receive data "`, data, `"`, "\n")
		case <- time.After(time.Second * 5) : // akan terpenuhi ketika tidak ada aktifitas penerimaan data dari channel dalam durasi 5 detik. Blok ini biasa disebut dengan blok TIMEOUT
			fmt.Println("timeout. no activities under 5 seconds")
			break loop
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	var messages = make(chan int)

	go sendData(messages)
	retreiveData(messages)
}
