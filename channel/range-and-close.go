/*
	For-range jika diterapkan pada channel berfungsi untuk handle penerimaan data. Setiap kali ada pengiriman data via channel, maka akan mentrigger perulangan for-range. Perulangan akan terus menerus seiring pengiriman data ke channel yang dipergunakan, dan perulangan akan di berhentikan apabila channel tersebut di close ataupun di nonaktifkan
*/
package main

import "fmt"
import "runtime"

func sendMessage(ch chan <- string) { // parameter ch hanya bisa mengirim data
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("Data %d", i)
	}
	close(ch)
}

// buat fungsi printMessage() untuk handle permintaan data.	
func printMessages(ch <- chan string) { // paramter ch hanya bisa menerima data
	for message := range ch {
		fmt.Println(message)
	}
}

// jalankan sendMessage() sebagai goroutine, dan fungsi sendMessage()
func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)
	go sendMessage(messages)
	printMessages(messages)
	// setelah 20 data sukses dikirim dan diterima, channel ch di non-aktifkan(close(ch))
}

