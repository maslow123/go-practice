package main 

import "fmt"
import "reflect"

type student struct {
	Name string
	Grade int
}

// buat method baru dengan nama getPropertyInfo() untuk mengambil dan menampilkan info tiap property milik struct student
func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	// cek variable objek tersebut merupakan pointer atau bukan
	if reflectValue.Kind() == reflect.Ptr {
		// Jika true, maka mengambil object reflect dengan menggunakan elem()
		reflectValue = reflectValue.Elem()
	}
	var reflectType = reflectValue.Type()
	// lakukan perulangan sebanyak jumlah property yang ada pada struct student, method NumField() akan mengembalikan jumlah property publik yang ada dalam struct
	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("Nama		:", reflectType.Field(i).Name) // mengambil key
		fmt.Println("Tipedata	:", reflectType.Field(i).Type) // mengambil tipe data
		fmt.Println("Nilai		:", reflectValue.Field(i).Interface()) // mengambil value
		fmt.Println()
	}
}

func main() {
	var s1 = &student{Name : "Fadhly", Grade: 2}
	s1.getPropertyInfo()
}