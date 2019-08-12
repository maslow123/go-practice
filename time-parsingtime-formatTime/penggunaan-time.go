package main

import "fmt"
import "time"

func main() {
	var time1 = time.Now()
	fmt.Printf("time1 %v\n", time1)
	// time1 2019-08-09 13:29:09.3162359 +0700 +07 m=+0.001000101

	var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC) // time date memiliki 8 buah parameter
	//					 (tahun, bulan, tanggal, jam, menit, detik, nanodetik, timezone)
	fmt.Printf("time2 %v\n", time2)
	// time2 2011-12-24 10:20:00 +0000 UTC

	/* Method milik time.Time */
	var now = time.Now()
	fmt.Println("year:", now.Year(), "month:", now.Month())
	// year: 2019 month: August
}