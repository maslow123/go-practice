package main

import "fmt"
import "time"

func main() {

	// var date, err = time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB") // error
							//(Format yang diinginkan, waktu sekarang)
	var date, err = time.Parse("06 Jan 15 03:04 MST", "02 Sep 15 08:00 WIB") // tidak error

	if err != nil {
		fmt.Println("error", err.Error())
		return
	}

	fmt.Println(date)
}