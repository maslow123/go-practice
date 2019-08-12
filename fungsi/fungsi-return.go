package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // untuk memastikan bahwa angka yang digenerate benar benar acak
	var randomValue int

	randomValue = randomWithRange(2,10)
	fmt.Println("Random number :", randomValue)
	randomValue = randomWithRange(2,10)
	fmt.Println("Random number :", randomValue)
	randomValue = randomWithRange(2,10)
	fmt.Println("Random number :", randomValue)
}

func randomWithRange(min, max int)int {
	var value = rand.Int() % (max - min + 1) + min
	// fmt.Println(rand.Int()) // nilai random 
	return value// mengembalikan nilai value
}

/* ########## PENGGUNAAN RETURN UNTUK MENGHENTIKAN PROSES ########## */
