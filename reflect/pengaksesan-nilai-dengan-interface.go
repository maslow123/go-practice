package main 

import "fmt"
import "reflect"

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("Tipe variable :", reflectValue.Type())
	fmt.Println("Nilai variable:", reflectValue.Interface()) // mengembalikan nilai interface kosong || interface{}. Dan nilai aslinya hanya bisa diakses dengan meng-casting interface kosong tersebut
	var nilai = reflectValue.Interface().(int)
	fmt.Println(nilai)
}