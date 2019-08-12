package main

import "fmt"
import "os"

func main() {
	var argsRaw = os.Args
	fmt.Printf("-> %#v\n", argsRaw) 
	// -> []string{"GOPATH", "banana", "potato", "ice cream"}
	
	var args = argsRaw[1:] // ambil slice ke 1 sampai seterusnya
	// var args = argsRaw[1:3] // ->[]string{"banana","potato"} (dimulai dari slice 1 hingga sebelum slice 3)
	fmt.Printf("-> %#v\n", args)
	// -> []string{"banana", "potato", "ice cream"}

	/* 
		RUN   = jalankan dengan cara $ go run namafile.go banana potato "ice cream"
		BUILD = jalankan dengan cara $ go build namafile.go
				namafile.exe banana potato "ice cream" (windows)
				./namafile banana potato "ice cream" (linux)
	*/
}