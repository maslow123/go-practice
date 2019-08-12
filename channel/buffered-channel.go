package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(2)

	messages := make(chan int,2) // memberikan nilai buffer 2, ketika bernilai 2 maka brrti jumlah buffer maksimal ada 3 dan pengiriman data index ke 0, 1, 2, 3 akan berjalan asynchronus, karena channel ditentukan nillai buffernya sebanyak 3(0,1,2,3) sisanya akan berjalan setelah ada salah satu dari keempat data yang sebelumnya dikirimkan sudah diterima

	go func() {
		for {
			i := <- messages
			fmt.Println("receive data", i)
		}
	}()

	for i:=0; i < 5; i++ {
		fmt.Println("send data", i)
		messages <- i
	}
}