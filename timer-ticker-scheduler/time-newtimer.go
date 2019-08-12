package main

import "fmt"
import "time"

func main() {
	var timer = time.NewTimer(4 * time.Second) // mengindikasikan bahwa nantinya akan ada data yang dikirimkan ke channel timer.C setelah 4 detik
	fmt.Println("start")
	<- timer.C // menandakan penerimaan data dari channel *timer.C*
	fmt.Println("finish")
}