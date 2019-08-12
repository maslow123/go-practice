package main

import "fmt"

func main() {
	type student struct {
		name		string
		height		float64
		age			int32
		isGraduated bool
		hobbies		[]string
	}

	var data = student {
		name		: "Fadhly",
		height		: 178.5,
		age			: 17,
		isGraduated : false,
		hobbies		: []string{"eating","sleeping"},
	}

	fmt.Printf("%b\n", data.age) // binner
	fmt.Printf("%c\n", 1400) // unicode
	fmt.Printf("%d\n", data.age) // memformat data numerik menjadi string
	fmt.Printf("%e\n", data.height) // memformat data numerik desimal kedalam bentuk notasi numerik
	fmt.Printf("%f\n", data.height) // memformat data numerik desimal, dengan lebar desimal bisa ditentukan. Defaultnya 6 digit
	fmt.Printf("%.1f\n", data.height)
	
	fmt.Printf("%g\n", 0.12312412312312412) // memformat data numerik desimal, dengan lebar desimal yang bisa ditentukan, cocok digunakan ketika jumlah digit desimal yang cukup banyak

	fmt.Printf("%o\n", data.age) // memformat data numerik, menjadi bentuk string oktal(8 digit)

	fmt.Printf("%p\n", &data.name) // memformat data pointer, mengembalikan alamat pointer referensi variabelnya
	fmt.Printf("%s\n", data.name) // memformat data string

	fmt.Printf("%T\n", data.name) // mengambil tipe data dari variabel yang akan di format

	fmt.Printf("%v\n", data) // digunakan untuk format apa saja, hasil kembaliannya adalah string nilai data aslinya.Jika data adalah objek cetakan struct, maka akan ditampilkan semua secara property berurutan

	fmt.Printf("%+v\n", data) // digunakan untuk memformat struct, mengembalikan nama tiap property dan nilainya berurutan sesuai dengan struktur struct

	fmt.Printf("%#v\n\n", data) // digunakan untuk memformat struct, mengembalikan nama dan nilai tiap property sesuai dengan struktur struct dan juga bagaimana objek tersebut dideklarasikan

	// ketika menampilkan objek yang deklarasinya adlaah menggunakan teknik anonymous struct
	var datas = struct {
		name		string
		height		float64
	}{
		name		: "Fadhly",
		height		: 178.5,
	}
	fmt.Printf("%#v\n", datas)
}