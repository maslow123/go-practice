/*
	Channel merupakan sebuah variable, dibuat dengan menggunakan kombinasi MAKE dan CHAN. Variabel channel memiliki
	satu tugas, menjadi pengirim dan penerima data
*/
package main 

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string) // deklarasi variable bertipe channel string

	// buat closure untuk membuat pesan string yang kemudian dikirim via channel
	var sayHelloTo = func(who string) { 
		var data = fmt.Sprintf("Hello %s", who)
		messages <- data // sedang berlangsung proses pengiriman data dari variabel yang berada dikanan lewat channel yang berada di kiri(variabel data dikirim lewat channel messages)
	}
	// fungsi sayHelloTo dieksekusi tiga kali sebagai geroutine berbeda. Menjadikan 3 proses ini berjalan secara asynchronus / tidak saling tunggu
	go sayHelloTo("M.Fadhly NR")
	go sayHelloTo("M.Rasya NF")
	go sayHelloTo("M.Hilman NM")

	var message1 = <-messages // jika disebelah kanan channel maka menandakan proses penerimaan data dari channel yang dikanan, untuk disimpan ke variable yang dikiri
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)
	


}