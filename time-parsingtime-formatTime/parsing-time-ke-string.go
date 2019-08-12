package main

import "fmt"
import "time"

func main() {

	var dates, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB") // parsing waktu sekarang 

	var dateS1 = dates.Format("Monday 02, January 2006 15:04 MST") // membentuk output string sesuai dengan layout format yang diinginkan
	fmt.Println("dateS1", dateS1)
	// Wednesday 02, September 2015 08:00 WIB

	var dateS2 = dates.Format(time.RFC3339)
	fmt.Println("dateS2", dateS2)
}