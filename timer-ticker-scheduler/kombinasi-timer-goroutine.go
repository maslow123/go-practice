package main

import "fmt"
import "os"
import "time"

// buat function timer(), nantinya fungsi ini dieksekusi sebagai goroutine. Didalam function timer() terdapat blok kode jika waktu sudah mencapai timeout, maka sebuah data dikirimkan lewat channel *ch*

func timer(timeout int, ch chan <- bool){ // func timer hanya dapat mengirim channel
	time.AfterFunc(time.Duration(timeout) * time.Second, func() {
		ch <- true // mengirim data ke channel ch
	})
}

// siapkan juga function watcher(). Fungsi ini juga akan dieksekusi sebagai goroutine. Tugasnya cukup sederhana, yaitu menerima data dari channel *ch* (jika ada penerimaan data, berarti sudah masuk waktu timeout), lalu menampilkan pesan bahwa waktu telah habis

func watcher(timeout int, ch <- chan bool) {
	<- ch // menerima data dari channel *ch*
	fmt.Println("\ntime out! no answer more than", timeout, "seconds")
	os.Exit(0)
}

// buat implementasinya di main
func main() {
	var timeout = 5
	var ch = make(chan bool)

	go timer(timeout, ch) // membuat goroutine timer
	go watcher(timeout, ch) // membuat goroutine watcher

	var input string
	fmt.Print("What is 725/25 ? ")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("The answer is right! congratulations !")
	} else {
		fmt.Println("The answer is wrong ! ")
	}
}