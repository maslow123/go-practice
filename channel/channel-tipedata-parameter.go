/*
	Variable channel bisa di-pass ke fungsi lain sebagai parameter. Cukup tambahkan keyword CHAN pada deklarasi parameter agar operasi pass channel variable bisa dilakukan
*/

package main 

import "fmt"
import "runtime"

func printMessage(what chan string) {
	fmt.Println(<-what)
}

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	for _,each := range []string{"fadhly","noor","rizqi"} { // iterasi data slice/array langsung pada saat inisialisasi
		// eksekusi goroutine pada iife
		go func(who string) {
			var data = fmt.Sprintf("Hello %s",who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}