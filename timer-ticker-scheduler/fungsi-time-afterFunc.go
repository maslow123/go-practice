package main

import "fmt"
import "time"

func main() {
	var ch = make(chan bool)

	time.AfterFunc( 4 * time.Second, func() {
		fmt.Println("expired")
		ch <- true // penerimaan data channel
	})

	fmt.Println("Start")
	<- ch
	fmt.Println("Finish")
}

/*
	- Jika tidak ada serah terima data lewat channel, maka eksekusi *time.AfterFunc()* adalah asynchronus(tidak blocking).
	- Jika ada serah terima data lewat channel, maka fungsi akan tetap berjalan asynchronus hingga baris kode dimana 	    penerimaan data channel dilakukan.
*/