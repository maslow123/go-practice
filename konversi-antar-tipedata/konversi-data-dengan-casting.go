package main 

import "fmt"

func main() {
	var a float64 = float64(24.00)
	fmt.Println(a) // 24

	var b int32 = int32(24.00)
	fmt.Println(b) // 24

/* Casting STRING <-----> BYTE */
	// mendapatkan slice byte dari sebuah data string dengan mengcastingnya ketype []byte
	
	var text1 = "halo" // string sebenarnya merupakan slice/array byte
	var bb = []byte(text1)

	fmt.Printf("%d %d %d %d\n", bb[0], bb[1], bb[2], bb[3]) // 104 97 108 111

	// mendapatkan string dari slice byte
	var byte1 = []byte{104, 97, 108, 111}
	var s = string(byte1)

	fmt.Println(s) // halo

/* Casting STRING -----> INT */
	var c int64 = int64('h')
	fmt.Println(c) // 104

	var d string = string(104)
	fmt.Println(d) // h
}