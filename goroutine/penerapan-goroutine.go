package main

import "fmt"
import "runtime"

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func Scanln(a ...interface{})(n int, err error){
	
	var s1,s2,s3 string
	fmt.Scanln(&s1, &s2, &s3)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	return
}

func main() {
	runtime.GOMAXPROCS(2) // menentukan jumlah core yang diaktifkan untuk eksekusi program

	go print(5, "halo") // pembuatan goroutine baru ditandai dengan diawali "go"
	print(5, "apakabar ?")

	var input string
	fmt.Scanln(&input) // mengakibatkan proses jalannya aplikasi berhenti dibaris itu(bloking) hingga user menekan tombol enter
	Scanln()
}