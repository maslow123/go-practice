package main

import "fmt"
import "time"

func main() {
	done := make(chan bool)
	ticker := time.NewTicker(time.Second)

	// buat goroutine untuk mengontrol kapan ticker harus di stop
	go func() {
		time.Sleep( 10 * time.Second) // wait 10 seconds
		done <- true // mengirim data ke channel done , dan setelah 10 detik maka masuk kedalam blok kode *for-select* dan ketika itu terjadi method *.Stop()* milik objek ticker dan dipanggil untuk menonaktifkan scheduler pada ticker tersebut
	}()

	for {
		select { // cek penerimaan data dari channel *done* dan *ticker.C*
		case <- done :
			ticker.Stop()
			return
		case t := <- ticker.C : // data yang dikirimkan *ticker.C* adalah data date-time kapan event itu terjadi
			fmt.Println("Hello !!", t)
		}
	}
}