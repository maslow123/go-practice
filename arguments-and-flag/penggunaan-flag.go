/*
	Flag memiliki kegunaan yang sama seperti arguments, yaitu untuk *parameterize* eksekusi program, dengan penulisan dalam bentuk key-value
*/
/*
	RUN = go run namafile.go -name="fadhly" -age=17
		
	BUILD = jalankan dengan cara $ go build namafile.go
			namafile.exe -name="fadhly" -age=17 (windows)
			./namafile -name="fadhly" -age=17 (linux)

			namafile.exe --help(untuk memunculkan hints || petunjuk arguments apa saja yang bisa dipakai)
			
*/

package main

import "fmt"
import "flag"

func main() {
	// disiapkan flag bertipe *string* dengan key adalah *name*, dan nilai defaultnya "anonymous"
	// var name = flag.String("name", "anonymous", "type your name")

	// var age = flag.Int64("age", 25, "type your age")

	// flag.Parse()
	// fmt.Printf("name\t: %s\n", *name) // nilai balik dari flag adalah pointer, maka dari itu kita perlu di-dereference agar dapat nilai aslinya (*name)
	// fmt.Printf("age\t: %d\n", *age)

	/* Deklarasi flag dengan cara passing reference Variable penampung data */
	var data1 = flag.String("name", "anonymous", "type your name")
	var data2 string
	flag.StringVar(&data2, "gender", "male", "type your gender") // ambil nilai lewat parameter pointer

	flag.Parse()

	fmt.Println(*data1)
	fmt.Println(data2)

}
