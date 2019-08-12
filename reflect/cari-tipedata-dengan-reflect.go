/*
	Reflection adalah teknik untuk inspeksi variable, mengambil informasi dari variable tersebut atau bahkan memanipulasi
	nya.Cakupan informasi yang bisa didapatkan lewat reflection sangat luas, seperti melihat struktur variable, tipe,
	nilai pointer dan banyak lagi

	reflect.ValueOf() = mengembalikan objek dalam tipe reflect.Value, yang mana berisikan informasi yang berhubungan
						dengan nilai pada variable yang dicari
	reflect.TypeOf() = mengembalikan objek dalam tipe reflect.Type, yang mana berisikan informasi yang berhubungan
					   dengan tipe data variable yang dicari
*/

package main

import "fmt"
import "reflect"

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number) // mengembalikan objek dalam tipe reflect.value yang berisikan informasi mengenai variable yang bersangkutan

	fmt.Println("Tipe  variable :", reflectValue.Type()) // mengembalikan tipe data variable dalam bentuk string

	// cek jenis tipe data dari variable reflectValue
	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai variable :", reflectValue.Int())
	}else {
		fmt.Println("Nilai variable :", reflectValue.String())
	}
}