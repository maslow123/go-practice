/*
	Kegunaan fungsi ini mirip *timeSleep()*, perbedaannya fungsi *timer.After()* akan mengembalikan data channel, sehingga perlu menggunakan tanda <- dalam penerapannya
*/

package main

import "fmt"
import "time"

func main() {
	fmt.Println("Start")
	<- time.After(4 * time.Second)
	fmt.Println("expired")
}