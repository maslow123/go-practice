package main

import "fmt"
import "reflect"

type student struct {
	Name string
	Grade int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	var s1 = &student{Name: "M.Fadhly", Grade: 2}
	fmt.Println("nama :", s1.Name)

	var reflectValue = reflect.ValueOf(s1)
	// cari reflect method
	var method = reflectValue.MethodByName("SetName")
	// eksekusi method
	method.Call([]reflect.Value {
		reflect.ValueOf("Noor Rizqi"),
	})

	fmt.Println("nama :", s1.Name)
}